package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/trob9/advent-of-code/benchmark"
)

func main() {
	year := flag.Int("year", 2025, "Year to run benchmarks for")
	maxDay := flag.Int("days", 25, "Maximum day to benchmark")
	updateReadme := flag.Bool("update-readme", true, "Update README with results")
	flag.Parse()

	fmt.Printf("Running benchmarks for year %d (days 1-%d)...\n", *year, *maxDay)

	results, err := benchmark.RunAll(*year, *maxDay)
	if err != nil {
		log.Fatalf("Error running benchmarks: %v", err)
	}

	fmt.Printf("\nCompleted %d benchmarks\n\n", len(results))

	// Print results
	fmt.Println("Results:")
	for _, r := range results {
		fmt.Printf("Day %d Part %d: %s\n", r.Day, r.Part, r.Unit)
	}

	// Update README if requested
	if *updateReadme {
		readmePath := filepath.Join("..", "..", "README.md")
		if err := benchmark.UpdateReadme(readmePath, *year, results); err != nil {
			log.Fatalf("Error updating README: %v", err)
		}
		fmt.Println("\nREADME.md updated successfully!")
	}
}
