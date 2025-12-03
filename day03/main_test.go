package main

import (
	"testing"

	"advent-of-code-2025/utils"
)

func TestPart1(t *testing.T) {
	expected := 357
	input := utils.MustReadLines("input_test.txt")
	got := Part1(input)
	if got != expected {
		t.Errorf("Part1() = %d, want %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 0
	input := utils.MustReadLines("input_test.txt")
	got := Part2(input)
	if got != expected {
		t.Errorf("Part2() = %d, want %d", got, expected)
	}
}
