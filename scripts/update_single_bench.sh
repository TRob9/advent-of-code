#!/bin/bash
# Update README with a single benchmark result
# Usage: ./scripts/update_single_bench.sh <day> <part> <benchmark_output>

set -euo pipefail

DAY=$1
PART=$2
BENCH_OUTPUT="$3"

# Parse ns/op from benchmark output
# Example: BenchmarkSolution-10    8533    140170 ns/op    8192 B/op    2 allocs/op
NS_PER_OP=$(echo "$BENCH_OUTPUT" | grep -oE '[0-9.]+ ns/op' | grep -oE '[0-9.]+' | head -1)

if [ -z "$NS_PER_OP" ]; then
    echo "❌ Could not parse benchmark result"
    exit 1
fi

# Convert to appropriate unit
format_unit() {
    local ns=$1
    if (( $(echo "$ns < 1000" | bc -l) )); then
        printf "%.2f ns/op" "$ns"
    elif (( $(echo "$ns < 1000000" | bc -l) )); then
        printf "%.2f µs/op" "$(echo "$ns / 1000" | bc -l)"
    elif (( $(echo "$ns < 1000000000" | bc -l) )); then
        printf "%.2f ms/op" "$(echo "$ns / 1000000" | bc -l)"
    else
        printf "%.2f s/op" "$(echo "$ns / 1000000000" | bc -l)"
    fi
}

FORMATTED=$(format_unit "$NS_PER_OP")

echo "📊 Day $DAY Part $PART: $FORMATTED"

# Update README.md
README="README.md"

if [ ! -f "$README" ]; then
    echo "❌ README.md not found"
    exit 1
fi

# Use sed/awk to update the specific row (use @ as delimiter to avoid / in units)
# First check if the row exists
if grep -q "^| $DAY " "$README"; then
    # Row exists, update it
    if [ "$PART" -eq 1 ]; then
        # Update Part 1
        sed -i '' "s@^| $DAY   | [^|]* |@| $DAY   | ⭐ $FORMATTED |@" "$README"
    else
        # Update Part 2
        sed -i '' "s@^| $DAY   | \([^|]*\) | [^|]* |@| $DAY   | \1 | ⭐ $FORMATTED |@" "$README"
    fi
else
    # Row doesn't exist, add it
    # Find the line with the table separator and add after it
    if [ "$PART" -eq 1 ]; then
        NEW_ROW="| $DAY   | ⭐ $FORMATTED | - |"
    else
        NEW_ROW="| $DAY   | - | ⭐ $FORMATTED |"
    fi

    # Insert after the separator line
    sed -i '' "/^|-----|--------|--------|$/a\\
$NEW_ROW
" "$README"
fi

echo "✅ Updated README.md"
