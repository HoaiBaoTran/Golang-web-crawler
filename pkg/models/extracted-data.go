package models

type ExtractedData struct {
	Title     string   `json:"title"`
	Heading   []string `json:"heading"`
	Paragraph []string `json:"paragraph"`
}
