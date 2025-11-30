package utils

import (
	"regexp"
	"strconv"
	"strings"
)

// ParseInts extracts all integers from a string
func ParseInts(s string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(s, -1)
	nums := make([]int, 0, len(matches))
	for _, match := range matches {
		n, err := strconv.Atoi(match)
		if err == nil {
			nums = append(nums, n)
		}
	}
	return nums
}

// SplitOnComma splits a string on commas and trims whitespace
func SplitOnComma(s string) []string {
	parts := strings.Split(s, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

// SplitOnSpace splits a string on whitespace
func SplitOnSpace(s string) []string {
	return strings.Fields(s)
}

// ToInt converts a string to an integer, panics on error
func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

// ToIntSafe converts a string to an integer, returns 0 on error
func ToIntSafe(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// ToInt64 converts a string to an int64, panics on error
func ToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// ExtractNumbers extracts all numbers from a string slice
func ExtractNumbers(lines []string) []int {
	var nums []int
	for _, line := range lines {
		nums = append(nums, ParseInts(line)...)
	}
	return nums
}

// SplitByEmptyLine splits lines by empty lines
func SplitByEmptyLine(lines []string) [][]string {
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

	return groups
}
