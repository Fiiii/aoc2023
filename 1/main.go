package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var counter int
var mapOfNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: 1")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
		return
	}

	scratcherFolderPath := fmt.Sprintf("%s/%s", dir, "1")
	inputFilePath := scratcherFolderPath + "/" + "input.txt"

	lines := make(chan string)
	processedLines := make(chan int)
	var wg sync.WaitGroup
	go readLines(inputFilePath, lines)

	workers := 10
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go processLines(lines, processedLines, &wg)
	}

	go func() {
		wg.Wait()
		close(processedLines)
	}()

	for num := range processedLines {
		counter += num
	}

	end := time.Now()
	fmt.Println("Counter:", counter)
	fmt.Println("Execution time: ", end.Sub(start))
}

func processLines(lines <-chan string, processedLines chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var num int
	var err error

	for line := range lines {
		strNumber := ""

		// get first number
		for i, v := range line {
			if byte(v) >= byte('0') && byte(v) <= byte('9') {
				strNumber += string(v)
				break
			} else {
				// check if it contains any of the words from numbers map
				for strNum, _ := range mapOfNumbers {
					if len(line[i:]) > len(strNum) {
						if strings.Contains(line[i:i+len(strNum)], strNum) == true {
							strNumber += strconv.Itoa(mapOfNumbers[strNum])
							break
						}
					}
				}
			}

			if len(strNumber) == 1 {
				break
			}
		}

		// get last number
		reversedline := reverseStr(line)
		for i, v := range reversedline {
			if byte(v) >= byte('0') && byte(v) <= byte('9') {
				strNumber += string(v)
				break
			} else {
				// check if it contains any of the words from numbers map
				for strNum, _ := range mapOfNumbers {
					if len(reversedline[i:]) > len(strNum) {
						rs := reverseStr(strNum)
						part := reversedline[i : i+len(rs)]
						if part == rs {
							strNumber += strconv.Itoa(mapOfNumbers[strNum])
							break
						}
					}
				}
			}

			if len(strNumber) == 2 {
				break
			}
		}

		if len(strNumber) > 1 {
			num, err = strconv.Atoi(string(strNumber[0]) + string(strNumber[len(strNumber)-1]))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			num, err = strconv.Atoi(string(strNumber[0]) + string(strNumber[0]))
			if err != nil {
				log.Fatal(err)
			}
		}

		processedLines <- num
	}
}

func readLines(filePath string, output chan<- string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		output <- line
	}

	close(output)
}

func reverseStr(str string) string {
	var result string

	for _, v := range str {
		result = string(v) + result
	}

	return result
}
