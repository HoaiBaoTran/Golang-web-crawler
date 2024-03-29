package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/services"
	handleTag "github.com/hoaibao/web-crawler/pkg/utils/handle-html-tag"
)

type Options struct {
	MaxDepth            int      `json:"max-depth"`
	Tag                 []string `json:"tag"`
	LevenshteinDistance int      `json:"levenshtein-distance"`
	Word                []string `json:"word"`
	WrappedTag          string   `json:"wrapped-tag"`
}

type RequestBody struct {
	Url     []string `json:"url"`
	Options Options  `json:"options"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Data models.ExtractedData `json:"data"`
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
	vars := mux.Vars(r)
	id, errId := vars["id"]
	if !errId {
		http.Error(w, "Invalid paragraph id", http.StatusBadRequest)
		return
	}

	fileName := fmt.Sprintf("%s.json", id)
	file, err := os.Open("json-files/" + fileName)
	if err != nil {
		http.Error(w, "Error retrieving extracted data", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/octet-stream")

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error copying file to response", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "File is ready",
	})
}

func (h *ExtractedDataHandler) CreateExtractedData(w http.ResponseWriter, r *http.Request) {
	var data RequestBody
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		errMessage := fmt.Sprintf("Invalid request body %s", err)
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}
	urlLink := data.Url
	if len(urlLink) <= 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	options := data.Options
	if options.MaxDepth <= 0 {
		errMessage := "Invalid max depth"
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	if options.LevenshteinDistance < 0 {
		errMessage := "Invalid levenshtein distance"
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	if !handleTag.IsValidHtmlTag(options.WrappedTag) {
		errMessage := "Invalid wrapped tag"
		http.Error(w, errMessage, http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Message: "Crawling data, wait a moment",
	})

	go func() {
		_, err := h.ExtractedDataService.CreateExtractedData(
			options.WrappedTag,
			options.MaxDepth,
			options.LevenshteinDistance,
			urlLink,
			options.Tag,
			options.Word,
		)
		if err != nil {
			http.Error(w, "Server", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(Response{
			Status:  http.StatusOK,
			Message: "Crawling data successfully",
		})
		// for _, extractedData := range responseData {
		// 	if err != nil {
		// 		fmt.Println("err: ", err)
		// 		http.Error(w, errMessage, http.StatusInternalServerError)
		// 		return
		// 	}
		// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name()))
		// }
	}()
}
