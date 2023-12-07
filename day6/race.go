package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const test = `Time:      71530
Distance:  940200`

const input = `Time:        53837288
Distance:   333163512891532`

func main() {
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
	races := make([][]int, 4)
	lines := strings.Split(input, "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])

	for i, time := range times {
		if i > 0 {
			t, err := strconv.Atoi(time)
			if err != nil {
				panic(err)
			}
			races[i-1] = []int{t}
		}
	}
	for i, distance := range distances {
		if i > 0 {
			d, err := strconv.Atoi(distance)
			if err != nil {
				panic(err)
			}
			races[i-1] = append(races[i-1], d)
		}
	}

	result := 1
	for _, race := range races {
		if len(race) == 0 {
			continue
		}
		time := race[0]
		distance := race[1]

		count := 0
		for x := 1; x < time; x++ {
			if (time-x)*x > distance {
				count++
			}
		}
		result *= count
	}
	fmt.Println(result)
}
