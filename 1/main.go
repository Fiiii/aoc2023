package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Hello, AoC! - day: 1")

	file, err := os.Open("/Users/arturfigiel/go/aoc2023/1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var counter int
	var strNumber string
	var num int

	for scanner.Scan() {
		strNumber = ""
		num = 0

		line := scanner.Text()
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

		counter += num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	end := time.Now()

	fmt.Println("Counter:", counter)
	fmt.Println("Execution time: ", end.Sub(start))
}
