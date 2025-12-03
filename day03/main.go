package main

import (
	"advent-of-code-2025/utils"
	"strconv"
	"strings"
)

func Part1(lines []string) int {
	resultChan := make(chan int, len(lines))

	for _, lineStr := range lines {
		go func(r string) {
			digits := strings.Split(r, "")

			maxCharge := findMaxCharge(digits, 2)

			resultChan <- maxCharge
			//log.Println(digits, maxCharge)
		}(lineStr)
	}
	total := 0
	for i := 0; i < len(lines); i++ {
		total += <-resultChan
	}
	return total
}

func findMaxCharge(digits []string, numberOfDigits int) int {
	currentCharge := 0
	currentIndex := 0
	remaining := numberOfDigits

	for remaining > 0 && currentIndex < len(digits) {
		nextMaxDigit, reachedIndex := findNextMaxDigit(digits, currentIndex, remaining)
		currentCharge = currentCharge*10 + nextMaxDigit
		currentIndex = reachedIndex + 1
		remaining--
	}

	return currentCharge
}

func findNextMaxDigit(digits []string, index int, remainingDigits int) (int, int) {
	maxDigit := 0
	reachedIndex := index
	for i := index; i < len(digits); i++ {
		if remainingDigits > len(digits)-i {
			break
		}
		digit, _ := strconv.Atoi(digits[i])
		if digit > maxDigit {
			maxDigit = digit
			reachedIndex = i
		}
	}
	return maxDigit, reachedIndex
}

func Part2(lines []string) int {
	resultChan := make(chan int, len(lines))

	for _, lineStr := range lines {
		go func(r string) {
			digits := strings.Split(r, "")

			maxCharge := findMaxCharge(digits, 12)

			resultChan <- maxCharge
			//log.Println(digits, maxCharge)
		}(lineStr)
	}
	total := 0
	for i := 0; i < len(lines); i++ {
		total += <-resultChan
	}
	return total
}

func main() {
	utils.Solution{
		Day:   3,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
