package convertJSON

import (
	"fmt"
	"strings"

	"github.com/hoaibao/web-crawler/pkg/models"
)

func GetJsonStringFromData(data models.ExtractedData) string {
	jsonString := fmt.Sprintf(`{
	"id": "%s",
	"title": "%s",
	"paragraphs": %s,
	"crawled-url": "%s",
	"related-url": %s,
	"img": %s,
	"line-count": %d,
	"word-count": %d,
	"character-count": %d,
	"average-word-length": %f,
	"frequency": %s
}`,
		data.Id,
		data.Title,
		GetJsonStringFromSliceString(data.Paragraph),
		data.CrawledUrl,
		GetJsonStringFromSliceString(data.RelatedUrl),
		GetJsonStringFromImgSlice(data.Img),
		data.LineCount,
		data.WordCount,
		data.CharCount,
		data.AverageWordLength,
		GetJsonStringFromMapFrequency(data.Frequency),
	)
	return jsonString
}

func GetJsonStringFromSliceString(data []string) string {
	jsonString := "[\n"
	for index, line := range data {
		line = strings.ReplaceAll(line, `"`, `\"`)
		jsonString += fmt.Sprintf("\t\t\"%s\"", line)
		if index < len(data)-1 {
			jsonString += ",\n"
		}
	}
	jsonString += "\n\t]"
	return jsonString
}

func GetJsonStringFromImgSlice(imgSlice []models.ImgStruct) string {
	jsonString := "[\n\t\t"
	for index, img := range imgSlice {
		imgJsonString := fmt.Sprintf(`{
			"img-src": "%s",
			"img-description": "%s"
		}`, img.Src, img.Description)

		if index < len(imgSlice) {
			imgJsonString += ",\n\t\t"
		}
		jsonString += imgJsonString
	}
	jsonString = jsonString[:len(jsonString)-4]
	jsonString += "\n\t]"
	return jsonString
}

func GetJsonStringFromMapFrequency(frequency map[string]int) string {
	jsonString := "{\n"

	for key, value := range frequency {
		jsonString += fmt.Sprintf("\t\t\"%s\": %d,\n", key, value)
	}
	jsonString = jsonString[:len(jsonString)-2]
	jsonString += "\n\t}"
	return jsonString
}
