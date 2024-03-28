package handlehtmltag

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	leDistance "github.com/hoaibao/web-crawler/pkg/utils/levenshtein-distance"
	"golang.org/x/net/html"
)

func IsValidHtmlTag(tag string) bool {
	_, err := html.Parse(strings.NewReader(tag))
	return err == nil
}

func ChangeContentToHtmlTag(paragraph, wordsReplace []string, wrappedTag string, levenshteinDistance int) []string {
	newParagraph := make([]string, len(paragraph))
	paraChannels := make([]chan string, len(paragraph))
	for index, smallPara := range paragraph {
		paraChannels[index] = ChangeContentSmallParagraph(smallPara, wrappedTag, wordsReplace, levenshteinDistance)
	}

	for i := range len(paragraph) {
		newParagraph[i] = <-paraChannels[i]
	}

	return newParagraph
}

func ChangeContentSmallParagraph(paragraph, wrappedTag string, wordsReplace []string, levenshteinDistance int) chan string {
	paraChan := make(chan string)
	go func(paragraph, wrappedTag string, wordsReplace []string, levenshteinDistance int) {

		tagSlice := strings.Split(wrappedTag, ">")
		tagSlice[0] += ">"
		tagSlice[1] += ">"
		var wordsNeedReplace []string

		punctuationPattern := regexp.MustCompile(`[[:punct:]]`)
		cleanedString := punctuationPattern.ReplaceAllString(paragraph, " ")
		wordSlice := strings.Split(cleanedString, " ")

		for _, wordPara := range wordSlice {
			for _, word := range wordsReplace {
				if leDistance.LevenshteinDistance(wordPara, word) <= levenshteinDistance &&
					!slices.Contains(wordsNeedReplace, wordPara) {
					wordsNeedReplace = append(wordsNeedReplace, wordPara)
				}
			}
		}

		for _, word := range wordsNeedReplace {
			newWord := fmt.Sprintf("%s%s%s", tagSlice[0], word, tagSlice[1])
			paragraph = strings.ReplaceAll(paragraph, word, newWord)
		}
		paraChan <- paragraph
	}(paragraph, wrappedTag, wordsReplace, levenshteinDistance)
	return paraChan
}
