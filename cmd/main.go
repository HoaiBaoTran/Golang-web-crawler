package main

import (
	"fmt"

	"github.com/hoaibao/web-crawler/pkg/utils/crawler"
)

func main() {
	myCrawler := crawler.CreateCrawler()
	// urlLink := "https://vnexpress.net/nong-dan-thu-nhap-gap-10-lan-neu-san-xuat-tom-lua-quy-mo-lon-4725097.html"
	urlLink := "http://localhost:8080/nong-dan"
	extractedData := myCrawler.CrawlData(urlLink)
	fmt.Println("data: ", extractedData)
}
