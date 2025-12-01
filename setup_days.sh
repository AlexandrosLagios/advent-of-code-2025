#!/bin/bash
cd /Users/alexandroslagios/src/Personal/advent-of-code-2025

for i in $(seq -f "%02g" 3 25); do
  echo "Creating day$i..."

  cat > "day$i/main.go" <<MAINEOF
package main

import (
	"fmt"
	"log"
	"os"

	"advent-of-code-2025/utils"
)

func Part1(lines []string) int {
	// TODO: Implement part 1
	return 0
}

func Part2(lines []string) int {
	// TODO: Implement part 2
	return 0
}

func main() {
	filename := "day$i/input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	lines, err := utils.ReadLines(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}
MAINEOF

  cat > "day$i/main_test.go" <<TESTEOF
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		// Add test input here
	}
	got := Part1(input)
	want := 0 // Update with expected value
	if got != want {
		t.Errorf("Part1() = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		// Add test input here
	}
	got := Part2(input)
	want := 0 // Update with expected value
	if got != want {
		t.Errorf("Part2() = %d, want %d", got, want)
	}
}
TESTEOF

  daynum=$(echo $i | sed 's/^0*//')
  cat > "day$i/README.md" <<READMEEOF
# Day $daynum

## Part 1

## Part 2
READMEEOF

  touch "day$i/input.txt"
done

echo "All done!"

