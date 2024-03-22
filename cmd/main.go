package main

import (
	"fmt"

	"github.com/hoaibao/web-crawler/pkg/handlers"
	"github.com/hoaibao/web-crawler/pkg/repositories"
	"github.com/hoaibao/web-crawler/pkg/services"
)

func main() {
	// myCrawler := crawler.CreateCrawler()
	// urlLink := "https://vnexpress.net/nong-dan-thu-nhap-gap-10-lan-neu-san-xuat-tom-lua-quy-mo-lon-4725097.html"
	// urlLink := "http://localhost:8080/nong-dan"
	// extractedData := myCrawler.CrawlData(urlLink)
	// fmt.Println("data: ", extractedData)

	extractedDataRepository := repositories.NewMemoryExtractedDataRepository()
	extractedDataService := services.NewExtractedDataService(extractedDataRepository)
	extractedDataHandler := handlers.NewExtractedDataHandler(extractedDataService)
	fmt.Println("Test ", extractedDataHandler)
}
