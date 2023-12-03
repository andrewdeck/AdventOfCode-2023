package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const test = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func main() {
	startTime := time.Now()
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(file)
	PartOne(input)

	// PartOne(test)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))
	// startTime = time.Now()
	// partTwo(test)
	// partTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))

	// var mem runtime.MemStats
	// runtime.ReadMemStats(&mem)

	// fmt.Printf("Allocated memory: %v bytes\n", mem.Alloc)
	// fmt.Printf("Total memory allocated (since start): %v bytes\n", mem.TotalAlloc)
}

func PartOne(input string) {
	sum := 0
	games := strings.Split(input, "\n")

	for _, game := range games {
		if possible, index := PossibleGame(game); possible {
			sum += index
			fmt.Printf("Valid Game: %v\n", index)
		}
	}
	fmt.Println(sum)
}

var realBag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func PossibleGame(game string) (bool, int) {
	gamePattern := regexp.MustCompile("Game ([0-9]+):")
	gameStrs := gamePattern.FindStringSubmatch(game)
	index, err := strconv.Atoi(gameStrs[1])
	if err != nil {
		panic(err)
	}
	max := make(map[string]int)

	reveals := strings.Split(gamePattern.ReplaceAllString(game, ""), ";")
	for _, reveal := range reveals {
		groups := strings.Split(reveal, ",")
		for _, group := range groups {
			fields := strings.Fields(group)
			color := fields[1]
			count, err := strconv.Atoi(fields[0])
			if err != nil {
				panic(err)
			}
			if count > max[color] {
				max[color] = count
			}
		}
	}

	possible := true
	for color, count := range realBag {
		if max[color] > count {
			possible = false
		}
	}
	return possible, index
}
