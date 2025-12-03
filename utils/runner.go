package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Solution represents a day's solution with two parts
type Solution struct {
	Day   int
	Part1 func([]string) int
	Part2 func([]string) int
}

const (
	totalRuns    = 10001
	warmupRuns   = 1
	measuredRuns = totalRuns - warmupRuns
)

// Run executes both parts of a solution with timing
func (s Solution) Run() {
	filename := fmt.Sprintf("day%02d/input.txt", s.Day)
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	lines, err := ReadLines(filename)
	if err != nil {
		log.Fatal(err)
	}

	result1, elapsed1 := Measure(func() int { return s.Part1(lines) })
	result2, elapsed2 := Measure(func() int { return s.Part2(lines) })

	fmt.Printf("Day %02d\n", s.Day)
	fmt.Printf("  Part 1: %d  (avg of %d runs: %s)\n", result1, measuredRuns, formatDuration(elapsed1))
	fmt.Printf("  Part 2: %d  (avg of %d runs: %s)\n", result2, measuredRuns, formatDuration(elapsed2))
	fmt.Printf("  Total avg: %s\n", formatDuration(elapsed1+elapsed2))
}

func Measure(fn func() int) (int, time.Duration) {
	if measuredRuns <= 0 {
		log.Fatal("measured runs must be positive")
	}

	var result int
	var sum time.Duration

	for i := 0; i < totalRuns; i++ {
		start := time.Now()
		current := fn()
		elapsed := time.Since(start)

		if i < warmupRuns {
			// Treat early invocations as warmups to avoid cold start noise,
			// but still capture the result for consistency checks.
			result = current
			continue
		}

		if current != result {
			log.Fatalf("inconsistent results between runs: got %d and %d", current, result)
		}

		sum += elapsed
	}

	return result, sum / time.Duration(measuredRuns)
}

func formatDuration(d time.Duration) string {
	switch {
	case d < time.Microsecond:
		return fmt.Sprintf("%d ns", d.Nanoseconds())
	case d < time.Millisecond:
		return fmt.Sprintf("%.3f us", float64(d)/float64(time.Microsecond))
	case d < time.Second:
		return fmt.Sprintf("%.3f ms", float64(d)/float64(time.Millisecond))
	default:
		return fmt.Sprintf("%.3f s", float64(d)/float64(time.Second))
	}
}
