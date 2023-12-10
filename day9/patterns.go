package main

import (
	"adventofcode"
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

const test = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	startTime := time.Now()
	// PartOne(test)
	PartOne(input)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))

	// startTime = time.Now()
	// PartTwo(test)
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	lines := strings.Split(input, "\n")
	history := make([][]int, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		history[i] = make([]int, len(fields))
		for j, field := range fields {
			history[i][j] = adventofcode.ParseInt(field)
		}
	}
	sum := 0
	for _, sequence := range history {
		sum += NextValueInSequence(sequence)
	}
	fmt.Println(sum)
}

func NextValueInSequence(sequence []int) int {
	if slices.Max(sequence) == 0 && slices.Min(sequence) == 0 {
		return 0
	}
	diffs := make([]int, len(sequence)-1)
	for x := 1; x < len(sequence); x++ {
		diffs[x-1] = sequence[x] - sequence[x-1]
	}
	return sequence[len(sequence)-1] + NextValueInSequence(diffs)
}
