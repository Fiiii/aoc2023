package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

var counter int

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: 1")
	file := "/Users/arturfigiel/go/aoc2023/1/input.txt"

	lines := make(chan string)
	processedLines := make(chan int)
	var wg sync.WaitGroup
	go readLines(file, lines)

	workers := 20
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go processLines(i, lines, processedLines, &wg)
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

func processLines(workerID int, lines <-chan string, processedLines chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	var num int
	var err error

	fmt.Println("Worker", workerID, "started")
	for line := range lines {
		strNumber := ""

		for _, v := range line {
			if byte(v) >= byte('0') && byte(v) <= byte('9') {
				strNumber += string(v)
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
