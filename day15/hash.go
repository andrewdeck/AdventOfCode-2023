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
	steps := strings.Split(input, ",")

	sum := 0
	for _, step := range steps {
		value := 0
		for _, char := range step {
			value += int(char)
			value *= 17
			value = value % 256
		}
		// fmt.Printf("%v %v\n", step, value)
		sum += value
	}

	fmt.Println(sum)
}
