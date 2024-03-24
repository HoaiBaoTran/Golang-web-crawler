package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hoaibao/web-crawler/pkg/services"
)

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
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errMessage := fmt.Sprintf("Invalid request body %s", err)
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}
	urlPath, isUrlPathExists := data["url"]
	if !isUrlPathExists {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	fmt.Println("Crawling data ....")
	response, err := h.ExtractedDataService.CreateExtractedData(urlPath)
	if err != nil {
		http.Error(w, "Server", http.StatusBadRequest)
		return
	}

	fmt.Println("Crawling data successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
