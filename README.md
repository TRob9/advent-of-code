# Advent of Code

My solutions for Advent of Code challenges.

## Quick Start (For Others Using This Repo)

1. Clone the repository
2. Run `make new` to create fresh boilerplate (archives my solutions)
3. Add your session cookie to `.session` file (see [Auto-Submission Setup](#auto-submission-setup))
4. Start the auto-fetch server: `./start_server.command`
5. Start solving: `make test-d1p1`

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
2. **Implement solution** in `part1.go` or `part2.go`
3. **Test and submit:**
   ```bash
   make test-d1p1   # Runs tests, then auto-submits if they pass
   ```

### Testing vs Running

```bash
# Run tests + auto-submit (if tests pass)
make test-d1p1

# Just run the solution (skip tests, skip submission)
make run-d1p1
```

Use `make run` when iterating on your solution. Use `make test` when ready to submit.

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
2. Delete the `2025/` directory
3. Create fresh boilerplate for all 12 days
4. Create `.session.example` file if `.session` doesn't exist

Perfect for:
- Starting a new year
- Letting others use your setup
- Resetting to clean slate

Your old solutions are safely archived!

## Benchmarks

```bash
# Benchmark individual solutions
make bench-d1p1   # Day 1, Part 1
make bench-d3p2   # Day 3, Part 2

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
