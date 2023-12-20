package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// file, err := os.ReadFile("test.txt")
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	startTime := time.Now()
	PartOne(input)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))

	// startTime = time.Now()
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	dish := ParseInput(input)

	for y := 1; y < len(dish); y++ {
		for x := 0; x < len(dish[y]); x++ {
			// roll rock north
			if dish[y][x] == 'O' {
				ry := y - 1
				for ry >= 0 && dish[ry][x] == '.' {
					dish[ry][x], dish[ry+1][x] = dish[ry+1][x], dish[ry][x]
					ry--
				}
			}
		}
	}
	load := 0
	for y := range dish {
		for _, char := range dish[y] {
			if char == 'O' {
				load += len(dish) - y
			}
		}
	}
	fmt.Println(load)
}

func PrintState(state [][]rune) {
	for _, row := range state {
		fmt.Println(string(row))
	}
}

func ParseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")

	dish := make([][]rune, len(lines))

	for i, line := range lines {
		dish[i] = []rune(line)
	}
	return dish
}
