#!/bin/bash
# Watch for file changes in day directories

YEAR=2025
DAY=$(date +%-d)
DAY_DIR="$HOME/Source/Personal-Projects/advent-of-code/${YEAR}/day$(printf "%02d" $DAY)"

echo "👀 Watching $DAY_DIR for changes..."
echo ""

# Use fswatch if available, otherwise fall back to stat polling
if command -v fswatch &> /dev/null; then
    fswatch -0 "$DAY_DIR" | while read -d "" event; do
        echo "[$(date +%H:%M:%S)] File changed: $(basename "$event")"
        if [[ "$event" == *.md ]]; then
            echo "  📄 problem.md updated"
        elif [[ "$event" == *.txt ]] && [[ "$event" != *testcases.txt ]]; then
            echo "  📥 input.txt updated"
        elif [[ "$event" == *testcases.txt ]]; then
            echo "  ✅ testcases.txt updated"
        fi
    done
else
    echo "Install fswatch for real-time monitoring: brew install fswatch"
    echo "Falling back to polling every 1 second..."
    
    while true; do
        if [[ -f "$DAY_DIR/problem.md" ]]; then
            echo "[$(date +%H:%M:%S)] ✅ problem.md exists"
        fi
        if [[ -f "$DAY_DIR/input.txt" ]]; then
            echo "[$(date +%H:%M:%S)] ✅ input.txt exists"
        fi
        if [[ -f "$DAY_DIR/testcases.txt" ]]; then
            echo "[$(date +%H:%M:%S)] ✅ testcases.txt exists"
        fi
        sleep 1
    done
fi
