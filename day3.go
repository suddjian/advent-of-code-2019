package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

func parseCode(code string) (bool, int, int) {
	distance, err := strconv.Atoi(code[1:])
	if err != nil {
		log.Fatalf("invalid code: %v", code)
	}
	sign := 1
	vertical := true
	switch code[0] {
	case 'U':
	case 'D':
		sign = -1
	case 'L':
		vertical = false
		sign = -1
	case 'R':
		vertical = false
	default:
		log.Fatalf("invalid code: %v", code)
	}
	return vertical, distance, sign
}

func walkPath(path []string, f func(p pos, s int)) {
	x, y, steps := 0, 0, 0
	for i := range path {
		vertical, distance, sign := parseCode(path[i])
		for i := 0; i < distance; i++ {
			if vertical {
				y += sign
			} else {
				x += sign
			}
			steps++
			f(pos{x, y}, steps)
		}
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func manhattan(x, y int) int {
	return abs(x) + abs(y)
}

func solve(a, b []string) int {
	board := make(map[pos]int)
	walkPath(a, func(p pos, steps int) {
		if board[p] == 0 {
			board[p] = steps
		}
	})

	minDistance := 0
	walkPath(b, func(p pos, steps int) {
		v := board[p]
		if v != 0 {
			distance := v + steps
			if minDistance == 0 || distance < minDistance {
				minDistance = distance
			}
		}
	})

	return minDistance
}

func main() {
	lineScanner := bufio.NewScanner(os.Stdin)
	solution := solve(scanLine(lineScanner), scanLine(lineScanner))
	fmt.Println(solution)
}

func scanLine(scanner *bufio.Scanner) []string {
	scanner.Scan()
	return strings.Split(scanner.Text(), ",")
}
