package routers

import (
	"github.com/gorilla/mux"
	"github.com/hoaibao/web-crawler/pkg/handlers"
)

func SetMainRouter() *mux.Router {
	return mux.NewRouter()
}

func SetExtractedDataRouter(extractedDataHandler *handlers.ExtractedDataHandler, mainRouter *mux.Router) {
	extractedDataRouter := mainRouter.PathPrefix("/api/v1/extracted-data").Subrouter()

	extractedDataRouter.HandleFunc("", extractedDataHandler.GetAllExtractedData).Methods("GET")
	extractedDataRouter.HandleFunc("/{id}", extractedDataHandler.GetExtractedDataById).Methods("GET")
	extractedDataRouter.HandleFunc("", extractedDataHandler.GetAllExtractedData).Methods("POST")
}
