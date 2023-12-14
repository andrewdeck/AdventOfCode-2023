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

func PartOne(input string) {
	images := ParseInput(input)
	sum := 0
	/* To summarize your pattern notes, add up the number of columns to the left of
	each vertical line of reflection; to that, also add 100 multiplied by the
	number of rows above each horizontal line of reflection. */
	for _, image := range images {
		if ok, index := HorizontalSymmetry(image, 0); ok {
			sum += (100 * (index + 1))
		} else if ok, index := VerticalSymmetry(image, 0); ok {
			sum += (index + 1)
		}
	}
	fmt.Println(sum)
}

func PartTwo(input string) {
	images := ParseInput(input)
	sum := 0
	/* To summarize your pattern notes, add up the number of columns to the left of
	each vertical line of reflection; to that, also add 100 multiplied by the
	number of rows above each horizontal line of reflection. */
	for _, image := range images {
		if ok, index := HorizontalSymmetry(image, 1); ok {
			sum += (100 * (index + 1))
		} else if ok, index := VerticalSymmetry(image, 1); ok {
			sum += (index + 1)
		}
	}
	fmt.Println(sum)
}

func ParseInput(input string) [][][]bool {
	chunks := strings.Split(input, "\n\n")

	images := make([][][]bool, len(chunks))

	for i, chunk := range chunks {
		lines := strings.Split(chunk, "\n")
		image := make([][]bool, len(lines))
		for y, line := range lines {
			row := make([]bool, len([]rune(line)))
			for x, char := range []rune(line) {
				row[x] = char == '#'
			}
			image[y] = row
		}
		images[i] = image
	}
	return images
}

func VerticalSymmetry(image [][]bool, errors int) (bool, int) {
	height := len(image)
	width := len(image[0])
	hasSymmetry := false
	index := 0
	for x := 0; x < width-1; x++ {
		diffCount := 0

		for dx := 0; x-dx >= 0 && x+dx+1 < width; dx++ {
			for y := 0; y < height; y++ {
				if image[y][x-dx] != image[y][x+dx+1] {
					diffCount++
				}
			}
		}

		if diffCount == errors {
			hasSymmetry = true
			index = x
			break
		}
	}

	return hasSymmetry, index
}

func HorizontalSymmetry(image [][]bool, errors int) (bool, int) {
	height := len(image)
	width := len(image[0])
	hasSymmetry := false
	index := 0

	for y := 0; y < height-1; y++ {
		diffCount := 0
		for dy := 0; y-dy >= 0 && y+dy+1 < height; dy++ {
			for x := 0; x < width; x++ {
				if image[y-dy][x] != image[y+dy+1][x] {
					diffCount++
				}
			}
		}
		if diffCount == errors {
			hasSymmetry = true
			index = y
			break
		}
	}
	return hasSymmetry, index
}
