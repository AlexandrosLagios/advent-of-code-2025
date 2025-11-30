package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file and returns all lines as a slice of strings
func ReadLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// MustReadLines reads lines or panics
func MustReadLines(filename string) []string {
	lines, err := ReadLines(filename)
	if err != nil {
		panic(fmt.Sprintf("failed to read file %s: %v", filename, err))
	}
	return lines
}

// ReadInts reads a file and returns all integers (one per line)
func ReadInts(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var nums []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

// ReadIntGrid reads a file where each line contains space-separated integers
func ReadIntGrid(filename string) ([][]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var grid [][]int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Fields(line)
		var row []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	return grid, nil
}

// ReadCharGrid reads a file as a 2D character grid
func ReadCharGrid(filename string) ([][]rune, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var grid [][]rune
	for _, line := range lines {
		if line == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	return grid, nil
}

// ReadGroups reads a file split by blank lines into groups
func ReadGroups(filename string) ([][]string, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}

	var groups [][]string
	var current []string

	for _, line := range lines {
		if line == "" {
			if len(current) > 0 {
				groups = append(groups, current)
				current = []string{}
			}
		} else {
			current = append(current, line)
		}
	}

	if len(current) > 0 {
		groups = append(groups, current)
	}

	return groups, nil
}

// ReadRaw reads entire file as a single string
func ReadRaw(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
