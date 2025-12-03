YEAR := 2025

.PHONY: new
new:
	@./scripts/new-setup.sh $(YEAR)

# Fetch targets
.PHONY: fetch
fetch:
	@./scripts/fetch.sh

fetch-%:
	@./scripts/fetch.sh $*

# Run targets (skip tests, just run solution)
run-d%p1:
	@cd $(YEAR)/day$(shell printf "%02d" $*) && SKIP_TESTS=1 SKIP_SUBMIT=1 go run -tags=part1 part1.go

run-d%p2:
	@cd $(YEAR)/day$(shell printf "%02d" $*) && SKIP_TESTS=1 SKIP_SUBMIT=1 go run -tags=part2 part2.go

# Test targets (runs tests + auto-submits if they pass)
test-d%p1:
	@cd $(YEAR)/day$(shell printf "%02d" $*) && go run -tags=part1 part1.go

test-d%p2:
	@cd $(YEAR)/day$(shell printf "%02d" $*) && go run -tags=part2 part2.go

# Benchmark targets
bench-d1p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day01/benchmark_test.go
	@cd $(YEAR)/day01 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day01/benchmark_test.go

bench-d1p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day01/benchmark_test.go
	@cd $(YEAR)/day01 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day01/benchmark_test.go

bench-d2p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day02/benchmark_test.go
	@cd $(YEAR)/day02 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day02/benchmark_test.go

bench-d2p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day02/benchmark_test.go
	@cd $(YEAR)/day02 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day02/benchmark_test.go

bench-d3p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day03/benchmark_test.go
	@cd $(YEAR)/day03 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day03/benchmark_test.go

bench-d3p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day03/benchmark_test.go
	@cd $(YEAR)/day03 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day03/benchmark_test.go

bench-d4p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day04/benchmark_test.go
	@cd $(YEAR)/day04 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day04/benchmark_test.go

bench-d4p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day04/benchmark_test.go
	@cd $(YEAR)/day04 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day04/benchmark_test.go

bench-d5p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day05/benchmark_test.go
	@cd $(YEAR)/day05 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day05/benchmark_test.go

bench-d5p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day05/benchmark_test.go
	@cd $(YEAR)/day05 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day05/benchmark_test.go

bench-d6p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day06/benchmark_test.go
	@cd $(YEAR)/day06 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day06/benchmark_test.go

bench-d6p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day06/benchmark_test.go
	@cd $(YEAR)/day06 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day06/benchmark_test.go

bench-d7p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day07/benchmark_test.go
	@cd $(YEAR)/day07 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day07/benchmark_test.go

bench-d7p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day07/benchmark_test.go
	@cd $(YEAR)/day07 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day07/benchmark_test.go

bench-d8p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day08/benchmark_test.go
	@cd $(YEAR)/day08 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day08/benchmark_test.go

bench-d8p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day08/benchmark_test.go
	@cd $(YEAR)/day08 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day08/benchmark_test.go

bench-d9p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day09/benchmark_test.go
	@cd $(YEAR)/day09 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day09/benchmark_test.go

bench-d9p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day09/benchmark_test.go
	@cd $(YEAR)/day09 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day09/benchmark_test.go

bench-d10p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day10/benchmark_test.go
	@cd $(YEAR)/day10 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day10/benchmark_test.go

bench-d10p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day10/benchmark_test.go
	@cd $(YEAR)/day10 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day10/benchmark_test.go

bench-d11p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day11/benchmark_test.go
	@cd $(YEAR)/day11 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day11/benchmark_test.go

bench-d11p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day11/benchmark_test.go
	@cd $(YEAR)/day11 && go test -tags=part2 -benchmem
	@rm $(YEAR)/day11/benchmark_test.go

bench-d12p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day12/benchmark_test.go
	@cd $(YEAR)/day12 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day12/benchmark_test.go

bench-d12p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day12/benchmark_test.go
	@cd $(YEAR)/day12 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day12/benchmark_test.go

# Run all benchmarks and update README
.PHONY: bench-all
bench-all:
	@echo "Running all benchmarks..."
	@cd benchmark/cmd && go run main.go
