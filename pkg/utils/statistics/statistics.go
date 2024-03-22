package statistics

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error while reading file", err)
	}
}

func CountFrequencyAndCalcAverage(line []string, wg *sync.WaitGroup) (frequency map[string]int, averageWordLength float64) {
	frequency = make(map[string]int)
	regexPattern := `(\w+[.]\w+)|(\w+)`
	re := regexp.MustCompile(regexPattern)
	var totalWordLength int32
	for _, line := range line {
		words := strings.Fields(line)
		for _, word := range words {
			formattedWords := re.FindStringSubmatch(word)
			if len(formattedWords) < 1 {
				continue
			}
			if _, isContain := frequency[formattedWords[0]]; !isContain {
				totalWordLength += int32(len(formattedWords[0]))
			}
			frequency[formattedWords[0]]++
		}
	}
	wg.Done()
	averageWordLength = float64(len(frequency)) / float64(totalWordLength)
	return
}

func WordAndCharCount(line string) (wordCount int, charCount int) {
	wordSlice := strings.Fields(line)
	wordCount = len(wordSlice)

	lineWithoutSpace := strings.ReplaceAll(line, " ", "")
	charCount = len(lineWithoutSpace)

	return
}

func WriteResultToFile(lineCount, wordCount, charCount int32, averageWordLength float64, frequency map[string]int, executionTime time.Duration) {
	mapString := ""
	for key, value := range frequency {
		mapString += fmt.Sprintf("[%s: %d]\n", key, value)
	}

	// lines := []string{
	// 	fmt.Sprintf("Number of lines: %d\n", lineCount),
	// 	fmt.Sprintf("Number of words: %d\n", wordCount),
	// 	fmt.Sprintf("Number of characters: %d\n", charCount),
	// 	fmt.Sprintf("Average word length: %f\n", averageWordLength),
	// 	fmt.Sprintln("Frequency:"),
	// 	fmt.Sprintln(mapString),
	// }

	// current_date := time.Now().Format("02-01-2006")
	// current_time := time.Now().Format("15-04-05")
	// outputFileName := fmt.Sprintf("rs_%s_%s.txt", current_date, current_time)
	// outputFile, err := os.Create(outputFileName)

	// if err != nil {
	// 	fmt.Println("error while writing")
	// 	fmt.Println(err)
	// }

	// w := bufio.NewWriter(outputFile)

	// for _, line := range lines {
	// 	w.WriteString(line)
	// }

	// w.Flush()

	// fmt.Println("Write file successfully")
	// outputFile.Close()

	fmt.Println("Line Count:", lineCount)
	fmt.Println("Word Count:", wordCount)
	fmt.Println("Character Count:", charCount)
	fmt.Println("Average Word Length:", averageWordLength)
	// fmt.Println("Frequency:\n", mapString)
	fmt.Println("Execution Time:", executionTime)
}
