//go:build part1

package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/trob9/advent-of-code/pkg/harness"
)

// AUTO-SUBMIT: When set to true, running the relevant make test command will auto submit answer if testcases pass (ie. make test-d1p1)
const autoSubmit = true

var Answer int //  Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 1))
	}

	h := harness.New(solve, &Answer, 1, opts...)

	// Run tests first
	if passed, err := h.RunTests(); err != nil {
		fmt.Printf("Test error: %v\n", err)
		return
	} else if !passed {
		return
	}

	// Run actual solution
	h.Run()
}

func solve(input []byte) {
	turns := strings.Split(strings.TrimSpace(string(input)), "\n")
	Answer = howManyZeros(convertTurns(turns))
}

func howManyZeros(turns []int) int {
	position := 50
	zeros := 0

	for _, turn := range turns {
		position = position + turn
		position = (position%100 + 100) % 100

		if position == 0 {
			zeros++
		}
	}

	return zeros
}

func convertTurns(turns []string) []int {
	convertedTurns := make([]int, len(turns))

	for i, turn := range turns {
		r := []rune(turn)

		if r[0] == 'L' {
			r[0] = '-'
		} else if r[0] == 'R' {
			r[0] = '+'
		}

		turn = string(r)
		converted, _ := strconv.Atoi(turn)
		convertedTurns[i] = converted
	}

	return convertedTurns
}
