package main

import "advent-of-code-2025/utils"

func Part1(lines []string) int {
	count := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[row]); col++ {
			if lines[row][col] != '@' {
				continue
			}
			adj := utils.CountNeighbors8(lines, col, row, '@')
			if adj < 4 {
				count++
			}
		}
	}

	return count
}

func Part2(lines []string) int {
	// Convert strings to rune slices for mutability
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	count := 0
	canContinueRemoving := true
	for canContinueRemoving {
		canContinueRemoving = false
		for row := 0; row < len(grid); row++ {
			for col := 0; col < len(grid[row]); col++ {
				if grid[row][col] != '@' {
					continue
				}
				adj := utils.CountNeighbors8(grid, col, row, '@')
				if adj < 4 {
					grid[row][col] = '.'
					count++
					canContinueRemoving = true
				}
			}
		}
	}
	return count
}

func main() {
	utils.Solution{
		Day:   4,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
