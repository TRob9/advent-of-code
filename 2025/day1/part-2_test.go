//go:build part2

package main

import (
	"os"
	"strings"
	"testing"
)

var input2 []string
var anchor2 int

func init() {
	data, _ := os.ReadFile("../input/day1.txt")
	input2 = strings.Split(string(data), "\n")
}

func BenchmarkPart2(b *testing.B) {
	var result int
	for b.Loop() {
		result = howManyZeros(convertTurns(input2))
	}
	anchor2 = result
}
