package statistics_test

import (
	"reflect"
	"sync"
	"testing"

	"github.com/hoaibao/web-crawler/pkg/utils/statistics"
)

func BenchmarkCountWordAndChar(b *testing.B) {
	content := "This is a sample line."

	wordCount, charCount := statistics.WordAndCharCount(content)

	expectedWordCount := 5
	expectedCharCount := 18
	if wordCount != expectedWordCount {
		b.Errorf("unexpected word count: got %d, want %d", wordCount, expectedWordCount)
	}
	if charCount != expectedCharCount {
		b.Errorf("unexpected char count: got %d, want %d", charCount, expectedCharCount)
	}
}

func BenchmarkFrequency(b *testing.B) {
	var wg sync.WaitGroup
	content := []string{
		"This is sample line 1.",
		"This is sample line 2.",
	}
	wg.Add(1)
	frequency, totalUniqueWordLength := statistics.CountFrequencyAndCalcAverage(content, &wg)
	wg.Wait()

	expectedFrequency := map[string]int{
		"This":   2,
		"is":     2,
		"sample": 2,
		"line":   2,
		"1":      1,
		"2":      1,
	}
	expectedTotalUniqueWordLength := 18

	if totalUniqueWordLength != int64(expectedTotalUniqueWordLength) {
		b.Errorf("unexpected total word length: got %d, want %d", totalUniqueWordLength, expectedTotalUniqueWordLength)
	}
	if !reflect.DeepEqual(frequency, expectedFrequency) {
		b.Errorf("Result %v does not match expected %v", frequency, expectedFrequency)
	}
}
