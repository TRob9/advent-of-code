#!/bin/bash
# Remove fmt.Println statements from solve functions in Advent of Code solutions
# This prevents benchmark output pollution

set -euo pipefail

YEAR="${1:-2025}"
FIXED_COUNT=0

# Find all part*.go files
for file in "$YEAR"/day*/part*.go; do
    if [ ! -f "$file" ]; then
        continue
    fi

    # Check if file has fmt.Println in solve function
    if grep -q "fmt\.Println" "$file"; then
        # Create a backup
        cp "$file" "$file.bak"

        # Remove lines containing fmt.Println (including multiline calls)
        # This handles both single-line and multi-line fmt.Println calls
        sed -i '' '/fmt\.Println/d' "$file"

        # Also remove lines with just fmt.Printf if they're printing Answer
        sed -i '' '/fmt\.Printf.*Answer/d' "$file"

        echo "✅ Cleaned $file"
        FIXED_COUNT=$((FIXED_COUNT + 1))

        # Remove backup if successful
        rm "$file.bak"
    fi
done

if [ $FIXED_COUNT -eq 0 ]; then
    echo "✨ No print statements found - code is clean!"
else
    echo ""
    echo "🧹 Cleaned $FIXED_COUNT file(s)"
    echo "   Removed fmt.Println statements that interfere with benchmarks"
fi

exit 0
