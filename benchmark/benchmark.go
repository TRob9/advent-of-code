package benchmark

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// Result represents benchmark results for a single part
type Result struct {
	Day     int
	Part    int
	NsPerOp float64
	Unit    string
}

// RunAll runs benchmarks for all days and parts, returning results
func RunAll(year int, maxDay int) ([]Result, error) {
	var results []Result

	for day := 1; day <= maxDay; day++ {
		for part := 1; part <= 2; part++ {
			result, err := runBenchmark(year, day, part)
			if err != nil {
				// Skip if benchmark doesn't exist yet
				continue
			}
			results = append(results, result)
		}
	}

	return results, nil
}

// runBenchmark runs a single benchmark and parses the result
func runBenchmark(year, day, part int) (Result, error) {
	dayDir := fmt.Sprintf("%d/day%02d", year, day)

	// Skip if solution is still using placeholder
	if isPlaceholder(year, day, part) {
		return Result{}, fmt.Errorf("solution is still using placeholder")
	}

	tag := fmt.Sprintf("part%d", part)

	// Get current directory and navigate to repo root
	currentDir, _ := os.Getwd()
	// If we're in benchmark/cmd, go up two levels to repo root
	repoRoot := currentDir
	if strings.HasSuffix(currentDir, "benchmark/cmd") {
		repoRoot = strings.TrimSuffix(currentDir, "/benchmark/cmd")
	}

	fullDayDir := fmt.Sprintf("%s/%s", repoRoot, dayDir)
	testFile := fmt.Sprintf("%s/benchmark_test.go", fullDayDir)
	templateFile := fmt.Sprintf("%s/benchmark/benchmark_test.template", repoRoot)

	// Copy template file
	templateContent, err := os.ReadFile(templateFile)
	if err != nil {
		return Result{}, fmt.Errorf("failed to read template: %w", err)
	}
	if err := os.WriteFile(testFile, templateContent, 0644); err != nil {
		return Result{}, fmt.Errorf("failed to write test file: %w", err)
	}
	// Clean up test file when done
	defer os.Remove(testFile)

	cmd := exec.Command("go", "test", "-tags="+tag, "-bench=.", "-benchmem", "-benchtime=1s")
	cmd.Dir = fullDayDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		return Result{}, err
	}

	// Parse benchmark output
	// Example: BenchmarkSolution-10    8533    140170 ns/op    8192 B/op    2 allocs/op
	re := regexp.MustCompile(`BenchmarkSolution\-\d+\s+\d+\s+([\d.]+)\s+ns/op`)
	matches := re.FindStringSubmatch(string(output))

	if len(matches) < 2 {
		return Result{}, fmt.Errorf("could not parse benchmark output")
	}

	nsPerOp, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return Result{}, err
	}

	return Result{
		Day:     day,
		Part:    part,
		NsPerOp: nsPerOp,
		Unit:    formatUnit(nsPerOp),
	}, nil
}

// isPlaceholder checks if a solution file still contains the placeholder code
func isPlaceholder(year, day, part int) bool {
	// Get current directory and navigate to repo root
	currentDir, _ := os.Getwd()
	repoRoot := currentDir
	if strings.HasSuffix(currentDir, "benchmark/cmd") {
		repoRoot = strings.TrimSuffix(currentDir, "/benchmark/cmd")
	}

	fileName := fmt.Sprintf("%s/%d/day%02d/part%d.go", repoRoot, year, day, part)
	content, err := os.ReadFile(fileName)
	if err != nil {
		return true // If we can't read the file, skip it
	}

	// Check for the placeholder comment
	return strings.Contains(string(content), "// Placeholder - replace with your solution logic")
}

// formatUnit converts nanoseconds to appropriate unit
func formatUnit(ns float64) string {
	switch {
	case ns < 1000:
		return fmt.Sprintf("%.2f ns/op", ns)
	case ns < 1000000:
		return fmt.Sprintf("%.2f µs/op", ns/1000)
	case ns < 1000000000:
		return fmt.Sprintf("%.2f ms/op", ns/1000000)
	default:
		return fmt.Sprintf("%.2f s/op", ns/1000000000)
	}
}

// UpdateReadme updates the README with benchmark results
func UpdateReadme(readmePath string, year int, results []Result) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}

	// Build the table
	var table strings.Builder
	table.WriteString(fmt.Sprintf("### %d Results (Apple M4 Pro)\n\n", year))
	table.WriteString("| Day | Part 1 | Part 2 |\n")
	table.WriteString("|-----|--------|--------|\n")

	// Group by day
	dayResults := make(map[int]map[int]string)
	for _, r := range results {
		if dayResults[r.Day] == nil {
			dayResults[r.Day] = make(map[int]string)
		}
		dayResults[r.Day][r.Part] = r.Unit
	}

	// Write table rows
	for day := 1; day <= 25; day++ {
		if dayResults[day] == nil {
			continue
		}
		part1 := dayResults[day][1]
		part2 := dayResults[day][2]

		// Add gold stars for completed parts
		if part1 == "" {
			part1 = "-"
		} else {
			part1 = "⭐ " + part1
		}
		if part2 == "" {
			part2 = "-"
		} else {
			part2 = "⭐ " + part2
		}
		table.WriteString(fmt.Sprintf("| %d   | %s | %s |\n", day, part1, part2))
	}

	// Replace the benchmark section
	re := regexp.MustCompile(`(?s)### \d+ Results.*?\n\n(\|.*?\n)*`)
	newContent := re.ReplaceAllString(string(content), table.String()+"\n")

	return os.WriteFile(readmePath, []byte(newContent), 0644)
}
