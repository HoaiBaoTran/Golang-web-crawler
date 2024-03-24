package statistics

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	result Result = Result{
		frequency: make(map[string]int),
	}
	chunkNumber = 5
)

type Result struct {
	lineCount, wordCount, charCount, totalWordLength int
	frequency                                        map[string]int
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error while reading file", err)
	}
}

func CountFrequencyFromLine(data string) <-chan map[string]int {
	frequencyChan := make(chan map[string]int)
	go func() {
		frequency := make(map[string]int)
		totalWordLength := 0
		punctuationPattern := regexp.MustCompile(`[[:punct:]]`)
		cleanedString := punctuationPattern.ReplaceAllString(data, " ")
		words := strings.Fields(cleanedString)
		for _, word := range words {
			if _, isContain := frequency[word]; !isContain {
				totalWordLength += len(word)
			}
			frequency[word]++
		}
		frequencyChan <- frequency
	}()

	return frequencyChan
}

func WordAndCharCount(line string) (int, int) {
	wordSlice := strings.Fields(line)
	wordCount := len(wordSlice)

	lineWithoutSpace := strings.ReplaceAll(line, " ", "")
	charCount := len(lineWithoutSpace)

	return wordCount, charCount
}

func WriteResultToFile(lineCount, wordCount, charCount int, averageWordLength float64, frequency map[string]int, executionTime time.Duration) {
	mapString := ""
	for key, value := range frequency {
		mapString += fmt.Sprintf("[%s: %d]\n", key, value)
	}

	lines := []string{
		fmt.Sprintf("Number of lines: %d\n", lineCount),
		fmt.Sprintf("Number of words: %d\n", wordCount),
		fmt.Sprintf("Number of characters: %d\n", charCount),
		fmt.Sprintf("Average word length: %f\n", averageWordLength),
		fmt.Sprintln("Frequency:"),
		fmt.Sprintln(mapString),
	}

	current_date := time.Now().Format("02-01-2006")
	current_time := time.Now().Format("15-04-05")
	outputFileName := fmt.Sprintf("results/rs_%s_%s.txt", current_date, current_time)
	outputFile, err := os.Create(outputFileName)

	if err != nil {
		fmt.Println("error while writing")
		fmt.Println(err)
	}

	w := bufio.NewWriter(outputFile)

	for _, line := range lines {
		w.WriteString(line)
	}

	fmt.Printf("Number of lines: %d\n", lineCount)
	fmt.Printf("Number of words: %d\n", wordCount)
	fmt.Printf("Number of characters: %d\n", charCount)
	fmt.Printf("Average word length: %f\n", averageWordLength)
	fmt.Println("Frequency:")
	// fmt.Println(result.frequency)
	// fmt.Println(mapString)
	fmt.Println("Execution Time:", executionTime)

	w.WriteString(fmt.Sprintln("Execution Time", executionTime))
	w.Flush()

	fmt.Println("Write file successfully")
	outputFile.Close()
}

func SplitFile(lines []string, chunkNumber int, wg *sync.WaitGroup) [][]string {
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

func HandleChunk(chunk []string) <-chan Result {
	chunkResultChan := make(chan Result)
	go func() {
		frequencyChan := CountFrequencyFromLine(strings.Join(chunk, " "))
		defer close(chunkResultChan)
		chunkResult := Result{}
		for _, line := range chunk {
			wordCount, charCount := WordAndCharCount(line)
			chunkResult.wordCount += wordCount
			chunkResult.charCount += charCount
		}
		chunkResult.frequency = <-frequencyChan
		chunkResultChan <- chunkResult
	}()
	return chunkResultChan
}

func Merge(workerChannels ...<-chan Result) <-chan Result {
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

func Statistics(paragraph []string) (lineCount, wordCount, charCount int, averageWordLength float64, frequency map[string]int) {
	var wg sync.WaitGroup
	result.lineCount = len(paragraph)

	wg.Add(chunkNumber)
	chunks := SplitFile(paragraph, chunkNumber, &wg)
	wg.Wait()

	results := make([]<-chan Result, len(chunks))
	for index, chunk := range chunks {
		results[index] = HandleChunk(chunk)
	}

	returnChan := Merge(results...)
	for i := range returnChan {
		result.charCount += i.charCount
		result.wordCount += i.wordCount
		for key, value := range i.frequency {
			if _, isKey := result.frequency[key]; !isKey {
				result.totalWordLength += len(key)
				result.frequency[key] = value
			} else {
				result.frequency[key] += value
			}
		}
	}

	lineCount = result.lineCount
	wordCount = result.wordCount
	charCount = result.charCount
	averageWordLength = float64(result.totalWordLength) / float64(len(result.frequency))
	frequency = result.frequency

	return
}
