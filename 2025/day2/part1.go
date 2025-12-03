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
		opts = append(opts, harness.WithSubmit(2025, 2))
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
	Answer = 0
	convertProductCodes(strings.TrimSpace(string(input)))
}

func convertProductCodes(input string) {
	pairings := strings.Split(input, ",")
	for _, p := range pairings {
		pair := strings.Split(p, "-")
		start, _ := strconv.Atoi(pair[0])
		end, _ := strconv.Atoi(pair[1])
		for i := start; i <= end; i++ {
			compare(i)
		}
	}
}

func compare(productCode int) {
	halfLength := len(strconv.Itoa(productCode)) / 2
	modifier := 1
	for i := 0; i < halfLength; i++ {
		modifier *= 10
	}
	lastHalf := productCode % modifier
	if strconv.Itoa(productCode) == strconv.Itoa(lastHalf)+strconv.Itoa(lastHalf) {
		Answer += productCode
	}
}
