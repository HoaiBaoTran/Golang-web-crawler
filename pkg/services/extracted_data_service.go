package services

import (
	"fmt"

	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/repositories"
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

func (s *ExtractedDataService) CreateExtractedData(urlPath, wrappedTag string, maxDepth, levenshteinDistance int, tag, words []string) ([]models.ExtractedData, error) {
	exitChan := make(chan bool)
	dataChan := make(chan models.ExtractedData)
	var data []models.ExtractedData

	myCrawler := crawler.CreateCrawler(maxDepth)

	go func(urlPath string) {
		myCrawler.CrawlWeb(urlPath, maxDepth, tag, exitChan, dataChan)
	}(urlPath)

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
			fmt.Println("extracted-data: ", extractedData.Title)
			data = append(data, extractedData)
		case <-exitChan:
			repositories.LogMessage("Finish crawling ", urlPath)
			for _, smallData := range data {
				fmt.Println("Data: ", smallData.Title)
			}
			// return s.ExtractedDataRepository.CreateExtractedData(data)
		}
	}

}
