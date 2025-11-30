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

	// Warmup: run timing functions multiple times to avoid cold start overhead
	for i := 0; i < 10; i++ {
		start := time.Now()
		_ = time.Since(start)
	}

	result1, elapsed1 := measure(func() int { return s.Part1(lines) })
	result2, elapsed2 := measure(func() int { return s.Part2(lines) })

	fmt.Printf("Part 1: %d (took %d ns)\n", result1, elapsed1.Nanoseconds())
	fmt.Printf("Part 2: %d (took %d ns)\n", result2, elapsed2.Nanoseconds())
	fmt.Printf("Total: %d ns\n", (elapsed1 + elapsed2).Nanoseconds())
}

func measure(fn func() int) (int, time.Duration) {
	start := time.Now()
	result := fn()
	return result, time.Since(start)
}
