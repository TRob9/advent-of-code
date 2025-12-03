//go:build part2

package main

import (
	"fmt"
	"os"
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
	limit := len(strconv.Itoa(productCode)) / 2
	runes := []rune(strconv.Itoa(productCode))
	for i := limit; i >= 1; i-- {
		if len(strconv.Itoa(productCode))%i == 0 {
			newRunes := []rune{}
			for j := 0; j < i; j++ {
				newRunes = append(newRunes, runes[j])
			}
			s := string(newRunes)
			if concatenate(s, len(strconv.Itoa(productCode))/i) == strconv.Itoa(productCode) {
				Answer += productCode
				break
			}
		}
	}
}

func concatenate(partial string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += partial
	}
	return result
}
