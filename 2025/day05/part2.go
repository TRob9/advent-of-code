//go:build part2

package main

import (
	"fmt"
	"os"
	"sort"
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
		opts = append(opts, harness.WithSubmit(2025, 5))
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
	cleanedInput := strings.Split(strings.TrimSpace(string(input)), "\n\n")
	ranges := strings.Split(strings.TrimSpace(string(cleanedInput[0])), "\n")
	rangeMap := make(map[int]int)
	for _, p := range ranges {
		pair := strings.Split(p, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		if rangeMap[start] < end {
			rangeMap[start] = end
		}
	}

	ordered := make([]int, 0, len(rangeMap))
	for k := range rangeMap {
		ordered = append(ordered, k)
	}

	sort.Ints(ordered)
	lastA := 0
	lastB := 0
	for _, val := range ordered {
		A := val
		B := rangeMap[val]
		if lastA == 0 && lastB == 0 {
			lastA = A
			lastB = B
			Answer += (B - A) + 1
			continue
		}
		if A == lastB {
			Answer += B - A
			lastA = A
			lastB = B
			continue
		}
		if A < lastB && B <= lastB {
			lastA = A
			lastB = B
			continue
		}
		if A < lastB && B > lastB {
			Answer += B - lastB
			lastA = A
			lastB = B
		} else {
			Answer += (B - A) + 1
			lastA = A
			lastB = B
		}
	}
}
