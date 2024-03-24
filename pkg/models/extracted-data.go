package models

type ExtractedData struct {
	Title             string         `json:"title"`
	Paragraph         []string       `json:"paragraph"`
	CrawledUrl        string         `json:"crawled-url"`
	RelatedUrl        []string       `json:"related-url"`
	Img               []ImgStruct    `json:"image"`
	LineCount         int            `json:"line-count"`
	WordCount         int            `json:"word-count"`
	CharCount         int            `json:"character-count"`
	AverageWordLength float64        `json:"average-word-length"`
	Frequency         map[string]int `json:"frequency"`
}

type ImgStruct struct {
	Src         string `json:"imgSrc"`
	Description string `json:"imgDescription"`
}
