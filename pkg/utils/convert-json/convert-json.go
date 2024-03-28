package convertJSON

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

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
	if len(data) <= 0 {
		return "[]"
	}
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
	if len(imgSlice) <= 0 {
		return "[]"
	}
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
	if len(frequency) <= 0 {
		return "{}"
	}
	jsonString := "{\n"

	for key, value := range frequency {
		jsonString += fmt.Sprintf("\t\t\"%s\": %d,\n", key, value)
	}
	jsonString = jsonString[:len(jsonString)-2]
	jsonString += "\n\t}"
	return jsonString
}

func WriteJsonFile(extractedData models.ExtractedData) (file *os.File, errMessage string, err error) {
	current_date := time.Now().Format("02-01-2006")
	current_time := time.Now().Format("15-04-05")
	outputFileName := fmt.Sprintf("json-files/%s_%s_%s.json", extractedData.Id, current_date, current_time)

	file, err = os.Create(outputFileName)
	if err != nil {
		errMessage = "Error opening JSON file"
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	jsonString := GetJsonStringFromData(extractedData)
	writer.Write([]byte(jsonString))

	return
}
