package harness

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// TestCase represents a single test case with input and expected output
type TestCase struct {
	Input    []byte
	Expected string
}

// Harness manages test execution and validation
type Harness struct {
	solveFn    func([]byte)
	answer     interface{}
	part       int
	year       int
	day        int
	autoSubmit bool
}

// Option is a functional option for configuring the Harness
type Option func(*Harness)

// WithSubmit enables auto-submission when tests pass
func WithSubmit(year, day int) Option {
	return func(h *Harness) {
		h.year = year
		h.day = day
		h.autoSubmit = true
	}
}

// New creates a new harness for the given solve function and part number
func New(solveFn func([]byte), answer interface{}, part int, opts ...Option) *Harness {
	h := &Harness{
		solveFn: solveFn,
		answer:  answer,
		part:    part,
	}
	for _, opt := range opts {
		opt(h)
	}
	return h
}

// RunTests reads the tests.txt file, runs all test cases for the specified part,
// and returns true if all tests pass
func (h *Harness) RunTests() (bool, error) {
	testCase, err := h.loadTestCase()
	if err != nil {
		return false, err
	}

	if testCase == nil {
		// No test case defined, skip testing
		return true, nil
	}

	// Run the solve function
	h.solveFn(testCase.Input)

	// Get the answer from the pointer
	var result string
	switch v := h.answer.(type) {
	case *int:
		result = strconv.Itoa(*v)
	case *string:
		result = *v
	default:
		return false, fmt.Errorf("unsupported answer type: %T", h.answer)
	}

	if result != testCase.Expected {
		return false, fmt.Errorf("test failed: expected %s, got %s", testCase.Expected, result)
	}

	fmt.Printf("✅ Part %d test passed (expected: %s, got: %s)\n", h.part, testCase.Expected, result)
	return true, nil
}

// loadTestCase reads the testcases.txt file and extracts the test case for the current part
func (h *Harness) loadTestCase() (*TestCase, error) {
	data, err := os.ReadFile("testcases.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil // No test file, skip
		}
		return nil, err
	}

	content := string(data)
	partMarker := fmt.Sprintf("*** Part %d ***", h.part)

	// Find the part section
	partIdx := strings.Index(content, partMarker)
	if partIdx == -1 {
		return nil, nil // Part not defined, skip
	}

	// Extract content after the part marker
	afterPart := content[partIdx+len(partMarker):]

	// Find the next part marker or end of file
	nextPartIdx := strings.Index(afterPart, "*** Part")
	if nextPartIdx != -1 {
		afterPart = afterPart[:nextPartIdx]
	}

	// Parse input and expected sections
	inputMarker := "input:"
	expectedMarker := "expected:"

	inputIdx := strings.Index(afterPart, inputMarker)
	expectedIdx := strings.Index(afterPart, expectedMarker)

	if inputIdx == -1 || expectedIdx == -1 {
		return nil, nil // Incomplete test case, skip
	}

	// Extract input
	inputSection := afterPart[inputIdx+len(inputMarker) : expectedIdx]
	input := strings.TrimSpace(inputSection)

	// Extract expected
	expectedSection := afterPart[expectedIdx+len(expectedMarker):]
	expected := strings.TrimSpace(expectedSection)

	return &TestCase{
		Input:    []byte(input),
		Expected: expected,
	}, nil
}

// Run executes the solution with the actual input.txt file
func (h *Harness) Run() error {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		return fmt.Errorf("failed to read input.txt: %v", err)
	}

	// Reset answer to zero value before running actual solution
	// (tests may have modified it)
	switch v := h.answer.(type) {
	case *int:
		*v = 0
	case *string:
		*v = ""
	}

	h.solveFn(input)

	// Get the answer from the pointer
	var result interface{}
	switch v := h.answer.(type) {
	case *int:
		result = *v
	case *string:
		result = *v
	default:
		return fmt.Errorf("unsupported answer type: %T", h.answer)
	}

	fmt.Printf("Part %d answer: %v\n", h.part, result)

	// Auto-submit if enabled (unless SKIP_SUBMIT is set)
	if h.autoSubmit && os.Getenv("SKIP_SUBMIT") == "" {
		return h.submitAnswer(result)
	}

	return nil
}

