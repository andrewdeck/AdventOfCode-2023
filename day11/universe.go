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

	// startTime = time.Now()
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

type Galaxy struct {
	x, y int
}

func PartOne(input string) {
	lines := strings.Split(input, "\n")
	// 1. parse universe
	universe := make([][]bool, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		row := make([]bool, len(runes))
		for x, r := range runes {
			row[x] = r == '#'
		}
		universe[i] = row
	}
	// 2. expand universe
	columns := []int{}
	rows := []int{}
	for y, row := range universe {
		noGalaxies := true
		for _, val := range row {
			if val {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			rows = append(rows, y)
		}
	}
	for x := range universe[0] {
		noGalaxies := true
		for _, row := range universe {
			if row[x] {
				noGalaxies = false
				break
			}
		}
		if noGalaxies {
			columns = append(columns, x)
		}
	}
	for i, row := range rows {
		rowToInsert := row + i
		universe = append(universe[:rowToInsert+1], universe[rowToInsert:]...)
		universe[rowToInsert] = make([]bool, len(universe[0]))
	}
	for i, col := range columns {
		colToInsert := col + i
		for y := range universe {
			universe[y] = append(universe[y][:colToInsert+1], universe[y][colToInsert:]...)
			universe[y][colToInsert] = false
		}
	}
	// 3. extract galaxy locations
	galaxies := []Galaxy{}
	for y, row := range universe {
		for x, val := range row {
			if val {
				galaxies = append(galaxies, Galaxy{x: x, y: y})
			}
		}
	}
	// 4. calculate shortest path between every pair, aka manhattan distance
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			a, b := galaxies[i], galaxies[j]
			dx := a.x - b.x
			dy := a.y - b.y
			if dx < 0 {
				dx *= -1
			}
			if dy < 0 {
				dy *= -1
			}
			sum += dx + dy
		}
	}
	fmt.Println(sum)
}
