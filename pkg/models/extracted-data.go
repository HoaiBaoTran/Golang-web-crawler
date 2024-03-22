package models

type ExtractedData struct {
	Title      string      `json:"title"`
	Paragraph  []string    `json:"paragraph"`
	CrawledUrl string      `json:"crawled-url"`
	RelatedUrl []string    `json:"related-url"`
	Img        []ImgStruct `json:"image"`
}

type ImgStruct struct {
	Src         string `json:"imgSrc"`
	Description string `json:"imgDescription"`
}
