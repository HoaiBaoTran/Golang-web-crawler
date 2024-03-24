package crawler

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	chromeDp "github.com/chromedp/chromedp"
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
	extractedData := models.ExtractedData{
		RelatedUrl: make([]string, 0),
	}
	extractedData.CrawledUrl = urlLink
	myCrawler.crawler.OnHTML(".container > .sidebar-1", func(e *colly.HTMLElement) {

		title := e.ChildText(".container > .sidebar-1 > .title-detail")
		if title != "" {
			extractedData.Title = title
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

		relatedUrls := e.ChildAttrs(".container > .sidebar-1 > article > div.box-tinlienquanv2 > article > h4 > a", "href")
		if len(relatedUrls) > 0 {
			extractedData.RelatedUrl = relatedUrls
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

func (myCrawler *MyCrawler) CrawlRelatedUrl(urlLink string) <-chan []string {
	relatedUrlChan := make(chan []string)
	go func(urlLink string) {
		opts := append(chromeDp.DefaultExecAllocatorOptions[:],
			chromeDp.Flag("headless", false),
			chromeDp.Flag("start-fullscreen", true),
			// chromeDp.Flag("enable-automation", false),
			// chromeDp.Flag("disable-extensions", false),
			// chromeDp.Flag("remote-debugging-port", "9222"),
		)

		ctxTimeOut, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		allocatorCtx, cancel := chromeDp.NewExecAllocator(ctxTimeOut, opts...)
		defer cancel()
		ctx, cancel := chromeDp.NewContext(allocatorCtx, chromeDp.WithLogf(log.Printf))
		defer cancel()

		var elements []*cdp.Node
		var relatedUrls []string
		tasks := chromeDp.Tasks{
			chromeDp.Navigate(urlLink),
			chromeDp.WaitVisible(".container > .sidebar-1 > article.fck_detail > .box-tinlienquanv2", chromeDp.ByQuery),
			chromeDp.Nodes(".container > .sidebar-1 > article.fck_detail > .box-tinlienquanv2 > article > h4 > a", &elements),
		}

		err := chromeDp.Run(ctx, tasks)
		if err != nil {
			log.Fatal("err: ", err)
		}

		for _, element := range elements {
			var attributes map[string]string
			err = chromeDp.Run(ctx, chromeDp.Attributes(element.FullXPath(), &attributes))
			if err != nil {
				fmt.Println("Failed to get attributes:", err)
				return
			}

			for key, value := range attributes {
				if key == "href" {
					relatedUrls = append(relatedUrls, value)
				}
			}
		}
		relatedUrlChan <- relatedUrls
	}(urlLink)
	return relatedUrlChan
}
