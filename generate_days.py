#!/usr/bin/env python3
import os

base_dir = "/Users/alexandroslagios/src/Personal/advent-of-code-2025"

main_template = """package main

import (
\t"fmt"
\t"log"
\t"os"

\t"advent-of-code-2025/utils"
)

func Part1(lines []string) int {{
\t// TODO: Implement part 1
\treturn 0
}}

func Part2(lines []string) int {{
\t// TODO: Implement part 2
\treturn 0
}}

func main() {{
\tfilename := "day{day}/input.txt"
\tif len(os.Args) > 1 {{
\t\tfilename = os.Args[1]
\t}}

\tlines, err := utils.ReadLines(filename)
\tif err != nil {{
\t\tlog.Fatal(err)
\t}}

\tfmt.Println("Part 1:", Part1(lines))
\tfmt.Println("Part 2:", Part2(lines))
}}
"""

test_template = """package main

import (
\t"testing"
)

func TestPart1(t *testing.T) {{
\tinput := []string{{
\t\t// Add test input here
\t}}
\tgot := Part1(input)
\twant := 0 // Update with expected value
\tif got != want {{
\t\tt.Errorf("Part1() = %d, want %d", got, want)
\t}}
}}

func TestPart2(t *testing.T) {{
\tinput := []string{{
\t\t// Add test input here
\t}}
\tgot := Part2(input)
\twant := 0 // Update with expected value
\tif got != want {{
\t\tt.Errorf("Part2() = %d, want %d", got, want)
\t}}
}}
"""

readme_template = """# Day {day_num}

## Part 1

## Part 2
"""

for i in range(1, 26):
    day = f"{i:02d}"
    day_dir = os.path.join(base_dir, f"day{day}")

    # Create main.go
    with open(os.path.join(day_dir, "main.go"), "w") as f:
        f.write(main_template.format(day=day))

    # Create main_test.go
    with open(os.path.join(day_dir, "main_test.go"), "w") as f:
        f.write(test_template)

    # Create README.md
    with open(os.path.join(day_dir, "README.md"), "w") as f:
        f.write(readme_template.format(day_num=i))

    # Create empty input.txt
    open(os.path.join(day_dir, "input.txt"), "w").close()

    print(f"Created files for day{day}")

print("All done!")

