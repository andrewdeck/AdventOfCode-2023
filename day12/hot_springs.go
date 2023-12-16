package main

import (
	"fmt"
	"os"
	"strconv"
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
	// PartTwo(input, 1000000)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	lines := strings.Split(input, "\n")

	totalPossibilities := 0

	for _, line := range lines {
		parts := strings.Fields(line)
		springsStr, pattern := parts[0], parts[1]

		possibilities := 0
		springs := []rune(springsStr)
		unknownIndexes := []int{}
		for i := 0; i < len(springs); i++ {
			if springs[i] == '?' {
				unknownIndexes = append(unknownIndexes, i)
			}
		}

		limit := 2
		for i := 1; i < len(unknownIndexes); i++ {
			limit *= 2
		}

		for i := 0; i < limit; i++ {
			bitmap := fmt.Sprintf("%024b", i)
			tempSprings := make([]rune, len(springs))
			copy(tempSprings, springs)
			bits := []rune(bitmap)
			for bitIdx, springIdx := range unknownIndexes {
				bit := bits[len(bits)-1-bitIdx]
				if bit == '1' {
					tempSprings[springIdx] = '#'
				} else {
					tempSprings[springIdx] = '.'
				}
			}
			if string(tempSprings) == "#.#.###" {
				fmt.Println("here")
			}
			lengths := []string{}
			curLength := 0
			for _, char := range tempSprings {
				if char == '#' {
					curLength++
				} else if curLength > 0 {
					lengths = append(lengths, strconv.Itoa(curLength))
					curLength = 0
				}
			}
			if curLength > 0 {
				lengths = append(lengths, strconv.Itoa(curLength))
			}

			if pattern == strings.Join(lengths, ",") {
				possibilities++
			}
		}
		totalPossibilities += possibilities
	}
	fmt.Println(totalPossibilities)
}
