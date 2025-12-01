# Advent of Code 2025

My solutions for [Advent of Code 2025](https://adventofcode.com/2025) in Go.

## Quick Start

```bash
# Run a solution
go run day01/main.go

# Run tests
go test ./day01

# Create new day (copy from day01)
cp -r day01 day06 && sed -i '' 's/Day:   1/Day:   6/' day06/main.go
```

## Structure

```
advent-of-code-2025/
├── day01/
│   ├── main.go          # Solution implementation
│   ├── main_test.go     # Tests for part 1
│   ├── input.txt        # Actual puzzle input
│   ├── input_test.txt   # Example input from problem
│   └── README.md        # Notes (optional)
├── day02/ ... day12/
│   └── ... (same structure)
└── utils/               # Shared utility functions
    ├── runner.go        # Solution runner with timing
    ├── input.go         # File reading helpers
    ├── mathutils.go     # Math operations
    └── ... (other utilities)
```

- `utils/` - Shared utility functions for:
  - **Solution runner** (handles file I/O, timing, and output)
  - File I/O (reading lines, grids, groups, etc.)
  - Math operations (GCD, LCM, distances, etc.)
  - Grid utilities (Point, Direction, neighbors, transformations)
  - Collections (Set, Queue, Stack, generic helpers)
  - Parsing helpers (extract integers, split strings, etc.)
  - Search algorithms (BFS, DFS, Dijkstra, A*)

## Running Solutions

```bash
# Run a specific day (shows timing in nanoseconds)
go run day01/main.go

# Run with custom input file
go run day01/main.go path/to/input.txt
```

## Generating a New Day

To create a new day (e.g., day06), you need to create 4 files:

### Option 1: Copy from existing day
```bash
# Copy day01 as a template
cp -r day01 day06

# Update the day number in main.go
sed -i '' 's/Day:   1/Day:   6/' day06/main.go

# Update test file
sed -i '' 's/day01/day06/g' day06/main_test.go

# Clear the input files
> day06/input.txt
> day06/input_test.txt
```

### Option 2: Create manually

**1. Create directory:**
```bash
mkdir day06
```

**2. Create `day06/main.go`:**
```go
package main

import "advent-of-code-2025/utils"

func Part1(lines []string) int {
    // TODO: Implement part 1
    return 0
}

func Part2(lines []string) int {
    // TODO: Implement part 2
    return 0
}

func main() {
    utils.Solution{
        Day:   6,
        Part1: Part1,
        Part2: Part2,
    }.Run()
}
```

**3. Create `day06/main_test.go`:**
```go
package main

import (
    "testing"

    "advent-of-code-2025/utils"
)

const expected = 0

func TestPart1(t *testing.T) {
    input := utils.MustReadLines("day06/input_test.txt")
    got := Part1(input)
    if got != expected {
        t.Errorf("Part1() = %d, want %d", got, expected)
    }
}
```

**4. Create empty input files:**
```bash
touch day06/input.txt day06/input_test.txt
```

**5. (Optional) Create `day06/README.md`:**
```markdown
# Day 6

## Part 1

## Part 2
```

## Testing

Each day has:
- `input_test.txt` - The example input from the problem description
- `main_test.go` - Test file with expected value at the top

Test pattern:
```go
const expected = 42 // Update with expected value for test input

func TestPart1(t *testing.T) {
    input := utils.MustReadLines("input_test.txt")
    got := Part1(input)
    if got != expected {
        t.Errorf("Part1() = %d, want %d", got, expected)
    }
}
```

Run tests:
```bash
# Test a specific day
go test ./day01

# Test all days
go test ./...

# Verbose output
go test -v ./day01
```

## Performance

Each solution automatically:
- Measures execution time in nanoseconds for both parts
- Includes warmup to avoid cold start overhead
- Reports individual and total execution times

