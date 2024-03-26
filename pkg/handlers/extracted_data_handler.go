package handlers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/services"
)

type Options struct {
	MaxDepth     int      `json:"max-depth"`
	Tag          []string `json:"tag"`
	EditDistance int      `json:"edit-distance"`
}

type RequestBody struct {
	Url     string  `json:"url"`
	Options Options `json:"options"`
}

type ExtractedDataHandler struct {
	ExtractedDataService services.ExtractedDataService
}

func NewExtractedDataHandler(extractedDataService *services.ExtractedDataService) *ExtractedDataHandler {
	return &ExtractedDataHandler{
		ExtractedDataService: *extractedDataService,
	}
}

func (h *ExtractedDataHandler) GetAllExtractedData(w http.ResponseWriter, r *http.Request) {

}

func (h *ExtractedDataHandler) GetExtractedDataById(w http.ResponseWriter, r *http.Request) {

}

func (h *ExtractedDataHandler) CreateExtractedData(w http.ResponseWriter, r *http.Request) {
	var data RequestBody
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errMessage := fmt.Sprintf("Invalid request body %s", err)
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}
	urlLink := data.Url
	if urlLink == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	options := data.Options
	if options.MaxDepth <= 0 {
		options.MaxDepth = 1
	}
	if options.EditDistance <= 0 {
		options.EditDistance = -1
	}

	responseData, err := h.ExtractedDataService.CreateExtractedData(
		urlLink,
		options.MaxDepth,
		options.EditDistance,
		options.Tag,
	)

	if err != nil {
		http.Error(w, "Server", http.StatusBadRequest)
		return
	}

	for _, extractedData := range responseData {

		file, errMessage, err := h.WriteJsonFile(extractedData)
		if err != nil {
			fmt.Println("err: ", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name()))
		w.WriteHeader(http.StatusOK)

		http.ServeFile(w, r, file.Name())
	}
}

func (h *ExtractedDataHandler) WriteJsonFile(extractedData models.ExtractedData) (file *os.File, errMessage string, err error) {
	current_date := time.Now().Format("02-01-2006")
	current_time := time.Now().Format("15-04-05")
	outputFileName := fmt.Sprintf("json-files/%s_%s_%s.json", extractedData.Id, current_date, current_time)

	file, err = os.Create(outputFileName)
	if err != nil {
		errMessage = "Error opening JSON file"
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	jsonData, err := json.Marshal(extractedData)
	if err != nil {
		errMessage = "Error parsing JSON"
		return
	}

	_, err = writer.Write(jsonData)
	if err != nil {
		errMessage = "Error writing JSON file"
		return
	}

	err = writer.Flush()
	if err != nil {
		errMessage = "Error writing JSON file"
		return
	}

	return
}
