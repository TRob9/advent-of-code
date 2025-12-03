//go:build part1

package main

import (
	"os"
	"strings"
	"testing"
)

var input1 []string
var anchor1 int

func init() {
	data, _ := os.ReadFile("../input/day1.txt")
	input1 = strings.Split(string(data), "\n")
}

func BenchmarkPart1(b *testing.B) {
	var result int
	for b.Loop() {
		result = howManyZeros(convertTurns(input1))
	}
	anchor1 = result
}
