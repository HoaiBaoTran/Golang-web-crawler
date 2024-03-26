package crawler

import (
	"context"
	"fmt"
	"log"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	chromeDp "github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
	"github.com/hoaibao/web-crawler/pkg/models"
	"github.com/hoaibao/web-crawler/pkg/repositories"
)

const (
	minPerLinkTest = 5 * time.Second
	minPerLink     = 1 * time.Minute
)

type MyCrawler struct {
	crawler *colly.Collector
	mutex   sync.Mutex
	visited map[string]bool
}

func CreateCrawler(maxDepth int) *MyCrawler {
	crawler := colly.NewCollector(
		colly.MaxDepth(maxDepth),
		colly.AllowedDomains("vnexpress.net", "localhost"),
	)

	return &MyCrawler{
		crawler: crawler,
		visited: make(map[string]bool),
	}
}

func (myCrawler *MyCrawler) CrawlWeb(urlLink string, depth, editDistance int, tag []string, exit chan bool, dataChan chan models.ExtractedData) {
	defer func() {
		exit <- true
	}()
	if depth <= 0 {
		fmt.Println("Exit")
		return
	}

	myCrawler.mutex.Lock()
	if _, ok := myCrawler.visited[urlLink]; !ok {
		myCrawler.visited[urlLink] = true
		myCrawler.mutex.Unlock()
	} else {
		myCrawler.mutex.Unlock()
		return
	}

	message := fmt.Sprintf("Crawling link: %s, depth: %d", urlLink, depth)
	repositories.MyLogger.LogMessage(message)

	extractedData := myCrawler.CrawlData(urlLink, tag)
	dataChan <- extractedData
	// Test data
	// sampleUrls := []string{
	// 	"http://localhost:8081/dong-vat",
	// 	"http://localhost:8081/nong-dan",
	// 	"http://localhost:8081/bong-da",
	// }
	//
	e := make(chan bool)
	for _, u := range extractedData.RelatedUrl {
		time.Sleep(minPerLinkTest)
		go myCrawler.CrawlWeb(u, depth-1, editDistance, tag, e, dataChan)
	}

	for i := 0; i < len(extractedData.RelatedUrl); i++ {
		<-e
	}

}

func (myCrawler *MyCrawler) CrawlData(urlLink string, tag []string) models.ExtractedData {
	extractedData := models.ExtractedData{}

	extractedData.CrawledUrl = urlLink
	relatedUrlChan := myCrawler.CrawlRelatedUrl(urlLink)

	myCrawler.crawler.OnHTML("body", func(e *colly.HTMLElement) {
		dataSource := e.Attr("data-source")
		dataSlice := strings.Split(dataSource, "-")
		id := dataSlice[len(dataSlice)-1]
		extractedData.Id = id
	})

	myCrawler.crawler.OnHTML(".top-detail > .container > .sidebar-1", func(e *colly.HTMLElement) {

		title := e.ChildText(".top-detail > .container > .sidebar-1 > .title-detail")
		if title != "" {
			if slices.Contains(tag, "h1") {
				title = fmt.Sprintf(`<h1>%s</h1>`, title)
			}
			extractedData.Title = title
		}

		lines := e.ChildTexts(".top-detail > .container > .sidebar-1 > p, .top-detail > .container > .sidebar-1 > article > p")
		for _, line := range lines {
			if slices.Contains(tag, "p") {
				line = fmt.Sprintf(`<p>%s</p>`, line)
			}
			extractedData.Paragraph = append(extractedData.Paragraph, line)
		}
		// if len(lines) > 0 {
		// 	if slices.Contains(tag, "p") {
		// 		extractedData.Paragraph = fmt.Sprintf(`<p>%s</p>`, )
		// 	} else {
		// 		extractedData.Paragraph = title
		// 	}
		// 	extractedData.Paragraph = append(extractedData.Paragraph, lines...)
		// }

		imgUrls := e.ChildAttrs(".top-detail > .container > .sidebar-1 > article > figure > .fig-picture > picture > img", "src")
		if len(imgUrls) > 0 {
			for _, value := range imgUrls {
				extractedData.Img = append(extractedData.Img, models.ImgStruct{
					Src: value,
				})
			}
		}

		imgDescriptions := e.ChildTexts(".top-detail > .container > .sidebar-1 > article > figure > figcaption > p")
		if len(imgDescriptions) > 0 {
			for index, value := range imgDescriptions {
				extractedData.Img[index].Description = value
			}
		}
	})

	myCrawler.crawler.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	myCrawler.crawler.Visit(urlLink)

	extractedData.RelatedUrl = <-relatedUrlChan
	return extractedData
}

func (myCrawler *MyCrawler) CrawlRelatedUrl(urlLink string) <-chan []string {
	relatedUrlChan := make(chan []string)
	go func(urlLink string) {
		opts := append(chromeDp.DefaultExecAllocatorOptions[:],
			// chromeDp.Flag("headless", false),
			chromeDp.Flag("headless", true),
			chromeDp.Flag("start-fullscreen", true),
			// chromeDp.Flag("enable-automation", false),
			// chromeDp.Flag("disable-extensions", false),
			// chromeDp.Flag("remote-debugging-port", "9222"),
		)

		// ctxTimeOut, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		// defer cancel()
		// allocatorCtx, cancel := chromeDp.NewExecAllocator(ctxTimeOut, opts...)
		// defer cancel()
		allocatorCtx, cancel := chromeDp.NewExecAllocator(context.Background(), opts...)
		defer cancel()
		ctx, cancel := chromeDp.NewContext(allocatorCtx, chromeDp.WithLogf(log.Printf))
		defer cancel()

		var elements []*cdp.Node
		relatedUrls := make([]string, 0)

		var found bool
		beforeTasks := chromeDp.Tasks{
			chromeDp.Navigate(urlLink),
			// chromeDp.Sleep(5 * time.Second),
			chromeDp.EvaluateAsDevTools(`document.getElementsByClassName("box-tinlienquanv2").length > 0 ? true : false`, &found),
		}
		err := chromeDp.Run(ctx, beforeTasks)
		if !found {
			fmt.Println("Not found element ", err)
			relatedUrlChan <- []string{}
			return
		} else {
			tasks := chromeDp.Tasks{
				chromeDp.Navigate(urlLink),
				chromeDp.WaitVisible(".top-detail > .container > .sidebar-1 > article.fck_detail > .box-tinlienquanv2 > article > h4 > a", chromeDp.ByQuery),
				chromeDp.Nodes(".top-detail > .container > .sidebar-1 > article.fck_detail > .box-tinlienquanv2 > article > h4 > a", &elements),
			}
			err := chromeDp.Run(ctx, tasks)
			if err != nil {
				fmt.Println("Not found element ", err)
				relatedUrlChan <- relatedUrls
				return
			}

			for _, element := range elements {
				var attributes map[string]string
				err = chromeDp.Run(ctx, chromeDp.Attributes(element.FullXPath(), &attributes))
				if err != nil {
					fmt.Println("Failed to get attributes:", err)
					relatedUrlChan <- relatedUrls
					return
				}

				for key, value := range attributes {
					if key == "href" {
						relatedUrls = append(relatedUrls, value)
					}
				}
			}
			relatedUrlChan <- relatedUrls
		}
	}(urlLink)
	return relatedUrlChan
}
