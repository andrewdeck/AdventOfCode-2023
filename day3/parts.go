package main

import (
	"adventofcode"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

const test = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

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

	startTime = time.Now()
	// PartTwo(test)
	PartTwo(input)
	fmt.Printf("Part 2: %v\n", time.Since(startTime))
	adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	data := make([][]rune, height)
	for i := range data {
		data[i] = make([]rune, width)
	}

	for y, line := range lines {
		runes := []rune(line)
		copy(data[y], runes)
	}

	// What is the sum of all of the part numbers in the engine schematic?
	sum := 0

	for y := range data {
		for x := 0; x < width; x++ {
			if unicode.IsDigit(data[y][x]) {
				end := x
				for xi := x + 1; xi < width && unicode.IsDigit(data[y][xi]); xi++ {
					end = xi
				}
				number, err := strconv.Atoi(string(data[y][x : end+1]))
				if err != nil {
					panic(err)
				}
				adjacentSymbol := false
				for _, dy := range []int{-1, 0, 1} {
					yi := y + dy
					for xi := x - 1; xi < end+2; xi++ {
						if yi >= 0 && yi < height && xi >= 0 && xi < width {
							if IsSymbol(data[yi][xi]) {
								adjacentSymbol = true
							}
						}
					}
				}

				if adjacentSymbol {
					sum += number
				}
				x = end + 1 // prevent double counting
			}
		}
	}

	fmt.Println(sum)
}

func IsSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func PartTwo(input string) {
	lines := strings.Split(input, "\n")
	height := len(lines)
	width := len(lines[0])
	data := make([][]rune, height)
	for i := range data {
		data[i] = make([]rune, width)
	}

	for y, line := range lines {
		runes := []rune(line)
		copy(data[y], runes)
	}

	gearMap := map[string][]int{}

	for y := range data {
		for x := 0; x < width; x++ {
			if unicode.IsDigit(data[y][x]) {
				end := x
				for xi := x + 1; xi < width && unicode.IsDigit(data[y][xi]); xi++ {
					end = xi
				}
				number, err := strconv.Atoi(string(data[y][x : end+1]))
				if err != nil {
					panic(err)
				}

				for _, dy := range []int{-1, 0, 1} {
					yi := y + dy
					for xi := x - 1; xi < end+2; xi++ {
						if yi >= 0 && yi < height && xi >= 0 && xi < width {
							if data[yi][xi] == '*' {
								key := fmt.Sprintf("%v:%v", xi, yi)
								gearMap[key] = append(gearMap[key], number)
							}
						}
					}
				}

				x = end + 1 // prevent double counting
			}
		}
	}

	sum := 0

	for _, gears := range gearMap {
		if len(gears) == 2 {
			sum += gears[0] * gears[1]
		}
	}

	fmt.Println(sum)
}
