package services

import (
	"time"

	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/repositories"
	convertJSON "github.com/hoaibao/web-crawler/pkg/utils/convert-json"
	"github.com/hoaibao/web-crawler/pkg/utils/crawler"
	handleTag "github.com/hoaibao/web-crawler/pkg/utils/handle-html-tag"
	"github.com/hoaibao/web-crawler/pkg/utils/statistics"
)

type ExtractedDataService struct {
	ExtractedDataRepository repositories.ExtractedDataRepository
}

func NewExtractedDataService(extractedDataRepository repositories.ExtractedDataRepository) *ExtractedDataService {
	return &ExtractedDataService{
		ExtractedDataRepository: extractedDataRepository,
	}
}

func (s *ExtractedDataService) GetAllExtractedData() ([]models.ExtractedData, error) {
	return s.ExtractedDataRepository.GetAllExtractedData()
}

func (s *ExtractedDataService) GetExtractedDataById(id int) (models.ExtractedData, error) {
	return s.ExtractedDataRepository.GetExtractedDataById(id)
}

func (s *ExtractedDataService) CreateExtractedData(wrappedTag string, maxDepth, levenshteinDistance int, urlPath, tag, words []string) ([]models.ExtractedData, error) {
	exitChan := make(chan bool)
	dataChan := make(chan models.ExtractedData)
	var responseData []models.ExtractedData

	myCrawler := crawler.CreateCrawler(maxDepth)

	for _, url := range urlPath {
		go func(url string) {
			myCrawler.CrawlWeb(url, maxDepth, tag, exitChan, dataChan)
		}(url)
		time.Sleep(10 * time.Second)
	}
	for {
		select {
		case extractedData := <-dataChan:
			lineCount, wordCount, charCount, averageWordLength, frequency := statistics.Statistics(extractedData.Paragraph)
			extractedData.LineCount = lineCount
			extractedData.WordCount = wordCount
			extractedData.CharCount = charCount
			extractedData.AverageWordLength = averageWordLength
			extractedData.Frequency = frequency

			if levenshteinDistance >= 0 && len(words) > 0 {
				extractedData.Paragraph = handleTag.ChangeContentToHtmlTag(extractedData.Paragraph, words, wrappedTag, levenshteinDistance)
			}

			go func(extractedData models.ExtractedData) {
				convertJSON.WriteJsonFile(extractedData)
				responseExtractedData, err := s.ExtractedDataRepository.CreateExtractedData(extractedData)
				if err != nil {
					repositories.LogMessage("Can't not insert into database", err)
				}
				responseData = append(responseData, responseExtractedData)
			}(extractedData)

		case <-exitChan:
			repositories.LogMessage("Finish crawling ", urlPath)
			return responseData, nil
		}
	}

}
