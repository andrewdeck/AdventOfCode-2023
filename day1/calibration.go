package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const test = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func main() {
	startTime := time.Now()
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	partOne(input)

	// partOne(test)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))
	startTime = time.Now()
	// partTwo(test)
	partTwo(input)
	fmt.Printf("Part 2: %v\n", time.Since(startTime))

	// var mem runtime.MemStats
	// runtime.ReadMemStats(&mem)

	// fmt.Printf("Allocated memory: %v bytes\n", mem.Alloc)
	// fmt.Printf("Total memory allocated (since start): %v bytes\n", mem.TotalAlloc)
}

func partOne(input string) {
	sum := 0
	lines := strings.Split(input, "\n")

	regex := regexp.MustCompile("[^0-9]+")

	for _, line := range lines {
		digitsArray := regex.Split(line, -1)
		digits := strings.Join(digitsArray, "")
		number := string(digits[0]) + string(digits[len(digits)-1])
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}

func partTwo(input string) {
	sum := 0
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		digits := ExtractNumbers(line)
		number := string(digits[0]) + string(digits[len(digits)-1])
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}

func ExtractNumbers(str string) string {
	newStr := []byte(str)
	result := []byte{}

	regex := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")

	words := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	location := regex.FindIndex(newStr)
	for location != nil {
		word := string(newStr[location[0]:location[1]])
		number := words[string(word)]
		if number == "" {
			number = word
		}
		result = append(result, []byte(number)...)
		newStr = newStr[location[0]+1:]
		location = regex.FindIndex(newStr)
	}

	return string(result)
}
