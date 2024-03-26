package services

import (
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/repositories"
	"github.com/hoaibao/web-crawler/pkg/utils/crawler"
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

func (s *ExtractedDataService) CreateExtractedData(urlPath string, maxDepth, editDistance int, tag []string) ([]models.ExtractedData, error) {
	exit := make(chan bool)
	dataChan := make(chan models.ExtractedData)
	var data []models.ExtractedData

	myCrawler := crawler.CreateCrawler(maxDepth)
	go myCrawler.CrawlWeb(urlPath, maxDepth, editDistance, tag, exit, dataChan)

	for {
		select {
		case extractedData := <-dataChan:
			lineCount, wordCount, charCount, averageWordLength, frequency := statistics.Statistics(extractedData.Paragraph)
			extractedData.LineCount = lineCount
			extractedData.WordCount = wordCount
			extractedData.CharCount = charCount
			extractedData.AverageWordLength = averageWordLength
			extractedData.Frequency = frequency
			data = append(data, extractedData)
		case <-exit:
			// close(dataChan)
			// time.Sleep(10 * time.Second)
			// fmt.Println("finish crawling")

			return s.ExtractedDataRepository.CreateExtractedData(data)
		}
	}

}
