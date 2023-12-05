package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: 4")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
		return
	}

	scratcherFolderPath := fmt.Sprintf("%s/%s", dir, "4")
	inputFilePath := scratcherFolderPath + "/" + "input.txt"

	fileData, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println(part1(fileData))
	fmt.Println(part2(fileData))

	end := time.Now()
	fmt.Println("Execution time: ", end.Sub(start))
}

func part2(file []byte) int {
	var sumOfCards int
	cumulativeCards := make(map[int]int)
	cardNumber := 1

	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		matches := 0
		line := scanner.Text()
		strNums := line[8:]
		splitted := strings.Split(strNums, "|")
		winningNumbers := extractNumbers(splitted[0])
		nums := extractNumbers(splitted[1])

		// setup original
		cumulativeCards[cardNumber] += 1

		// validate copies
		for _, num := range nums {
			for _, winningNum := range winningNumbers {
				if num == winningNum {
					matches++
				}
			}
		}

		if matches > 0 {
			// loop on matches
			for j := 0; j < cumulativeCards[cardNumber]; j++ {
				for i := 1; i <= matches; i++ {
					cumulativeCards[cardNumber+i] += 1
				}
			}
		}

		cardNumber++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, v := range cumulativeCards {
		sumOfCards += v
	}

	return sumOfCards
}

func part1(file []byte) int {
	var counter int
	scanner := bufio.NewScanner(bytes.NewReader(file))
	for scanner.Scan() {
		matches := 0
		line := scanner.Text()
		strNums := line[8:]
		splitted := strings.Split(strNums, "|")
		winningNumbers := extractNumbers(splitted[0])
		nums := extractNumbers(splitted[1])
		for _, num := range nums {
			for _, winningNum := range winningNumbers {
				if num == winningNum {
					if matches > 0 {
						matches *= 2
					} else {
						matches += 1
					}
				}
			}
		}

		counter += matches
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return counter
}

func extractNumbers(input string) []string {
	re := regexp.MustCompile(`\b\d+\b`)
	matches := re.FindAllString(input, -1)

	return matches
}
