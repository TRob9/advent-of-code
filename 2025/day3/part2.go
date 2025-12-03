//go:build part2

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
	banks := strings.Split(strings.TrimSpace(string(input)), "\n")

	for _, battery := range banks {
		runes := []rune(battery)
		totalDigits := 12

		var A, B, C, D, E, F, G, H, I, J, K, L int
		var AIndex, BIndex, CIndex, DIndex, EIndex, FIndex, GIndex, HIndex, IIndex, JIndex, KIndex int

		length := len(runes)

		for i := 0; i <= length-totalDigits; i++ {
			val := int(runes[i] - '0')
			if val > A {
				A = val
				AIndex = i
			}
		}

		for i := AIndex + 1; i <= length-(totalDigits-1); i++ {
			val := int(runes[i] - '0')
			if val > B {
				B = val
				BIndex = i
			}
		}

		for i := BIndex + 1; i <= length-(totalDigits-2); i++ {
			val := int(runes[i] - '0')
			if val > C {
				C = val
				CIndex = i
			}
		}

		for i := CIndex + 1; i <= length-(totalDigits-3); i++ {
			val := int(runes[i] - '0')
			if val > D {
				D = val
				DIndex = i
			}
		}

		for i := DIndex + 1; i <= length-(totalDigits-4); i++ {
			val := int(runes[i] - '0')
			if val > E {
				E = val
				EIndex = i
			}
		}

		for i := EIndex + 1; i <= length-(totalDigits-5); i++ {
			val := int(runes[i] - '0')
			if val > F {
				F = val
				FIndex = i
			}
		}

		for i := FIndex + 1; i <= length-(totalDigits-6); i++ {
			val := int(runes[i] - '0')
			if val > G {
				G = val
				GIndex = i
			}
		}

		for i := GIndex + 1; i <= length-(totalDigits-7); i++ {
			val := int(runes[i] - '0')
			if val > H {
				H = val
				HIndex = i
			}
		}

		for i := HIndex + 1; i <= length-(totalDigits-8); i++ {
			val := int(runes[i] - '0')
			if val > I {
				I = val
				IIndex = i
			}
		}

		for i := IIndex + 1; i <= length-(totalDigits-9); i++ {
			val := int(runes[i] - '0')
			if val > J {
				J = val
				JIndex = i
			}
		}

		for i := JIndex + 1; i <= length-(totalDigits-10); i++ {
			val := int(runes[i] - '0')
			if val > K {
				K = val
				KIndex = i
			}
		}

		for i := KIndex + 1; i < length; i++ {
			val := int(runes[i] - '0')
			if val > L {
				L = val
			}
		}

		number := A*100000000000 +
			B*10000000000 +
			C*1000000000 +
			D*100000000 +
			E*10000000 +
			F*1000000 +
			G*100000 +
			H*10000 +
			I*1000 +
			J*100 +
			K*10 + L

		Answer += number

		A, B, C, D, E, F, G, H, I, J, K, L = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
		AIndex, BIndex, CIndex, DIndex, EIndex, FIndex, GIndex, HIndex, IIndex, JIndex, KIndex = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	}
}
