package models

type ExtractedData struct {
	Title     string      `json:"title"`
	Heading   []string    `json:"heading"`
	Paragraph []string    `json:"paragraph"`
	Img       []ImgStruct `json:"image"`
	Url       []string    `json:"url"`
}

type ImgStruct struct {
	Src         string `json:"imgSrc"`
	Description string `json:"imgDescription"`
}
