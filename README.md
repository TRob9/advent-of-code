# Advent of Code

My solutions for Advent of Code challenges.

## Note

These solutions are preserved exactly as they were when I first solved each puzzle. I stopped working on them the moment I got the correct answer. They are unoptimised and most certainly are inefficient, as I wish to document my progress authentically as I grow as an engineer.

I'll come back to these later when I'm more experienced and refactor them to see how much I've improved.

## Running Solutions

Each solution can be run independently with Go:

```bash
go run part-1.go
go run part-2.go
```

## Benchmarks

Run benchmarks for each day:

```bash
# Part 1
cd 2025/day1 && go test -tags=part1 -bench=BenchmarkPart1 -benchmem

# Part 2
cd 2025/day1 && go test -tags=part2 -bench=BenchmarkPart2 -benchmem
```

### 2025 Results (Apple M4 Pro)

| Day | Part 1 | Part 2 |
|-----|--------|--------|
| 1   | 140.17 µs/op | 123.51 µs/op |
| 2   | 96.39 ms/op | 580.88 ms/op |
| 3   | 47.74 µs/op | 60.66 µs/op |

## Language

Solutions are written in Go.
