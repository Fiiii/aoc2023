package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X, Y int
}

func (p Point) addPoint(d Point) Point {
	return Point{p.X + d.X, p.Y + d.Y}
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting working directory:", err)
		return
	}

	scratcherFolderPath := fmt.Sprintf("%s/%s", dir, "3")
	inputFilePath := scratcherFolderPath + "/" + "input.txt"

	file, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	fmt.Println(part1(file))
	fmt.Println(part2(file))
}

func getSymbols(file []byte) map[Point]string {
	symbols := map[Point]string{}

	for y, s := range strings.Fields(string(file)) {
		for x, r := range s {
			if r != '.' && !unicode.IsDigit(r) {
				symbols[Point{x, y}] = string(r)
			}
		}
	}
	return symbols
}

func getEngineParts(file []byte, symbols map[Point]string) map[Point][]int {
	engineParts := map[Point][]int{}
	re := regexp.MustCompile(`\d+`)
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for y, s := range strings.Fields(string(file)) {
		for _, match := range re.FindAllStringIndex(s, -1) {
			keys := map[Point]bool{}
			for x := match[0]; x < match[1]; x++ {
				for _, d := range directions {
					keys[Point{x, y}.addPoint(d)] = true
				}
			}

			n, _ := strconv.Atoi(s[match[0]:match[1]])
			for p := range keys {
				if _, exists := symbols[p]; exists {
					engineParts[p] = append(engineParts[p], n)
				}
			}
		}
	}
	return engineParts

}
func part1(file []byte) int {
	symbols := getSymbols(file)
	engineParts := getEngineParts(file, symbols)
	partNumbers := 0
	for _, values := range engineParts {
		for _, value := range values {
			partNumbers += value
		}
	}
	return partNumbers
}

func part2(file []byte) int {
	symbols := getSymbols(file)
	engine_parts := getEngineParts(file, symbols)
	gear_ratio := 0
	for index, neighbors := range engine_parts {
		if symbols[index] == "*" && len(neighbors) == 2 {
			gear_ratio += neighbors[0] * neighbors[1]
		}
	}
	return gear_ratio
}
