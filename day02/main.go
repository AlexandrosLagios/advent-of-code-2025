package main

import (
	"advent-of-code-2025/utils"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	line := lines[0]
	ranges := strings.Split(line, `,`)

	resultChan := make(chan int, len(ranges))

	for _, rangeStr := range ranges {
		go func(r string) {
			parts := strings.Split(r, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			sum := findTwiceRepeatingNumbers(start, end)

			resultChan <- sum
		}(rangeStr)
	}

	total := 0
	for i := 0; i < len(ranges); i++ {
		total += <-resultChan
	}

	return total
}

func findTwiceRepeatingNumbers(start, end int) int {
	sum := 0

	startDigits := len(strconv.Itoa(start))
	endDigits := len(strconv.Itoa(end))

	for digitCount := startDigits; digitCount <= endDigits; digitCount++ {
		if digitCount%2 != 0 {
			continue
		}

		halfDigits := digitCount / 2

		minFirstHalf := powerOf10(halfDigits - 1)
		maxFirstHalf := powerOf10(halfDigits) - 1

		for firstHalf := minFirstHalf; firstHalf <= maxFirstHalf; firstHalf++ {
			firstHalfString := strconv.Itoa(firstHalf)
			fullNumberString := firstHalfString + firstHalfString
			fullNumber, _ := strconv.Atoi(fullNumberString)

			if fullNumber >= start && fullNumber <= end {
				sum += fullNumber
			}

			if fullNumber > end {
				break
			}
		}
	}

	return sum
}

func powerOf10(n int) int {
	result := 1
	for i := 0; i < n; i++ {
		result *= 10
	}
	return result
}

func Part2(lines []string) int {
	line := lines[0]
	ranges := strings.Split(line, `,`)

	resultChan := make(chan int, len(ranges))

	for _, rangeStr := range ranges {
		go func(r string) {
			parts := strings.Split(r, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])

			seen := make(map[int]bool)
			sum := findRepeatingNumbers(end, start, seen)

			resultChan <- sum
		}(rangeStr)
	}

	total := 0
	for i := 0; i < len(ranges); i++ {
		total += <-resultChan
	}

	return total
}

func findRepeatingNumbers(end int, start int, seen map[int]bool) int {
	sum := 0
	maxDigits := len(strconv.Itoa(end))

	for repeatCount := 2; repeatCount <= maxDigits; repeatCount++ {
		for patternLength := 1; patternLength*repeatCount <= maxDigits; patternLength++ {
			minPattern := powerOf10(patternLength - 1)
			if patternLength == 1 {
				minPattern = 1
			}
			maxPattern := powerOf10(patternLength) - 1

			for pattern := minPattern; pattern <= maxPattern; pattern++ {
				patternString := strconv.Itoa(pattern)
				fullNumberString := strings.Repeat(patternString, repeatCount)
				fullNumber, _ := strconv.Atoi(fullNumberString)

				if fullNumber > end {
					break
				}
				if fullNumber >= start && !seen[fullNumber] {
					seen[fullNumber] = true
					sum += fullNumber
				}
			}
		}
	}
	return sum
}

func main() {
	utils.Solution{
		Day:   2,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
