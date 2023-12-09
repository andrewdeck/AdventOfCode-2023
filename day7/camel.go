package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

const test = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

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

type CamelHand struct {
	cards      []rune
	wager      int
	ordinality int
}

type ByRank []CamelHand

func (a ByRank) Len() int      { return len(a) }
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool {
	if a[i].ordinality != a[j].ordinality {
		return a[i].ordinality < a[j].ordinality
	} else {
		for x := 0; x < 5; x++ {
			if a[i].cards[x] == a[j].cards[x] {
				continue
			}
			if CardIsLess(a[i].cards[x], a[j].cards[x]) {
				return true
			} else {
				return false
			}
		}
		// this should never be reached
		panic("uh oh")
	}
}

var cardRank = []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

func HandOrdinality(hand []rune) int {
	cardCount := map[rune]int{}
	for _, card := range hand {
		cardCount[card]++
	}
	counts := []int{}
	for _, count := range cardCount {
		counts = append(counts, count)
	}
	countMap := map[int]int{}
	for _, count := range counts {
		countMap[count]++
	}
	// 13 unique cards + 6 hands
	// higher ordinality is a winning hand
	if countMap[5] == 1 {
		return 6
	} else if countMap[4] == 1 {
		return 5
	} else if countMap[3] == 1 && countMap[2] == 1 {
		return 4
	} else if countMap[3] == 1 {
		return 3
	} else if countMap[2] == 2 {
		return 2
	} else if countMap[2] == 1 {
		return 1
	} else {
		return 0
	}
}

func CardIsLess(a, b rune) bool {
	aIndex := slices.Index(cardRank, a)
	bIndex := slices.Index(cardRank, b)
	return aIndex > bIndex
}

func PartOne(input string) {
	lines := strings.Split(input, "\n")
	hands := make([]CamelHand, len(lines))
	for i, line := range lines {
		fields := strings.Fields(line)
		wager, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		cards := []rune(fields[0])
		hands[i] = CamelHand{
			cards:      cards,
			wager:      wager,
			ordinality: HandOrdinality(cards),
		}
	}

	sort.Sort(ByRank(hands))

	sum := 0
	for i, hand := range hands {
		fmt.Println(string(hand.cards))
		sum += (i + 1) * hand.wager
	}
	fmt.Println(sum)
}

// func PartTwo(input string) {

// }
