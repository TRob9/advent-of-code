# Advent of Code

My solutions for Advent of Code challenges.

## Quick Start (For Others Using This Repo)

1. Clone the repository
2. Run `make new` to create fresh boilerplate (archives my solutions)
3. Add your session cookie to `.session` file (see [Auto-Submission Setup](#auto-submission-setup))
4. Start the auto-fetch server: `./start_server.command`
5. Write your solution in the `solve()` function
6. Run it: `make run-dXpY` (just runs your code)
7. Test it: `make test-dXpY` (runs tests, auto-submits if they pass)

Where X is the day number and Y is the part (1 or 2). Example: `make test-d1p1`

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

## Auto-Fetch Server

The auto-fetch server handles everything automatically:

**Start the server:**
```bash
# Double-click on macOS
./start_server.command

# Or run manually
cd server && node server.js
```

The server will:
- ⏰ Auto-fetch at 4:00:05 PM AEDT when puzzles unlock
- 📥 Download problem description → `dayN/problem.md`
- 📥 Download your personal input → `dayN/input.txt`
- 🤖 Use Claude SDK to extract and populate test cases → `testcases.txt`
- 🔄 Auto-fetch Part 2 when you complete Part 1
- 💤 Run continuously in the background

**Manual fetch (server must be running):**
```bash
make fetch        # Fetch today's puzzle
make fetch-<day>  # Fetch specific day (e.g., make fetch-4)
```

Note: Manual fetch commands require the server to be running for Claude-powered test case extraction.

## Solving Workflow

1. **Wait for auto-fetch** (or run `make fetch-<day>`)
2. **Implement your solution** in the `solve()` function (`part1.go` or `part2.go`)
3. **Iterate with `make run-dXpY`** to see your answer without running tests
4. **Submit with `make test-dXpY`** when ready - runs tests, auto-submits if they pass

### Testing vs Running

```bash
# Just run your solution (no tests, no submission)
make run-dXpY

# Test + auto-submit (runs tests, submits if they pass)
make test-dXpY
```

Where X is the day number and Y is the part (1 or 2). Examples: `make run-d1p1`, `make test-d12p2`

Use `make run-dXpY` when developing. Use `make test-dXpY` when ready to submit.

**Note:** Auto-submission is controlled by the `autoSubmit` constant in each solution file (default: `true`).

### How Testing Works

The test harness will:
1. ✅ Run your solution against test cases from `testcases.txt`
2. ✅ If tests pass, run against your real input from `input.txt`
3. ✅ Auto-submit the answer to Advent of Code

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

## Starting Fresh (make new)

The `make new` command archives existing solutions and creates fresh boilerplate:

```bash
make new
```

This will:
1. Archive current solutions to `archive/2025_TIMESTAMP.tar.gz`
2. Delete and recreate the `2025/` directory with fresh boilerplate for all 12 days
3. Create `.session.example` file if `.session` doesn't exist

Perfect for:
- Starting a new year
- Letting others use your setup
- Resetting to clean slate

Your old solutions are safely archived!

## Benchmarks

```bash
# Benchmark individual solutions (where X is day, Y is part 1 or 2)
make bench-dXpY   # Examples: make bench-d1p1, make bench-d3p2

# Run all benchmarks and update README table
make bench-all
```

Or run directly:

```bash
cd 2025/day1 && go test -tags=part1 -bench=BenchmarkPart1 -benchmem
cd benchmark/cmd && go run main.go  # Run all and update README
```

### 2025 Results (Apple M4 Pro)

| Day | Part 1 | Part 2 |
|-----|--------|--------|
| 1   | 187.12 µs/op | 171.11 µs/op |
| 2   | 101.64 ms/op | 602.38 ms/op |
| 3   | 52.70 µs/op | 66.46 µs/op |










## Language

Solutions are written in Go.