// getSessionCookie reads the session cookie from .session file or environment variable
func getSessionCookie() string {
	// Try reading from .session file, checking current dir and parent dirs
	dir, _ := os.Getwd()
	for dir != "/" {
		sessionFile := fmt.Sprintf("%s/.session", dir)
		if data, err := os.ReadFile(sessionFile); err == nil {
			return strings.TrimSpace(string(data))
		}
		// Move up one directory
		dir = fmt.Sprintf("%s/..", dir)
		dir, _ = filepath.Abs(dir)
	}

	// Fall back to environment variable
	return os.Getenv("AOC_SESSION")
}

// submitAnswer submits the answer to Advent of Code
func (h *Harness) submitAnswer(answer interface{}) error {
	session := getSessionCookie()
	if session == "" {
		fmt.Println("⚠️  Skipping submission: No session cookie found")
		fmt.Println("To enable auto-submission:")
		fmt.Println("  1. Go to adventofcode.com and log in")
		fmt.Println("  2. Open browser DevTools (F12) → Application → Cookies")
		fmt.Println("  3. Copy the 'session' cookie value")
		fmt.Println("  4. Paste it into a .session file in the project root")
		fmt.Println("  (Or set AOC_SESSION environment variable)")
		return nil
	}

	submitURL := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", h.year, h.day)

	// Prepare form data
	data := url.Values{}
	data.Set("level", strconv.Itoa(h.part))
	data.Set("answer", fmt.Sprint(answer))

	// Create request
	req, err := http.NewRequest("POST", submitURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", session))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to submit answer: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response
	bodyStr := string(body)

	if strings.Contains(bodyStr, "That's the right answer") {
		fmt.Println("🎉 Correct! Answer submitted successfully!")

		// For Part 1: copy part1.go to part2.go and notify server to fetch Part 2
		if h.part == 1 {
			// Copy part1.go to part2.go
			if err := h.copyPart1ToPart2(); err != nil {
				fmt.Printf("⚠️  Failed to copy part1.go to part2.go: %v\n", err)
			}

			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				notifyServerForPart2(h.day)
			}()
			wg.Wait()
		}

		return nil
	} else if strings.Contains(bodyStr, "That's not the right answer") {
		// Extract the message
		re := regexp.MustCompile(`<article><p>([^<]+)</p></article>`)
		if matches := re.FindStringSubmatch(bodyStr); len(matches) > 1 {
			fmt.Printf("❌ Wrong answer: %s\n", matches[1])
		} else {
			fmt.Println("❌ Wrong answer")
		}
		return fmt.Errorf("incorrect answer")
	} else if strings.Contains(bodyStr, "You gave an answer too recently") {
		fmt.Println("⏳ Please wait before submitting again")
		return fmt.Errorf("rate limited")
	} else if strings.Contains(bodyStr, "Did you already complete it") {
		fmt.Println("✓ Already completed")
		return nil
	} else {
		fmt.Println("⚠️  Unknown response from server")
		return nil
	}
}

// notifyServerForPart2 sends an HTTP request to the local server to trigger Part 2 fetch
func notifyServerForPart2(day int) {
	// Brief delay for AoC to fully unlock Part 2
	time.Sleep(500 * time.Millisecond)

	url := "http://localhost:3030/fetchPart2"
	payload := map[string]int{"day": day}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("⚠️  Failed to notify server: %v\n", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("⚠️  Failed to notify server: %v\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("⚠️  Auto-fetch server not running - Part 2 won't auto-fetch")
		fmt.Println("   (Start server with ./start_server.command)")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("📡 Notified server - Part 2 will auto-fetch shortly")
	} else {
		fmt.Printf("⚠️  Server responded with status %d\n", resp.StatusCode)
	}
}

// copyPart1ToPart2 copies the part1.go file to part2.go, updating tags and part numbers
func (h *Harness) copyPart1ToPart2() error {
	// Read part1.go
	part1Content, err := os.ReadFile("part1.go")
	if err != nil {
		return fmt.Errorf("failed to read part1.go: %v", err)
	}

	// Convert to string and update tags/references
	content := string(part1Content)
	content = strings.ReplaceAll(content, "//go:build part1", "//go:build part2")

	// Update harness.New() call to use part 2
	// This handles both formats: harness.New(solve, &Answer, 1, opts...)
	// Need to use regexp to replace the number after the third parameter
	re := regexp.MustCompile(`(harness\.New\([^,]+,\s*[^,]+,\s*)1(\s*,)`)
	content = re.ReplaceAllString(content, "${1}2${2}")

	// Write to part2.go
	if err := os.WriteFile("part2.go", []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write part2.go: %v", err)
	}

	fmt.Println("📝 Copied part1.go to part2.go")
	return nil
}
