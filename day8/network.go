package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
)

const test = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

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

	startTime = time.Now()
	PartTwo(test)
	PartTwo(input)
	fmt.Printf("Part 2: %v\n", time.Since(startTime))
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

func PartTwo(input string) {
	lines := strings.Split(input, "\n")
	network := map[string][]string{}

	pattern := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	directions := []rune(lines[0])
	for _, line := range lines[2:] {
		pieces := pattern.FindStringSubmatch(line)
		network[pieces[1]] = []string{pieces[2], pieces[3]}
	}
	locations := []string{}
	for loc := range network {
		if strings.HasSuffix(loc, "A") {
			locations = append(locations, loc)
		}
	}

	locationCycle := make([]int, len(locations))
	for i := 0; i < math.MaxInt; i++ {
		dir := directions[i%len(directions)]
		for j := range locations {
			if dir == 'L' {
				locations[j] = network[locations[j]][0]
			} else {
				locations[j] = network[locations[j]][1]
			}
		}
		for j, loc := range locations {
			if locationCycle[j] != 0 {
				continue
			}
			if strings.HasSuffix(loc, "Z") {
				locationCycle[j] = i + 1
			}
		}
		missingCyle := false
		for _, cycle := range locationCycle {
			if cycle == 0 {
				missingCyle = true
				break
			}
		}
		if !missingCyle {
			break
		}
	}
	fmt.Println(LCM(locationCycle[0], locationCycle[1], locationCycle[2:]...))
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
