package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const test = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

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
	// startTime = time.Now()
	// partTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))

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
		fmt.Println(number)
		num, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println(sum)
}
