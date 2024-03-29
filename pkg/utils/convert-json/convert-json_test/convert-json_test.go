package convertjson_test

import (
	"testing"

	"github.com/hoaibao/web-crawler/pkg/models"
	convertJSON "github.com/hoaibao/web-crawler/pkg/utils/convert-json"
)

func TestGetJsonStringFromMapFrequency(t *testing.T) {
	frequency := map[string]int{
		"this": 1,
		"is":   2,
	}
	result := convertJSON.GetJsonStringFromMapFrequency(frequency)
	expectedResult := `{
		"this": 1,
		"is": 2
	}`
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func TestGetJsonStringFromImgSlice(t *testing.T) {
	imgSlice := []models.ImgStruct{
		models.ImgStruct{
			Src:         "img-src1",
			Description: "img-description-1",
		},
		models.ImgStruct{
			Src:         "img-src2",
			Description: "img-description-2",
		},
	}

	result := convertJSON.GetJsonStringFromImgSlice(imgSlice)
	expectedResult := `[
		{
			"img-src": "img-src1",
			"img-description": "img-description-1"
		},
		{
			"img-src": "img-src2",
			"img-description": "img-description-2"
		}
	]`
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func TestGetJsonStringFromSliceString(t *testing.T) {
	stringSlice := []string{
		"This is line 1",
		"This is line 2",
	}

	result := convertJSON.GetJsonStringFromSliceString(stringSlice)
	expectedResult := `[
		"This is line 1",
		"This is line 2"
	]`
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func TestGetJsonStringFromData(t *testing.T) {
	data := models.ExtractedData{
		Id:         "1",
		Title:      "title",
		Paragraph:  []string{"line 1", "line 2"},
		CrawledUrl: "crawled-url",
		RelatedUrl: []string{"related-1", "related-2"},
		Img: []models.ImgStruct{
			{
				Src:         "img-src1",
				Description: "img-description-1",
			},
			{
				Src:         "img-src2",
				Description: "img-description-2",
			},
		},
		LineCount:         8,
		WordCount:         10,
		CharCount:         11,
		AverageWordLength: 12,
		Frequency: map[string]int{
			"this": 1,
			"is":   2,
		},
	}

	result := convertJSON.GetJsonStringFromData(data)
	expectedResult := `{
	"id": "1",
	"title": "title",
	"paragraphs": [
		"line 1",
		"line 2"
	],
	"crawled-url": "crawled-url",
	"related-url": [
		"related-1",
		"related-2"
	],
	"img": [
		{
			"img-src": "img-src1",
			"img-description": "img-description-1"
		},
		{
			"img-src": "img-src2",
			"img-description": "img-description-2"
		}
	],
	"line-count": 8,
	"word-count": 10,
	"character-count": 11,
	"average-word-length": 12.000000,
	"frequency": {
		"this": 1,
		"is": 2
	}
}`
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}
