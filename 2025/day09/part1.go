//go:build part1

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	// "github.com/trob9/advent-of-code/internal/grid"
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
			if x == i {
				continue
			}
			parts2 := strings.Split(line2, ",")
			xi2, _ := strconv.Atoi(parts2[0])
			yi2, _ := strconv.Atoi(parts2[1])
			maxX := max(xi2, xi)
			maxY := max(yi2, yi)
			x3 := min(xi2, xi)
			y3 := min(yi2, yi)
			side1 := maxX - x3 + 1
			side2 := maxY - y3 + 1
			Area := side1 * side2
			if Area > Max {
				Max = Area
			}
		}
	}
	fmt.Println(Max)
	Answer = 0 // Placeholder - replace with your solution logic
}
