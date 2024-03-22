package main

import (
	"fmt"
	"net/http"

	"github.com/hoaibao/web-crawler/pkg/handlers"
	"github.com/hoaibao/web-crawler/pkg/repositories"
	"github.com/hoaibao/web-crawler/pkg/routers"
	"github.com/hoaibao/web-crawler/pkg/services"
)

func main() {

	extractedDataRepository := repositories.NewMemoryExtractedDataRepository()
	extractedDataService := services.NewExtractedDataService(extractedDataRepository)
	extractedDataHandler := handlers.NewExtractedDataHandler(extractedDataService)

	mainRouter := routers.SetMainRouter()
	routers.SetExtractedDataRouter(extractedDataHandler, mainRouter)

	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", mainRouter)
}
