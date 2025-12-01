package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	baseDir := "/Users/alexandroslagios/src/Personal/advent-of-code-2025"

	mainTemplate := `package main

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
	filename := "day%s/input.txt"
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
`

	testTemplate := `package main

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
		t.Errorf("Part1() = %%d, want %%d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		// Add test input here
	}
	got := Part2(input)
	want := 0 // Update with expected value
	if got != want {
		t.Errorf("Part2() = %%d, want %%d", got, want)
	}
}
`

	readmeTemplate := `# Day %d

## Part 1

## Part 2
`

	for i := 6; i <= 25; i++ {
		dayStr := fmt.Sprintf("%02d", i)
		dayDir := filepath.Join(baseDir, "day"+dayStr)

		// Create main.go
		mainContent := fmt.Sprintf(mainTemplate, dayStr)
		if err := os.WriteFile(filepath.Join(dayDir, "main.go"), []byte(mainContent), 0644); err != nil {
			fmt.Printf("Error creating main.go for day%s: %v\n", dayStr, err)
			continue
		}

		// Create main_test.go
		if err := os.WriteFile(filepath.Join(dayDir, "main_test.go"), []byte(testTemplate), 0644); err != nil {
			fmt.Printf("Error creating main_test.go for day%s: %v\n", dayStr, err)
			continue
		}

		// Create README.md
		readmeContent := fmt.Sprintf(readmeTemplate, i)
		if err := os.WriteFile(filepath.Join(dayDir, "README.md"), []byte(readmeContent), 0644); err != nil {
			fmt.Printf("Error creating README.md for day%s: %v\n", dayStr, err)
			continue
		}

		// Create empty input.txt
		if err := os.WriteFile(filepath.Join(dayDir, "input.txt"), []byte(""), 0644); err != nil {
			fmt.Printf("Error creating input.txt for day%s: %v\n", dayStr, err)
			continue
		}

		fmt.Printf("Created files for day%s\n", dayStr)
	}

	fmt.Println("All done!")
}
