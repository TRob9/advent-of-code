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
		opts = append(opts, harness.WithSubmit(2025, 6))
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
	nums1, nums2, nums3, nums4 := []int{}, []int{}, []int{}, []int{}
	operators := []string{}
	for i, line := range lines {
		if i == 0 {
			fields := strings.Fields(line)
			for _, field := range fields {
				num, _ := strconv.Atoi(field)
				nums1 = append(nums1, num)
			}
		}
		if i == 1 {
			fields := strings.Fields(line)
			for _, field := range fields {
				num, _ := strconv.Atoi(field)
				nums2 = append(nums2, num)
			}
		}
		if i == 2 {
			fields := strings.Fields(line)
			for _, field := range fields {
				num, _ := strconv.Atoi(field)
				nums3 = append(nums3, num)
			}
		}
		if i == 3 {
			fields := strings.Fields(line)
			for _, field := range fields {
				num, _ := strconv.Atoi(field)
				nums4 = append(nums4, num)
			}
		}
		if i == 4 {
			fields := strings.Fields(line)
			for _, field := range fields {
				operators = append(operators, field)
			}
		}
	}
	for i, operator := range operators {
		if operator == "*" {
			Answer += nums1[i] * nums2[i] * nums3[i] * nums4[i]
		}
		if operator == "+" {
			Answer += nums1[i] + nums2[i] + nums3[i] + nums4[i]
		}
	}
	fmt.Println(Answer)
}
