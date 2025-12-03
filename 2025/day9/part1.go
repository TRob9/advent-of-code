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

var Answer int // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 9))
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
	// Parse input - uncomment the pattern you need, delete the rest

	// a) Array of strings (lines)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// b) Array of strings (comma-separated)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), ",")

	// c) Array of integers (lines)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([]int, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i], _ = strconv.Atoi(line)
	// }

	// d) Array of integers (comma-separated)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make([]int, len(parts))
	// for i, part := range parts {
	// 	cleanedInput[i], _ = strconv.Atoi(strings.TrimSpace(part))
	// }

	// e) Map of string to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[string]int)
	// for _, part := range parts {
	// 	cleanedInput[strings.TrimSpace(part)] = 0
	// }

	// f) Map of int to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[int]int)
	// for _, part := range parts {
	// 	key, _ := strconv.Atoi(strings.TrimSpace(part))
	// 	cleanedInput[key] = 0
	// }

	// g) 2D grid of characters
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]rune, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i] = []rune(line)
	// }

	// h) 2D grid of integers (space-separated)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]int, len(lines))
	// for i, line := range lines {
	// 	parts := strings.Fields(line)
	// 	cleanedInput[i] = make([]int, len(parts))
	// 	for j, part := range parts {
	// 		cleanedInput[i][j], _ = strconv.Atoi(part)
	// 	}
	// }

	// i) Single integer
	// cleanedInput, _ := strconv.Atoi(strings.TrimSpace(string(input)))

	// j) Single string
	// cleanedInput := strings.TrimSpace(string(input))

	Answer = 0 // Placeholder - replace with your solution logic
}
