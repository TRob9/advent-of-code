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

var Answer int // Change to string if needed

func main() {
	var opts []harness.Option
	if autoSubmit {
		opts = append(opts, harness.WithSubmit(2025, 3))
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
	Answer = 0
	banks := strings.Split(strings.TrimSpace(string(input)), "\n")

	for _, battery := range banks {
		runes := []rune(battery)

		var A int
		var AIndex int
		var B int

		for i, r := range runes {
			converted := int(r - '0')

			if converted > A && i != len(runes)-1 {
				A = converted
				AIndex = i
			}
		}
		for i, r := range runes {
			converted := int(r - '0')

			if converted > B && i > AIndex {
				B = converted
			}
		}

		Answer += A*10 + B
		A = 0
		B = 0
		AIndex = 0
	}
}
