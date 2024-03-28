package repositories

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/hoaibao/web-crawler/pkg/database"
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/utils/logger"
	goDotEnv "github.com/joho/godotenv"
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

	err := goDotEnv.Load("/Users/hoaibao/Desktop/Workspace/Go/FPT_Assignments/web-crawler/.env")
	CheckError(err, "Can't load value from .env")

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

func (r *MemoryExtractedDataRepository) CreateExtractedData(extractedData models.ExtractedData) (models.ExtractedData, error) {

	LogMessage("Data Repo: ", extractedData.Title)
	insertFrequencyStatement := `INSERT INTO word_frequency(word, frequency, extracted_data_id) VALUES `
	for key, frequency := range extractedData.Frequency {
		insertFrequencyStatement += fmt.Sprintf("('%s', %d, (SELECT id FROM new_extracted_data)), ", key, frequency)
	}
	insertFrequencyStatement = insertFrequencyStatement[:len(insertFrequencyStatement)-2]
	sqlStatement := fmt.Sprintf(`
		WITH new_extracted_data AS (
			INSERT INTO extracted_data(data_id, line_count, word_count, char_count, average_word_length)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id
		) %s`, insertFrequencyStatement)

	// LogMessage(sqlStatement)
	result, err := r.DB.Exec(
		sqlStatement,
		extractedData.Id,
		extractedData.LineCount,
		extractedData.WordCount,
		extractedData.CharCount,
		extractedData.AverageWordLength,
	)
	CheckError(err, "Can't insert into database ")
	rowsAffected, err := result.RowsAffected()
	CheckError(err, "Can't get rows affected")
	LogMessage("Number of rows affected:", rowsAffected)

	return extractedData, nil
}
