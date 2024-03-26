package repositories

import "github.com/hoaibao/web-crawler/pkg/models"

type ExtractedDataRepository interface {
	GetAllExtractedData() ([]models.ExtractedData, error)
	GetExtractedDataById(id int) (models.ExtractedData, error)
	CreateExtractedData(data []models.ExtractedData) ([]models.ExtractedData, error)
}
