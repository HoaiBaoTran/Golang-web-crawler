package models

type ExtractedData struct {
	Title             string         `json:"title"`
	Paragraph         []string       `json:"paragraph"`
	CrawledUrl        string         `json:"crawled-url"`
	RelatedUrl        []string       `json:"related-url"`
	Img               []ImgStruct    `json:"image"`
	LineCount         int32          `json:"line-count"`
	WordCount         int32          `json:"word-count"`
	CharCount         int64          `json:"character-count"`
	AverageWordLength float64        `json:"average-word-length"`
	Frequency         map[string]int `json:"frequency"`
}

type ImgStruct struct {
	Src         string `json:"imgSrc"`
	Description string `json:"imgDescription"`
}
