package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const test = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

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

	// startTime = time.Now()
	// PartTwo(test)
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

type Range struct {
	destination int
	source      int
	offset      int
}

/*
Returns:

	[]seeds
	map of [source]destination
	map of [source][]Range
*/
func ParseInput(input string) ([]int, map[string]string, map[string][]Range) {
	lines := strings.Split(input, "\n")
	seedStrs := strings.Fields(lines[0])[1:]
	seeds := make([]int, len(seedStrs))
	for i, seed := range seedStrs {
		seeds[i] = ParseInt(seed)
	}
	// seed-to-soil map:
	mapPattern := regexp.MustCompile(`(\w+)-to-(\w+)\smap:`)

	var source string
	sourceMap := map[string]string{}
	rangesForSource := map[string][]Range{}
	for x := 2; x < len(lines); x++ {
		line := lines[x]
		matches := mapPattern.FindStringSubmatch(line)
		if matches != nil {
			source = matches[1]
			sourceMap[source] = matches[2]
		} else if line != "" {
			fields := strings.Fields(line)
			rangesForSource[source] = append(rangesForSource[source], Range{
				destination: ParseInt(fields[0]),
				source:      ParseInt(fields[1]),
				offset:      ParseInt(fields[2]),
			})
		}
	}

	return seeds, sourceMap, rangesForSource
}

func ParseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func PartOne(input string) {
	seeds, sourceMap, rangesForSource := ParseInput(input)

	lowestLoc := math.MaxInt

	for _, seed := range seeds {
		source := "seed"
		sourceId := seed
		for source != "location" {
			destination := sourceMap[source]
			ranges := rangesForSource[source]
			for _, r := range ranges {
				if sourceId >= r.source && sourceId < r.source+r.offset {
					sourceId = r.destination + (sourceId - r.source)
					break
				}
			}
			source = destination
		}
		if sourceId < lowestLoc {
			lowestLoc = sourceId
		}
	}

	fmt.Println(lowestLoc)
}

func PartTwo(input string) {

}
