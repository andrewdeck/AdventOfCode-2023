package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

const test = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func main() {
	startTime := time.Now()
	// PartOne(test)
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	PartOne(input)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))

	// startTime = time.Now()
	// PartTwo(test)
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	cards := strings.Split(input, "\n")
	cardPattern := regexp.MustCompile("Card [0-9]+:")

	sum := 0.0

	for _, card := range cards {
		parts := strings.Split(cardPattern.ReplaceAllString(card, ""), "|")
		winning := strings.Fields(parts[0])
		mine := strings.Fields(parts[1])

		count := 0
		for _, num := range mine {
			for _, w := range winning {
				if num == w {
					count++
				}
			}
		}
		if count > 0 {
			sum += math.Pow(2, float64(count-1))
		}
	}

	fmt.Println(sum)
}
