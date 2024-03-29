package levenshteindistancetest_test

import (
	"testing"

	levenshteinDistance "github.com/hoaibao/web-crawler/pkg/utils/levenshtein-distance"
)

func TestLevenshteinDistance(t *testing.T) {
	str1 := "Capybara"
	str2 := "Lion"

	result := levenshteinDistance.LevenshteinDistance(str1, str2)

	expectedNumber := 8
	if expectedNumber != result {
		t.Errorf("unexpected word count: got %d, want %d", result, expectedNumber)
	}

}
