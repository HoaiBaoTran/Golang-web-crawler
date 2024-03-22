package services

import (
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/repositories"
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

func (s *ExtractedDataService) CreateExtractedData() (models.ExtractedData, error) {
	return s.ExtractedDataRepository.CreateExtractedData(models.ExtractedData{})
}
