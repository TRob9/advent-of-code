//go:build part1

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/trob9/advent-of-code/pkg/harness"
)

// AUTO-SUBMIT: When set to true, running the relevant make test command will auto submit answer if testcases pass (ie. make test-d1p1)
const autoSubmit = true

var Answer int = 0 // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 4))
	}

	h := harness.New(solve, &Answer, 1, opts...)

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
	cleanedInput := make([][]rune, len(lines))
	for i, line := range lines {
		cleanedInput[i] = []rune(line)
	}

	rows := len(cleanedInput)
	cols := len(cleanedInput[0])

	// Check corners
	if cleanedInput[0][0] == 64 {
		Answer++
	}
	if cleanedInput[0][cols-1] == 64 {
		Answer++
	}
	if cleanedInput[rows-1][0] == 64 {
		Answer++
	}
	if cleanedInput[rows-1][cols-1] == 64 {
		Answer++
	}

	// Inner grid excluding outermost rows and columns
	for i := 1; i < rows-1; i++ { // excluding outermost left and right columns
		for j := 1; j < cols-1; j++ { // excluding outermost top and bottom rows
			count := 0
			if cleanedInput[i][j] == 64 {
				if cleanedInput[i][j-1] == 64 { // left side check
					count++
				}
				if cleanedInput[i][j+1] == 64 { // right side check
					count++
				}
				for e := -1; e <= 1; e++ {
					if cleanedInput[i-1][j+e] == 64 { // upper row check
						count++
					}
					if cleanedInput[i+1][j+e] == 64 { // lower row check
						count++
					}
				}
				if count < 4 {
					Answer++
				}
				count = 0
			}
		}
	}

	// Top row excluding corners
	for j := 1; j < cols-1; j++ {
		if cleanedInput[0][j] == 64 {
			count := 0
			if cleanedInput[0][j-1] == 64 { // left side check
				count++
			}
			if cleanedInput[0][j+1] == 64 { // right side check
				count++
			}
			for e := -1; e <= 1; e++ {
				if cleanedInput[1][j+e] == 64 { // lower row check
					count++
				}
			}
			if count < 4 {
				Answer++
			}
			count = 0
		}
	}

	// Bottom row excluding corners
	for j := 1; j < cols-1; j++ {
		if cleanedInput[rows-1][j] == 64 {
			count := 0
			if cleanedInput[rows-1][j-1] == 64 { // left side check
				count++
			}
			if cleanedInput[rows-1][j+1] == 64 { // right side check
				count++
			}
			for e := -1; e <= 1; e++ {
				if cleanedInput[rows-2][j+e] == 64 { // upper row check
					count++
				}
			}
			if count < 4 {
				Answer++
			}
			count = 0
		}
	}

	// Leftmost column excluding corners
	for i := 1; i < rows-1; i++ {
		if cleanedInput[i][0] == 64 {
			count := 0
			if cleanedInput[i-1][0] == 64 { // above check
				count++
			}
			if cleanedInput[i+1][0] == 64 { // below check
				count++
			}
			for e := -1; e <= 1; e++ {
				if cleanedInput[i+e][1] == 64 { // right side check
					count++
				}
			}
			if count < 4 {
				Answer++
			}
			count = 0
		}
	}

	// Rightmost column excluding corners
	for i := 1; i < rows-1; i++ {
		if cleanedInput[i][cols-1] == 64 {
			count := 0
			if cleanedInput[i-1][cols-1] == 64 { // above check
				count++
			}
			if cleanedInput[i+1][cols-1] == 64 { // below check
				count++
			}
			for e := -1; e <= 1; e++ {
				if cleanedInput[i+e][cols-2] == 64 { // left side check
					count++
				}
			}
			if count < 4 {
				Answer++
			}
			count = 0
		}
	}
}
