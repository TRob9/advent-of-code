//go:build part2

package main

import (
	"fmt"
	"os"

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

	Max := 0
	for i, line := range lines {
		parts := strings.Split(line, ",")
		xi, _ := strconv.Atoi(parts[0])
		yi, _ := strconv.Atoi(parts[1])
		for x, line2 := range lines {
	Answer = 0 // Placeholder - replace with your solution logic
}
