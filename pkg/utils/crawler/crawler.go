package crawler

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/hoaibao/web-crawler/pkg/models"
)

type MyCrawler struct {
	crawler *colly.Collector
}

func CreateCrawler() *MyCrawler {
	crawler := colly.NewCollector(
		colly.MaxDepth(1),
		// colly.AllowedDomains("vnexpress.net"),
	)

	return &MyCrawler{
		crawler: crawler,
	}
}

func (myCrawler *MyCrawler) CrawlData(urlLink string) models.ExtractedData {
	extractedData := models.ExtractedData{}
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

		imgUrls := e.ChildAttrs(".container > .sidebar-1 > article > figure > .fig-picture > picture > img", "src")
		if len(imgUrls) > 0 {
			for _, value := range imgUrls {
				extractedData.Img = append(extractedData.Img, models.ImgStruct{
					Src: value,
				})
			}
		}

		imgDescriptions := e.ChildTexts(".container > .sidebar-1 > article > figure > figcaption > p")
		if len(imgDescriptions) > 0 {
			for index, value := range imgDescriptions {
				extractedData.Img[index].Description = value
			}
		}
	})

	myCrawler.crawler.OnHTML("a[href]", func(e *colly.HTMLElement) {
		urlLink := e.Attr("href")
		regexPattern := `(https:\/\/\S+)`
		re := regexp.MustCompile(regexPattern)
		if re.MatchString(urlLink) {
			extractedData.Url = append(extractedData.Url, urlLink)
		}
	})

	myCrawler.crawler.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	myCrawler.crawler.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	myCrawler.crawler.Visit(urlLink)

	return extractedData
}
