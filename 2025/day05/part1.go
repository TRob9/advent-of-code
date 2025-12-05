//go:build part1

package main

import (
	"fmt"
	"github.com/trob9/advent-of-code/pkg/harness"
	"os"
	"strconv"
	"strings"
)

// AUTO-SUBMIT: When set to true, running the relevant make test command will auto submit answer if testcases pass (ie. make test-d1p1)
const autoSubmit = true

var Answer int // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 5))
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
	cleanedInput := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	ranges := strings.Split(strings.TrimSpace(string(cleanedInput[0])), "\n")
	ingredientsRaw := strings.Split(strings.TrimSpace((cleanedInput[1])), "\n")
	ingredients := make([]int, len(ingredientsRaw))
	for i, line := range ingredientsRaw {
		ingredients[i], _ = strconv.Atoi(line)
	}
	rangeMap := make(map[int]int)
	for _, p := range ranges {
		pair := strings.Split(p, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		if rangeMap[start] < end {
			rangeMap[start] = end
		}
	}
	for _, ingredient := range ingredients {
		Count := 0
		for i, value := range rangeMap {
			if ingredient >= i && ingredient <= value {
				Count++
			}
		}
		if Count > 0 {
			Answer++
			Count = 0
		}
	}
}
