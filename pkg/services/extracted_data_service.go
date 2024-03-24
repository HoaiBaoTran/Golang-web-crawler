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

func (s *ExtractedDataService) CreateExtractedData(urlPath string) (models.ExtractedData, error) {

	myCrawler := crawler.CreateCrawler()

	relatedUrlChan := myCrawler.CrawlRelatedUrl(urlPath)

	// urlLink := "https://vnexpress.net/nong-dan-thu-nhap-gap-10-lan-neu-san-xuat-tom-lua-quy-mo-lon-4725097.html"
	// urlLink := "http://localhost:8081/nong-dan"
	extractedData := myCrawler.CrawlData(urlPath)

	lineCount, wordCount, charCount, averageWordLength, frequency := statistics.Statistics(extractedData.Paragraph)
	extractedData.LineCount = lineCount
	extractedData.WordCount = wordCount
	extractedData.CharCount = charCount
	extractedData.AverageWordLength = averageWordLength
	extractedData.Frequency = frequency

	relatedUrl := <-relatedUrlChan
	extractedData.RelatedUrl = relatedUrl

	return s.ExtractedDataRepository.CreateExtractedData(extractedData)
}
