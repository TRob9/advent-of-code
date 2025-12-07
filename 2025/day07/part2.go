//go:build part2

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/trob9/advent-of-code/pkg/harness"
)

// AUTO-SUBMIT: When set to true, running the relevant make test command will auto submit answer if testcases pass (ie. make test-d1p1)
const autoSubmit = true

var Answer int // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 7))
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

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	lines[1] = lines[0]

	cleanedInput := make([][]rune, len(lines))
	for i, line := range lines {
		cleanedInput[i] = []rune(line)
	}
	// create a matching intGrid because converting runes to ints and back is a bad time
	intGrid := make([][]int, len(cleanedInput))
	for i := range intGrid {
		intGrid[i] = make([]int, len(lines))
	}

	for r, row := range cleanedInput {
		for c, char := range row {
			if char == 'S' {
				intGrid[r][c] = 1 // add the position of S to the intGrid as 1
			}
		}
	}

	for r := range cleanedInput {
		if r == 0 { /// skip that first row type shit
			continue
		}
		for c, char := range cleanedInput[r] {
			if char == '^' {
				if intGrid[r-1][c] >= 1 { // if the number above the split icon is more than zero
					intGrid[r][c-1]++ // split to the side, adding one to each
					intGrid[r][c+1]++
				}
			}
			if intGrid[r][c] == 0 && intGrid[r-1][c] >= 1 { // if above number is more than zero
				intGrid[r][c] = intGrid[r-1][c] // the below number becomes the above number, propagating down
			}
		}
	}

	// Sum da last row
	for _, num := range intGrid[len(intGrid)-1] { // length of the array minus one cuz len does not start at 0 bruh
		Answer += num
	}
	fmt.Println(Answer)
	fmt.Println(intGrid)
}
