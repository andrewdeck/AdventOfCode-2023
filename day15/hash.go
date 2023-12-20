package main

import (
	"adventofcode"
	"fmt"
	"os"
	"regexp"
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

	startTime = time.Now()
	PartTwo(input)
	fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func Hash(s string) int {
	value := 0
	for _, char := range s {
		value += int(char)
		value *= 17
		value = value % 256
	}
	return value
}

func PartOne(input string) {
	steps := strings.Split(input, ",")

	sum := 0
	for _, step := range steps {
		sum += Hash(step)
	}

	fmt.Println(sum)
}

type Lens struct {
	label       string
	focalLength int
}

func PartTwo(input string) {
	steps := strings.Split(input, ",")

	instructionPattern := regexp.MustCompile(`(\w+)([-=])(\d*)`)

	hashmap := make(map[int][]Lens)

	for _, step := range steps {
		groups := instructionPattern.FindStringSubmatch(step)
		stepLabel, operator, stepFocalLength := groups[1], groups[2], groups[3]
		hash := Hash(stepLabel)
		switch operator {
		case "=":
			existingLens := false
			for i, lens := range hashmap[hash] {
				if lens.label == stepLabel {
					existingLens = true
					hashmap[hash][i].focalLength = adventofcode.ParseInt(stepFocalLength)
					break
				}
			}
			if !existingLens {
				hashmap[hash] = append(hashmap[hash], Lens{
					label:       stepLabel,
					focalLength: adventofcode.ParseInt(stepFocalLength),
				})
			}
		case "-":
			for i, lens := range hashmap[hash] {
				if lens.label == stepLabel {
					hashmap[hash] = append(hashmap[hash][:i], hashmap[hash][i+1:]...)
				}
			}
		}
	}

	focusPower := 0
	for hash, lenses := range hashmap {
		boxPower := 0
		for pos, lens := range lenses {
			boxPower += (hash + 1) * (pos + 1) * lens.focalLength
		}
		focusPower += boxPower
	}
	fmt.Println(focusPower)
}
