#!/bin/bash

# Script to manually fetch Advent of Code problem, input, and populate test cases
# Usage: ./scripts/fetch.sh [day]
# If no day is provided, uses current day of month

set -e

YEAR=2025
DAY=${1:-$(date +%-d)}  # Default to current day if not provided

# Validate day is 1-25
if [ "$DAY" -lt 1 ] || [ "$DAY" -gt 25 ]; then
    echo "Error: Invalid day: $DAY (must be 1-25)"
    exit 1
fi

DAY_DIR="${YEAR}/day$(printf "%02d" $DAY)"

# Check if day directory exists
if [ ! -d "$DAY_DIR" ]; then
    echo "Error: Day directory not found: $DAY_DIR"
    echo "   Run 'make new' to create boilerplate first"
    exit 1
fi

# Check for session cookie
SESSION=$(cat .session 2>/dev/null | tr -d '\n' || echo "")
if [ -z "$SESSION" ] || [ "$SESSION" = "your_session_cookie_here" ]; then
    echo "Error: No session cookie found"
    echo "   Add your session cookie to .session file"
    exit 1
fi

echo "Fetching Day $DAY..."
echo ""

# Fetch problem HTML (temporary, will convert to markdown)
echo "Downloading problem..."
TEMP_HTML="/tmp/aoc_problem_${DAY}.html"
if ! curl -s -H "Cookie: session=$SESSION" \
    "https://adventofcode.com/${YEAR}/day/${DAY}" \
    -o "$TEMP_HTML"; then
    echo "Error: Failed to fetch problem"
    exit 1
fi

# Check if problem was fetched successfully
if grep -q "404 Not Found" "$TEMP_HTML" 2>/dev/null; then
    echo "Error: Problem not available yet (404)"
    rm "$TEMP_HTML"
    exit 1
fi

# Fetch input
echo "Downloading input..."
if ! curl -s -H "Cookie: session=$SESSION" \
    "https://adventofcode.com/${YEAR}/day/${DAY}/input" \
    -o "${DAY_DIR}/input.txt"; then
    echo "Error: Failed to fetch input"
    exit 1
fi

# Convert HTML to markdown
echo "Converting Converting to markdown..."
node -e "
const fs = require('fs');
const TurndownService = require('turndown');

const html = fs.readFileSync('${TEMP_HTML}', 'utf8');
const articleRegex = /<article[^>]*class=\"day-desc\"[^>]*>(.*?)<\/article>/gs;
const articles = [];
let match;
while ((match = articleRegex.exec(html)) !== null) {
    articles.push(match[1]);
}

const turndownService = new TurndownService({
    headingStyle: 'atx',
    codeBlockStyle: 'fenced'
});

let markdown = '';
articles.forEach((article, index) => {
    markdown += turndownService.turndown(article);
    if (index < articles.length - 1) {
        markdown += '\n\n---\n\n';
    }
});

fs.writeFileSync('${DAY_DIR}/problem.md', markdown);

// Output number of parts found for the shell script
console.log(articles.length);
" 2>/dev/null > /tmp/aoc_parts_count.txt

PARTS_COUNT=$(cat /tmp/aoc_parts_count.txt)
rm /tmp/aoc_parts_count.txt
rm "$TEMP_HTML"  # Clean up temp HTML file

if [ -z "$PARTS_COUNT" ] || [ "$PARTS_COUNT" = "0" ]; then
    echo "Error: Failed to parse problem description"
    exit 1
fi

echo "Success: Found $PARTS_COUNT part(s)"
echo ""

# Populate test cases with Claude
echo "AI: Populating test cases with Claude..."

# Check if Part 1 has already been populated (not empty)
PART1_POPULATED=$(grep -A 5 "^\*\*\* Part 1 \*\*\*" "${DAY_DIR}/testcases.txt" | grep -A 2 "^input:" | tail -n 1 | grep -v "^expected:" | grep -v "^$" | wc -l | tr -d ' ')

