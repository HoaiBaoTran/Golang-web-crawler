package statistics

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type Result struct {
	lineCount, wordCount       int32
	charCount, totalWordLength int64
	frequency                  map[string]int
}

var (
	result Result = Result{
		frequency: make(map[string]int),
	}
	chunkNumber = 5
)

func merge(workerChannels ...<-chan Result) <-chan Result {
	var wg sync.WaitGroup
	out := make(chan Result)

	copyToOutput := func(c <-chan Result) {
		for item := range c {
			out <- item
		}
		wg.Done()
	}

	wg.Add(len(workerChannels))
	for _, c := range workerChannels {
		go copyToOutput(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func handleChunk(chunk []string) <-chan Result {
	chunkResultChan := make(chan Result)
	go func() {
		defer close(chunkResultChan)
		chunkResult := Result{}
		for _, line := range chunk {
			wordCount, charCount := WordAndCharCount(line)
			chunkResult.wordCount += int32(wordCount)
			chunkResult.charCount += int64(charCount)
		}
		chunkResultChan <- chunkResult

	}()
	return chunkResultChan
}

func splitFile(lines []string, chunkNumber int, wg *sync.WaitGroup) [][]string {
	chunks := make([][]string, chunkNumber)
	lengthPerChunks := len(lines) / chunkNumber

	for i := 0; i < chunkNumber; i++ {
		go func(i int) {
			var chunkLines []string
			if i == chunkNumber-1 {
				chunkLines = lines[i*lengthPerChunks:]
			} else {
				chunkLines = lines[i*lengthPerChunks : (i+1)*lengthPerChunks]
			}
			chunks[i] = chunkLines
			wg.Done()
		}(i)
	}

	return chunks
}

func Statistics(lines []string) (lineCount, wordCount int32, charCount int64, averageWordLength float64, frequency map[string]int) {
	var wg sync.WaitGroup

	wg.Add(chunkNumber + 1)
	go func() {
		result.frequency, result.totalWordLength = CountFrequencyAndCalcAverage(lines, &wg)
	}()

	chunks := splitFile(lines, chunkNumber, &wg)

	results := make([]<-chan Result, len(chunks))
	for index, chunk := range chunks {
		results[index] = handleChunk(chunk)
	}

	returnChan := merge(results...)
	for i := range returnChan {
		result.charCount += i.charCount
		result.wordCount += i.wordCount
	}

	wg.Wait()
	fmt.Println("Result", result)
	// lineCount = result.lineCount
	// wordCount = result.wordCount
	// charCount = result.charCount
	// averageWordLength = float64(result.totalWordLength) / float64(len(result.frequency))
	// frequency = result.frequency
	return
}

func CountFrequencyAndCalcAverage(line []string, wg *sync.WaitGroup) (frequency map[string]int, totalWordLength int64) {
	frequency = make(map[string]int)
	regexPattern := `(\w+[.]\w+)|(\w+)`
	re := regexp.MustCompile(regexPattern)
	for _, line := range line {
		words := strings.Fields(line)
		for _, word := range words {
			formattedWords := re.FindStringSubmatch(word)
			if len(formattedWords) < 1 {
				continue
			}
			if _, isContain := frequency[formattedWords[0]]; !isContain {
				totalWordLength += int64(len(formattedWords[0]))
			}
			frequency[formattedWords[0]]++
		}
	}
	wg.Done()
	return
}

func WordAndCharCount(line string) (wordCount int32, charCount int64) {
	wordSlice := strings.Fields(line)
	wordCount = int32(len(wordSlice))

	lineWithoutSpace := strings.ReplaceAll(line, " ", "")
	charCount = int64(len(lineWithoutSpace))

	return
}
