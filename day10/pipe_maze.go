package main

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

const test = `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`

func main() {
	// file, err := os.ReadFile("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// input := string(file)
	startTime := time.Now()
	PartOne(test)
	// PartOne(input)
	fmt.Printf("Part 1: %v\n", time.Since(startTime))

	// startTime = time.Now()
	// PartTwo(test)
	// PartTwo(input)
	// fmt.Printf("Part 2: %v\n", time.Since(startTime))
	// adventofcode.PrintMemoryUsage()
}

func PartOne(input string) {
	pipes, start := ParsePipes(input)
	paths := []Pipe{}

	if start.pos.x < len(pipes[0]) {
		eastPipe := pipes[start.pos.y][start.pos.x+1]
		if slices.Contains(eastPipe.outputs, 'W') {
			paths = append(paths, eastPipe)
		}
	}
	if start.pos.x > 0 {
		westPipe := pipes[start.pos.y][start.pos.x-1]
		if slices.Contains(westPipe.outputs, 'E') {
			paths = append(paths, westPipe)
		}
	}
	if start.pos.y > 0 {
		northPipe := pipes[start.pos.y-1][start.pos.x]
		if slices.Contains(northPipe.outputs, 'S') {
			paths = append(paths, northPipe)
		}
	}
	if start.pos.y < len(pipes) {
		southPipe := pipes[start.pos.y+1][start.pos.x]
		if slices.Contains(southPipe.outputs, 'N') {
			paths = append(paths, southPipe)
		}
	}

	pipeA := paths[0]
	pipeB := paths[1]
	steps := 1
	for pipeA.pos.x != pipeB.pos.x && pipeA.pos.y != pipeB.pos.y {

	}
}

type Position struct {
	x, y int
}

type Pipe struct {
	outputs []rune
	pos     Position
}

func isOppositeDirection(a, b rune) bool {
	m := map[rune]int{
		'N': 1,
		'S': -1,
		'W': 2,
		'E': -2,
	}
	return m[a]+m[b] == 0
}

func (p *Pipe) nextPosition(input rune) Position {
	var direction rune
	for _, output := range p.outputs {
		if !isOppositeDirection(input, output) {
			direction = output
		}
	}

	newPos := Position{
		x: p.pos.x,
		y: p.pos.y,
	}
	switch direction {
	case 'N':
		newPos.y--
	case 'S':
		newPos.y++
	case 'E':
		newPos.x++
	case 'W':
		newPos.x--
	}
	return newPos
}

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
*/
func ParsePipes(input string) ([][]Pipe, Pipe) {
	lines := strings.Split(input, "\n")

	pipes := make([][]Pipe, len(lines))
	var start Pipe
	for x, line := range lines {
		pipes[x] = make([]Pipe, len(line))
	}

	for y, line := range lines {
		for x, pipe := range []rune(line) {
			if pipe == '.' {
				continue
			}
			p := Pipe{
				pos: Position{
					x: x,
					y: y,
				},
			}
			switch pipe {
			case '|':
				p.outputs = []rune{'N', 'S'}
			case '-':
				p.outputs = []rune{'E', 'W'}
			case 'L':
				p.outputs = []rune{'E', 'N'}
			case 'J':
				p.outputs = []rune{'N', 'W'}
			case '7':
				p.outputs = []rune{'S', 'W'}
			case 'F':
				p.outputs = []rune{'E', 'S'}
			}
			pipes[y][x] = p
			if pipe == 'S' {
				start = p
			}
		}
	}

	return pipes, start
}
