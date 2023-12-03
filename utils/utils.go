package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string, output chan<- string) {
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
