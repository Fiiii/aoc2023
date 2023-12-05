package main

import (
	"aoc2023/utils"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var counter int

// PART 1
//var maxRed = 12
//var maxGreen = 13
//var maxBlue = 14

// PART 1
//var bagLimits = map[string]int{
//	"red":   maxRed,
//	"green": maxGreen,
//	"blue":  maxBlue,
//}

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: 1")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
		return
	}

	scratcherFolderPath := fmt.Sprintf("%s/%s", dir, "2")
	inputFilePath := scratcherFolderPath + "/" + "input.txt"

	lines := make(chan string)
	processedLines := make(chan int)
	var wg sync.WaitGroup
	go utils.ReadLines(inputFilePath, lines)

	workers := 100
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go processLines(lines, processedLines, &wg)
	}

	go func() {
		wg.Wait()
		close(processedLines)
	}()

	for num := range processedLines {
		fmt.Println("Processed game", num)
		counter += num
	}

	end := time.Now()
	fmt.Println("Counter:", counter)
	fmt.Println("Execution time: ", end.Sub(start))
}

func processLines(lines <-chan string, processedLines chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	reGame := regexp.MustCompile(`Game (\d+):`)

	for line := range lines {
		fmt.Println("line", line)
		splitSubGames := strings.Split(line, ";")
		match := reGame.FindStringSubmatch(line)
		gameIDStr := match[1]
		gameID, err := strconv.Atoi(gameIDStr)
		if err != nil {
			fmt.Println("Error converting game ID to integer")
		}
		fmt.Println(gameID)

		re := regexp.MustCompile(`(\d+)\s(\w+)`)

		colors := make(map[string]int)
		for _, subGame := range splitSubGames {
			matches := re.FindAllStringSubmatch(subGame, -1)

			for _, m := range matches {
				number, _ := strconv.Atoi(m[1])
				color := strings.ToLower(m[2])
				if colors[color] < number {
					colors[color] = number
				}
			}

			// PART 1
			//if (colors["red"] <= maxRed || colors["red"] == 0) && (colors["green"] <= maxGreen || colors["green"] == 0) && (colors["blue"] <= maxBlue || colors["blue"] == 0) {
			//	isLimitReached = false
			//} else {
			//	isLimitReached = true
			//	fmt.Println(colors, isLimitReached)
			//	break
			//}
		}

		// PART 2
		multiplier := 1
		for _, number := range colors {
			multiplier *= number
		}

		processedLines <- multiplier

		// PART 1
		//if !isLimitReached {
		//	processedLines <- gameID
		//}
	}
}
