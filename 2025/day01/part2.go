//go:build part2

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/trob9/advent-of-code/pkg/harness"
)

// AUTO-SUBMIT: When set to true, running the relevant make test command will auto submit answer if testcases pass (ie. make test-d1p1)
const autoSubmit = true

var Answer int // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 1))
	}

	h := harness.New(solve, &Answer, 2, opts...)

	// Run tests first (unless SKIP_TESTS is set)
	if os.Getenv("SKIP_TESTS") == "" {
		if passed, err := h.RunTests(); err != nil {
			fmt.Printf("Test error: %v\n", err)
			return
		} else if !passed {
			return
		}
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
		// normalise the turns to not exceed more than one full rotation, apply zero count for each normalisation needed as will pass 0 each time
		for turn > 100 {
			turn -= 100
			zeros++
		}
		for turn < -100 {
			turn += 100
			zeros++
		}
		// apply the turn to the wheel
		newpos := position + turn
		// account for the position being higher than the upper limit of 99, apply another zero count if not already on or landing on zero
		if newpos > 99 {
			newpos -= 100
			if newpos != 0 && newpos != 100 && position != 0 {
				zeros++
			}
		}
		// account for the position being lower than the lower limit of 0, apply another zero count if not already on or landing on zero
		if newpos < 0 {
			newpos += 100
			if newpos != 0 && newpos != -100 && position != 0 {
				zeros++
			}
		}
		// capture the count of each time the new position is 0
		if newpos == 0 {
			zeros++
		}
		// assign the new position
		position = newpos
	}
	return zeros
}

func convertTurns(turns []string) []int {
	convertedTurns := make([]int, len(turns))
	for i, turn := range turns {
		// convert input string to array of runs
		r := []rune(turn)
		// replace first letter with + or - accordingly so that strconv.Atoi will convert to valid integer
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
