# Advent of Code

My solutions for Advent of Code challenges.

## Quick Start (For Others Using This Repo)

1. Clone the repository
2. Run `make new` to create fresh boilerplate (archives my solutions)
3. Add your session cookie to `.session` file
4. Start solving: `make test-d1p1`

## Note

These solutions are preserved exactly as they were when I first solved each puzzle. I stopped working on them the moment I got the correct answer. They are unoptimised and most certainly are inefficient, as I wish to document my progress authentically as I grow as an engineer.

I'll come back to these later when I'm more experienced and refactor them to see how much I've improved.

## Structure

Each day follows a consistent pattern:

```
2025/
└── day1/
    ├── part1.go      # Part 1 solution
    ├── part2.go      # Part 2 solution
    ├── input.txt     # Your puzzle input
    ├── testcases.txt # Test cases + expected outputs
    └── problem.md    # Problem description (auto-converted from HTML)
```

Note: `benchmark_test.go` is generated on-demand when running benchmarks, not stored in the repo.

## Solving a New Puzzle

### Option 1: Manual (Fastest)

1. Navigate to the day's directory (e.g., `cd 2025/day4`)
2. Paste example + expected output into `testcases.txt`:
   ```
   *** Part 1 ***
   input:
   <paste example input from problem>
   expected:
   <expected answer from problem>
   ```
3. Paste your puzzle input into `input.txt`
4. Implement the `solve()` function in `part1.go`
5. Test and auto-submit: `make test-d4p1`

### Option 2: Auto-Fetch with Claude

```bash
make fetch-4  # Fetches problem, input, and populates testcases.txt
```

This will:
- 📥 Download problem description and input
- 🤖 Use Claude to extract test cases automatically
- 📝 Populate `testcases.txt` in the correct format

Then just:
1. Implement the `solve()` function in `part1.go`
2. Test and auto-submit: `make test-d4p1`

### How Testing Works

The harness will:
- ✅ Run your solution against the test case
- ✅ If test passes, run against real input
- ✅ Auto-submit answer (if session cookie is set)

## Auto-Submission Setup

To enable auto-submission:

1. Log in to [adventofcode.com](https://adventofcode.com)
2. Open DevTools (F12) → Application/Storage → Cookies
3. Copy the `session` cookie value
4. Create a `.session` file in the project root:
   ```bash
   echo "your_session_cookie_here" > .session
   ```

The `.session` file is git-ignored for security.

Now `make test-dXpY` will auto-submit if tests pass!

## Auto-Fetch Server (Optional)

Want problems and inputs automatically downloaded at unlock time (4:00 PM EST)?

**Start the server:**
```bash
# Double-click on macOS
./start_server.command

# Or run manually
cd server && node server.js
```

The server will:
- ⏰ Wait until 4:00:05 PM each day
- 📥 Auto-fetch problem description → `dayN/problem.md`
- 📥 Auto-fetch your personal input → `dayN/input.txt`
- 🤖 Use Claude SDK to populate `testcases.txt` automatically
- 🔄 Auto-fetch Part 2 when Part 1 completes successfully
- 💤 Run continuously in the background

Perfect for getting started as soon as puzzles unlock! You can also use `make fetch` or `make fetch-<day>` to manually fetch any day.

## Starting Fresh (make new)

The `make new` command archives existing solutions and creates fresh boilerplate:

```bash
make new
```

This will:
1. Archive current solutions to `archive/2025_TIMESTAMP.tar.gz`
2. Delete the `2025/` directory
3. Create fresh boilerplate for all 12 days
4. Create `.session` file if it doesn't exist

Perfect for:
- Starting a new year
- Letting others use your setup
- Resetting to clean slate

Your old solutions are safely archived!

## Running Solutions

Using the Makefile:

```bash
# Run solutions
make run-d1p1   # Day 1, Part 1
make run-d1p2   # Day 1, Part 2
make run-d12p2  # Day 12, Part 2
```

Or run directly with Go:

```bash
cd 2025/day1 && go run -tags=part1 part1.go
cd 2025/day1 && go run -tags=part2 part2.go
```

## Benchmarks

Using the Makefile:

```bash
# Benchmark individual solutions
make bench-d1p1   # Day 1, Part 1
make bench-d3p2   # Day 3, Part 2

# Run all benchmarks and update README table
make bench-all
```

Or run directly with Go:

```bash
cd 2025/day1 && go test -tags=part1 -bench=BenchmarkPart1 -benchmem
cd 2025/day1 && go test -tags=part2 -bench=BenchmarkPart2 -benchmem

# Run all and update README
cd benchmark/cmd && go run main.go
```

### 2025 Results (Apple M4 Pro)

| Day | Part 1 | Part 2 |
|-----|--------|--------|
| 1   | 187.12 µs/op | 171.11 µs/op |
| 2   | 101.64 ms/op | 602.38 ms/op |
| 3   | 52.70 µs/op | 66.46 µs/op |









## Language

Solutions are written in Go.
