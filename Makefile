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

# Benchmark targets - pattern rules to reduce duplication
bench-d%p1:
	@cp benchmark/benchmark_test.template $(YEAR)/day$(shell printf "%02d" $*)/benchmark_test.go
	@OUTPUT=$$(cd $(YEAR)/day$(shell printf "%02d" $*) && go test -tags=part1 -bench=. -benchmem 2>&1) && \
		echo "$$OUTPUT" && \
		./scripts/update_single_bench.sh $* 1 "$$OUTPUT"
	@rm $(YEAR)/day$(shell printf "%02d" $*)/benchmark_test.go

bench-d%p2:
	@cp benchmark/benchmark_test.template $(YEAR)/day$(shell printf "%02d" $*)/benchmark_test.go
	@OUTPUT=$$(cd $(YEAR)/day$(shell printf "%02d" $*) && go test -tags=part2 -bench=. -benchmem 2>&1) && \
		echo "$$OUTPUT" && \
		./scripts/update_single_bench.sh $* 2 "$$OUTPUT"
	@rm $(YEAR)/day$(shell printf "%02d" $*)/benchmark_test.go

# Run all benchmarks and update README
.PHONY: bench-all
bench-all:
	@echo "Running all benchmarks..."
	@cd benchmark/cmd && go run main.go

# Lint - clean up code
.PHONY: lint
lint:
	@echo "🧹 Cleaning print statements from solutions..."
	@./scripts/clean_prints.sh $(YEAR)
