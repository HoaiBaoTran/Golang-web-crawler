package levenshteindistancetest_test

import (
	"testing"

	handlehtmltag "github.com/hoaibao/web-crawler/pkg/utils/handle-html-tag"
)

func TestIsValidHtmlTag(t *testing.T) {
	str1 := `<b class='test-class'></b>`

	result := handlehtmltag.IsValidHtmlTag(str1)

	expectedResult := true
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}

}

func TestChangeContentToHtmlTag(t *testing.T) {
	paragraph := []string{
		"This is test line 1",
		"This is test line 2",
	}

	expectedResult := []string{
		"Th<b>is</b> <b>is</b> test line 1",
		"Th<b>is</b> <b>is</b> test line 2",
	}

	result := handlehtmltag.ChangeContentToHtmlTag(paragraph, []string{"is"}, "<b></b>", 0)

	if !slicesEqual(expectedResult, result) {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func TestChangeContentSmallParagraph(t *testing.T) {
	paragraph := "This is test line 1"

	expectedResult := "Th<b>is</b> <b>is</b> test line 1"

	resultChan := handlehtmltag.ChangeContentSmallParagraph(paragraph, "<b></b>", []string{"is"}, 0)
	result := <-resultChan
	if expectedResult != result {
		t.Errorf("unexpected word count: got %v, want %v", result, expectedResult)
	}
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
