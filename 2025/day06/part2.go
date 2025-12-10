//go:build part2

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
		opts = append(opts, harness.WithSubmit(2025, 6))
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

	lines := strings.Split(string(input), "\n")

	var runes1, runes2, runes3, runes4, runes5 []rune
	for i, line := range lines { // could have been a switch case thing okay bro I get it
		if i == 0 {
			runes1 = []rune(line)
		}
		if i == 1 {
			runes2 = []rune(line)
		}
		if i == 2 {
			runes3 = []rune(line)
		}
		if i == 3 {
			runes4 = []rune(line)
		}
		if i == 4 {
			runes5 = []rune(line)
		}
	}

	var outerNums [][]int
	var innerNums []int
	var operators []string

	for i := len(runes1) - 1; i >= 0; i-- {
		r1 := runes1[i]
		r2 := runes2[i]
		r3 := runes3[i]
		r4 := runes4[i]
		r5 := runes5[i]

		numString := strings.TrimSpace(string(r1) + string(r2) + string(r3) + string(r4))
		if numString == "" {
			continue
		}

		number, _ := strconv.Atoi(numString)
		innerNums = append(innerNums, number)

		if r5 == '*' {
			operators = append(operators, "*")
			outerNums = append(outerNums, innerNums)
			innerNums = nil
		}
		if r5 == '+' {
			operators = append(operators, "+")
			outerNums = append(outerNums, innerNums)
			innerNums = nil
		}
	}
	for i, operator := range operators {
		if operator == "*" {
			product := 1
			for _, num := range outerNums[i] {
				product *= num
			}
			Answer += product
		}
		if operator == "+" {
			for _, num := range outerNums[i] {
				Answer += num
			}
		}
	}
}
