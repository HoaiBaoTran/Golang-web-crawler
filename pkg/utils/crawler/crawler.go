package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/hoaibao/web-crawler/pkg/models"
)

type MyCrawler struct {
	crawler *colly.Collector
}

var (
	extractedData = models.ExtractedData{}
)

func CreateCrawler() *MyCrawler {
	crawler := colly.NewCollector(
		colly.MaxDepth(1),
		// colly.AllowedDomains("vnexpress.net"),
	)

	return &MyCrawler{
		crawler: crawler,
	}
}

func (myCrawler *MyCrawler) CrawlData(urlLink string) {
	myCrawler.crawler.OnHTML(".container > .sidebar-1", func(e *colly.HTMLElement) {
		title := e.ChildText(".container > .sidebar-1 > .title-detail")
		if title != "" {
			extractedData.Title = title
			fmt.Println("title", title)
		}

		lines := e.ChildTexts(".container > .sidebar-1 > p, .container > .sidebar-1 > article > p")
		if len(lines) > 0 {
			extractedData.Paragraph = append(extractedData.Paragraph, lines...)
		}
	})

	myCrawler.crawler.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	myCrawler.crawler.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	myCrawler.crawler.Visit(urlLink)

	fmt.Println("result: ", extractedData)
}
