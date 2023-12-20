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

	startTime = time.Now()
	PartTwo(input)
	fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartTwo(input string) {
	dish := ParseInput(input)

	loadMap := map[int]int{} // load => number of times seen load
	loads := []int{}
	lastCycle := 0
	repeatLength := 0

	for cycle := 0; cycle < 100000; cycle++ {
		RollNorth(dish)
		RollWest(dish)
		RollSouth(dish)
		RollEast(dish)
		load := Load(dish)

		loadMap[load]++
		loads = append(loads, load)
		if loadMap[load] > 8 {
			found, length := CheckForRepetition(loads)
			if found {
				lastCycle = cycle
				repeatLength = length
				break
			}
		}
	}
	if lastCycle == 0 && repeatLength == 0 {
		panic("Did not find repeating pattern")
	}
	offset := (1000000000 - lastCycle) % repeatLength

	fmt.Println(loads[lastCycle-repeatLength+offset-1])
}

func CheckForRepetition(loads []int) (bool, int) {
	found := false
	length := 0
	lastLoad := loads[len(loads)-1]
	for d := 2; d < len(loads)/2; d++ {
		if loads[len(loads)-1-d] == lastLoad {
			// verify all numbers within d repeat
			valid := true
			for i := 1; i < d; i++ {
				if loads[len(loads)-1-i] != loads[len(loads)-1-d-i] {
					valid = false
					break
				}
			}

			if valid {
				found = true
				length = d
				break
			}
		}
	}
	return found, length
}

func RollNorth(dish [][]rune) {
	for y := 1; y < len(dish); y++ {
		for x := 0; x < len(dish[y]); x++ {
			// roll rock north
			if dish[y][x] == 'O' {
				ry := y - 1
				for ry >= 0 && dish[ry][x] == '.' {
					dish[ry][x], dish[ry+1][x] = dish[ry+1][x], dish[ry][x]
					ry--
				}
			}
		}
	}
}

func RollSouth(dish [][]rune) {
	for y := len(dish) - 2; y >= 0; y-- {
		for x := 0; x < len(dish[y]); x++ {
			// roll rock south
			if dish[y][x] == 'O' {
				ry := y + 1
				for ry < len(dish) && dish[ry][x] == '.' {
					dish[ry][x], dish[ry-1][x] = dish[ry-1][x], dish[ry][x]
					ry++
				}
			}
		}
	}
}

func RollWest(dish [][]rune) {
	for x := 1; x < len(dish[0]); x++ {
		for y := 0; y < len(dish); y++ {
			// roll rock west
			if dish[y][x] == 'O' {
				rx := x - 1
				for rx >= 0 && dish[y][rx] == '.' {
					dish[y][rx], dish[y][rx+1] = dish[y][rx+1], dish[y][rx]
					rx--
				}
			}
		}
	}
}
func RollEast(dish [][]rune) {
	for x := len(dish[0]) - 2; x >= 0; x-- {
		for y := 0; y < len(dish); y++ {
			// roll rock west
			if dish[y][x] == 'O' {
				rx := x + 1
				for rx < len(dish[0]) && dish[y][rx] == '.' {
					dish[y][rx], dish[y][rx-1] = dish[y][rx-1], dish[y][rx]
					rx++
				}
			}
		}
	}
}

func PartOne(input string) {
	dish := ParseInput(input)

	RollNorth(dish)

	load := Load(dish)
	fmt.Println(load)
}

func PrintState(state [][]rune) {
	for _, row := range state {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func Load(dish [][]rune) int {
	load := 0
	for y := range dish {
		for _, char := range dish[y] {
			if char == 'O' {
				load += len(dish) - y
			}
		}
	}
	return load
}

func ParseInput(input string) [][]rune {
	lines := strings.Split(input, "\n")

	dish := make([][]rune, len(lines))

	for i, line := range lines {
		dish[i] = []rune(line)
	}
	return dish
}
