package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

const test = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

func main() {
	// PartOne(test)
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)

	startTime := time.Now()
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
	network := map[string][]string{}

	pattern := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	directions := []rune(lines[0])
	for _, line := range lines[2:] {
		pieces := pattern.FindStringSubmatch(line)
		network[pieces[1]] = []string{pieces[2], pieces[3]}
	}
	location := "AAA"
	for i := 0; i < math.MaxInt; i++ {
		dir := directions[i%len(directions)]
		if dir == 'L' {
			location = network[location][0]
		} else {
			location = network[location][1]
		}
		if location == "ZZZ" {
			fmt.Println(i + 1)
			break
		}
	}
}
