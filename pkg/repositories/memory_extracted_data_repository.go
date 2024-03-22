package repositories

import (
	"database/sql"
	"os"

	"github.com/hoaibao/web-crawler/pkg/database"
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/utils/logger"
)

var (
	MyLogger = logger.InitLogger()
)

type MemoryExtractedDataRepository struct {
	ExtractedData map[int]models.ExtractedData
	DB            *sql.DB
}

func CheckError(err error, msg string) {
	if err != nil {
		// log.Fatal(err, msg)
		MyLogger.ConsoleLogger.Error(msg, err)
		MyLogger.FileLogger.Error(msg, err)
	}
}

func LogMessage(args ...interface{}) {
	MyLogger.ConsoleLogger.Infoln(args)
	MyLogger.FileLogger.Infoln(args)
}

func NewMemoryExtractedDataRepository() *MemoryExtractedDataRepository {
	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}

	db, err := database.NewConnection(config)
	CheckError(err, "Can't connect database")

	return &MemoryExtractedDataRepository{
		ExtractedData: make(map[int]models.ExtractedData, 0),
		DB:            db,
	}
}

func (r *MemoryExtractedDataRepository) GetAllExtractedData() ([]models.ExtractedData, error) {
	return []models.ExtractedData{}, nil
}

func (r *MemoryExtractedDataRepository) GetExtractedDataById(id int) (models.ExtractedData, error) {
	return models.ExtractedData{}, nil
}

func (r *MemoryExtractedDataRepository) CreateExtractedData(data models.ExtractedData) (models.ExtractedData, error) {
	return models.ExtractedData{}, nil
}
