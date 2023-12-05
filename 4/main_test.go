package main

import (
	"testing"
)

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func Test_Main_Day_4(t *testing.T) {
	t.Log("Testing AoC 2023 - Day 4")

	file := []byte(`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`)

	winningCards := part1(file)
	t.Log(winningCards)

	if winningCards != 13 {
		t.Fatalf("%s\t It should return 13 as the winning cards, but returned %d", Failed, winningCards)
	}
	t.Logf("%s\t It should return 13 as the winning cards, and it did", Success)

	cumulativeWinning := part2(file)
	if cumulativeWinning != 30 {
		t.Fatalf("%s\t It should return 30 as cumlative cards, but returned %d", Failed, cumulativeWinning)
	}
	t.Logf("%s\t It should return 30 as cumlative cards, and it did", Success)
}
