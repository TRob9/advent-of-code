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

# Run targets
run-d%p1:
	@cd $(YEAR)/day$* && go run -tags=part1 part1.go

run-d%p2:
	@cd $(YEAR)/day$* && go run -tags=part2 part2.go

# Test targets (runs tests + auto-submits if they pass)
test-d%p1:
	@cd $(YEAR)/day$* && go run -tags=part1 part1.go

test-d%p2:
	@cd $(YEAR)/day$* && go run -tags=part2 part2.go

# Benchmark targets
bench-d1p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day1/benchmark_test.go
	@cd $(YEAR)/day1 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day1/benchmark_test.go

bench-d1p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day1/benchmark_test.go
	@cd $(YEAR)/day1 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day1/benchmark_test.go

bench-d2p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day2/benchmark_test.go
	@cd $(YEAR)/day2 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day2/benchmark_test.go

bench-d2p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day2/benchmark_test.go
	@cd $(YEAR)/day2 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day2/benchmark_test.go

bench-d3p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day3/benchmark_test.go
	@cd $(YEAR)/day3 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day3/benchmark_test.go

bench-d3p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day3/benchmark_test.go
	@cd $(YEAR)/day3 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day3/benchmark_test.go

bench-d4p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day4/benchmark_test.go
	@cd $(YEAR)/day4 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day4/benchmark_test.go

bench-d4p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day4/benchmark_test.go
	@cd $(YEAR)/day4 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day4/benchmark_test.go

bench-d5p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day5/benchmark_test.go
	@cd $(YEAR)/day5 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day5/benchmark_test.go

bench-d5p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day5/benchmark_test.go
	@cd $(YEAR)/day5 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day5/benchmark_test.go

bench-d6p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day6/benchmark_test.go
	@cd $(YEAR)/day6 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day6/benchmark_test.go

bench-d6p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day6/benchmark_test.go
	@cd $(YEAR)/day6 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day6/benchmark_test.go

bench-d7p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day7/benchmark_test.go
	@cd $(YEAR)/day7 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day7/benchmark_test.go

bench-d7p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day7/benchmark_test.go
	@cd $(YEAR)/day7 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day7/benchmark_test.go

bench-d8p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day8/benchmark_test.go
	@cd $(YEAR)/day8 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day8/benchmark_test.go

bench-d8p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day8/benchmark_test.go
	@cd $(YEAR)/day8 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day8/benchmark_test.go

bench-d9p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day9/benchmark_test.go
	@cd $(YEAR)/day9 && go test -tags=part1 -bench=. -benchmem
	@rm $(YEAR)/day9/benchmark_test.go

bench-d9p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day9/benchmark_test.go
	@cd $(YEAR)/day9 && go test -tags=part2 -bench=. -benchmem
	@rm $(YEAR)/day9/benchmark_test.go

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
