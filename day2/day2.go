package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func run(state []int) bool {
	for i := 0; i < len(state)-3; i += 4 {
		switch state[i] {
		case 1:
			//addition
			if state[i+3] >= len(state) || state[i+1] >= len(state) || state[i+2] >= len(state) {
				return false
			}
			state[state[i+3]] = state[state[i+1]] + state[state[i+2]]
		case 2:
			//mult
			if state[i+3] >= len(state) || state[i+1] >= len(state) || state[i+2] >= len(state) {
				return false
			}
			state[state[i+3]] = state[state[i+1]] * state[state[i+2]]
		default:
			break
		}
	}
	return true
}

func reset(state []int, strinput []string) {
	for i := 0; i < len(strinput); i++ {
		state[i], _ = strconv.Atoi(strinput[i])
	}
}

func main() {
	lineScanner := bufio.NewScanner(os.Stdin)
	lineScanner.Scan()
	strinput := strings.Split(lineScanner.Text(), ",")
	state := make([]int, len(strinput))

	for noun := 0; state[0] != 19690720 && noun < 10000; noun++ {
		for verb := 0; state[0] != 19690720 && verb < 10000; verb++ {
			reset(state, strinput)
			state[1] = noun
			state[2] = verb
			run(state)
		}
	}

	for i := 0; i < len(state); i++ {
		strinput[i] = strconv.Itoa(state[i])
	}

	println(strings.Join(strinput, ","))
}