if [ "$PARTS_COUNT" = "1" ] || [ "$PART1_POPULATED" = "0" ]; then
    # Part 1 only OR Part 1 not yet populated (fresh fetch with both parts)
    if [ "$PARTS_COUNT" = "1" ]; then
        POPULATE_PARTS="Part 1 only (Part 2 will be added later when it unlocks)"
    else
        POPULATE_PARTS="Both Part 1 and Part 2"
    fi

    PROMPT="Extract the example test cases from this Advent of Code problem and write them to testcases.txt.

CRITICAL FORMAT REQUIREMENTS:
The testcases.txt file is parsed by automated testing. You MUST follow this EXACT format with NO deviations:

*** Part 1 ***
input:
<raw example input - no code blocks, no formatting, just the literal text>
expected:
<just the number or answer - nothing else>

$([ "$PARTS_COUNT" != "1" ] && echo "
*** Part 2 ***
input:
<raw example input - no code blocks, no formatting, just the literal text>
expected:
<just the number or answer - nothing else>
")

Rules:
1. Do NOT add markdown code blocks (\`\`\`) around the input
2. Do NOT add explanations, comments, or extra text
3. Do NOT add quotes around the input or expected value
4. Do NOT add any text before or after the format above
5. The input must be the EXACT text from the example (preserve spacing, newlines, etc.)
6. The expected value must be ONLY the answer (e.g., \"7\" not \"the answer is 7\")
$([ "$PARTS_COUNT" = "1" ] && echo "7. Leave the Part 2 section empty (it will be populated later when Part 2 unlocks)" || echo "7. If Part 2 uses a different example than Part 1, use the Part 2 example
8. If Part 2 uses the same example, the input may be identical but expected will differ")
9. If there are multiple examples, use the FIRST one shown in the problem

Task: Read the problem description from ${DAY_DIR}/problem.md and write the formatted test cases to ${DAY_DIR}/testcases.txt.
Populate: ${POPULATE_PARTS}"

else
    # Part 1 already populated, append Part 2 only
    PROMPT="TASK: Read problem.md and edit testcases.txt to add Part 2.

STEP 1: Read ${DAY_DIR}/problem.md to find the Part 2 example test case.

STEP 2: Use the Edit tool to replace the empty Part 2 section in ${DAY_DIR}/testcases.txt.

CRITICAL: The file already has Part 1 populated. You MUST use the Edit tool to replace ONLY the empty Part 2 section (from \"*** Part 2 ***\" to the end of the file).

Required format for Part 2:
*** Part 2 ***
input:
<raw example input - no code blocks, no formatting, just the literal text>
expected:
<just the number or answer - nothing else>

Edit Requirements:
1. Use Edit tool, NOT Write tool (preserve Part 1)
2. old_string must be the current empty Part 2 section
3. new_string must be the populated Part 2 section
4. Do NOT add markdown code blocks (\`\`\`) around the input
5. Do NOT add explanations, comments, or extra text
6. The input must be the EXACT text from the example (preserve spacing, newlines, etc.)
7. The expected value must be ONLY the answer (e.g., \"7\" not \"the answer is 7\")
8. If Part 2 uses the same example as Part 1, the input will be identical but expected will differ

Execute this task now by reading problem.md and editing testcases.txt with the Edit tool."

fi

# Call Claude SDK via the server's populate endpoint
if ! curl -s -X POST http://localhost:3030/populate \
    -H "Content-Type: application/json" \
    -d "{\"day\": $DAY, \"part\": $([ "$PARTS_COUNT" = "1" ] && echo 1 || echo 2), \"prompt\": $(echo "$PROMPT" | jq -Rs .)}" \
    > /dev/null 2>&1; then
    echo "Warning: Server not running - trying direct Claude call..."

    # Fallback: manual update instructions
    echo ""
    echo "Manual: Manual instructions:"
    echo "   1. Read ${DAY_DIR}/problem.md"
    echo "   2. Extract example test case(s)"
    echo "   3. Update ${DAY_DIR}/testcases.txt with the format shown above"
    exit 0
fi

echo "Success: Test cases populated!"
echo ""
echo "Complete: Fetch complete for Day $DAY!"
echo ""
echo "Next steps:"
echo "  1. Review ${DAY_DIR}/testcases.txt"
echo "  2. Start solving in ${DAY_DIR}/part1.go"
echo "  3. Run 'make test-d${DAY}p1' to test and submit"
