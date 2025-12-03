#!/bin/bash

# Script to archive existing solutions and create fresh boilerplate
# Usage: ./scripts/new-setup.sh <year>

set -e

YEAR=${1:-2025}
ARCHIVE_DIR="archive"
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

echo "Archiving existing solutions..."

# Create archive directory if it doesn't exist
mkdir -p "$ARCHIVE_DIR"

# Archive the year's directory if it exists
if [ -d "$YEAR" ]; then
    ARCHIVE_NAME="${ARCHIVE_DIR}/${YEAR}_${TIMESTAMP}.tar.gz"
    tar -czf "$ARCHIVE_NAME" "$YEAR"
    echo "Success: Archived to $ARCHIVE_NAME"

    # Remove old directory
    rm -rf "$YEAR"
    echo "Success: Removed old $YEAR directory"
else
    echo "Info: No existing $YEAR directory to archive"
fi

echo ""
echo "Creating fresh boilerplate for $YEAR..."

# Create year directory
mkdir -p "$YEAR"

# Create 12 days of boilerplate
for day in {1..12}; do
    DAY_DIR="$YEAR/day$day"
    mkdir -p "$DAY_DIR"

    # Create part1.go
    cat > "$DAY_DIR/part1.go" << 'EOF'
//go:build part1

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
		opts = append(opts, harness.WithSubmit(YEAR, DAY))
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
	// Parse input - uncomment the pattern you need, delete the rest

	// a) Array of strings (lines)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// b) Array of strings (comma-separated)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), ",")

	// c) Array of integers (lines)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([]int, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i], _ = strconv.Atoi(line)
	// }

	// d) Array of integers (comma-separated)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make([]int, len(parts))
	// for i, part := range parts {
	// 	cleanedInput[i], _ = strconv.Atoi(strings.TrimSpace(part))
	// }

	// e) Map of string to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[string]int)
	// for _, part := range parts {
	// 	cleanedInput[strings.TrimSpace(part)] = 0
	// }

	// f) Map of int to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[int]int)
	// for _, part := range parts {
	// 	key, _ := strconv.Atoi(strings.TrimSpace(part))
	// 	cleanedInput[key] = 0
	// }

	// g) 2D grid of characters
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]rune, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i] = []rune(line)
	// }

	// h) 2D grid of integers (space-separated)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]int, len(lines))
	// for i, line := range lines {
	// 	parts := strings.Fields(line)
	// 	cleanedInput[i] = make([]int, len(parts))
	// 	for j, part := range parts {
	// 		cleanedInput[i][j], _ = strconv.Atoi(part)
	// 	}
	// }

	// i) Single integer
	// cleanedInput, _ := strconv.Atoi(strings.TrimSpace(string(input)))

	// j) Single string
	// cleanedInput := strings.TrimSpace(string(input))

	Answer = 0 // Placeholder - replace with your solution logic
}
EOF
    sed -i '' "s/YEAR/$YEAR/g" "$DAY_DIR/part1.go"
    sed -i '' "s/DAY/$day/g" "$DAY_DIR/part1.go"

    # Create part2.go
    cat > "$DAY_DIR/part2.go" << 'EOF'
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
		opts = append(opts, harness.WithSubmit(YEAR, DAY))
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
	// Parse input - uncomment the pattern you need, delete the rest

	// a) Array of strings (lines)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), "\n")

	// b) Array of strings (comma-separated)
	// cleanedInput := strings.Split(strings.TrimSpace(string(input)), ",")

	// c) Array of integers (lines)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([]int, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i], _ = strconv.Atoi(line)
	// }

	// d) Array of integers (comma-separated)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make([]int, len(parts))
	// for i, part := range parts {
	// 	cleanedInput[i], _ = strconv.Atoi(strings.TrimSpace(part))
	// }

	// e) Map of string to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[string]int)
	// for _, part := range parts {
	// 	cleanedInput[strings.TrimSpace(part)] = 0
	// }

	// f) Map of int to int (comma-separated keys)
	// parts := strings.Split(strings.TrimSpace(string(input)), ",")
	// cleanedInput := make(map[int]int)
	// for _, part := range parts {
	// 	key, _ := strconv.Atoi(strings.TrimSpace(part))
	// 	cleanedInput[key] = 0
	// }

	// g) 2D grid of characters
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]rune, len(lines))
	// for i, line := range lines {
	// 	cleanedInput[i] = []rune(line)
	// }

	// h) 2D grid of integers (space-separated)
	// lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	// cleanedInput := make([][]int, len(lines))
	// for i, line := range lines {
	// 	parts := strings.Fields(line)
	// 	cleanedInput[i] = make([]int, len(parts))
	// 	for j, part := range parts {
	// 		cleanedInput[i][j], _ = strconv.Atoi(part)
	// 	}
	// }

	// i) Single integer
	// cleanedInput, _ := strconv.Atoi(strings.TrimSpace(string(input)))

	// j) Single string
	// cleanedInput := strings.TrimSpace(string(input))

	Answer = 0 // Placeholder - replace with your solution logic
}
EOF
    sed -i '' "s/YEAR/$YEAR/g" "$DAY_DIR/part2.go"
    sed -i '' "s/DAY/$day/g" "$DAY_DIR/part2.go"

    # Benchmark file is generated on-demand from benchmark/benchmark_test.template
    # No need to create it in boilerplate

    # Create empty input.txt
    touch "$DAY_DIR/input.txt"

    # Create testcases.txt template
    cat > "$DAY_DIR/testcases.txt" << 'EOF'
*** Part 1 ***
input:


expected:


*** Part 2 ***
input:


expected:

EOF


    echo "Success: Created day$day"
done

echo ""
echo "Security: Creating .session file..."

if [ -f ".session" ]; then
    echo "Info: .session file already exists, skipping"
else
    cat > ".session" << 'EOF'
your_session_cookie_here
EOF
    echo "Success: Created .session file (add your session cookie)"
fi

echo ""
echo "Complete: Setup complete!"
echo ""
echo "Next steps:"
echo "  1. Add your session cookie to .session file"
echo "  2. Start solving puzzles!"
echo "  3. Use 'make test-d1p1' to test and submit"
