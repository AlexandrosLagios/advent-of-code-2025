package main

import (
	"advent-of-code-2025/utils"
)

func Part1(lines []string) int {
	var currentPosition = 50
	var result = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := GetDirection(line[0])
		distance := utils.ToInt(line[1:])

		currentPosition = Rotate(currentPosition, direction, distance)
		if currentPosition == 0 {
			result++
		}
	}
	return result
}

type Direction int

const (
	Left  Direction = -1
	Right Direction = 1
)

func GetDirection(char byte) Direction {
	switch char {
	case 'L':
		return Left
	case 'R':
		return Right
	default:
		panic("invalid direction: " + string(char))
	}
}

func Rotate(previous int, direction Direction, distance int) int {
	newPosition := previous + int(direction)*distance
	return (newPosition%100 + 100) % 100
}

func Part2(lines []string) int {
	var currentPosition = 50
	var result = 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := GetDirection(line[0])
		distance := utils.ToInt(line[1:])

		previousPosition := currentPosition
		currentPosition = Rotate(currentPosition, direction, distance)

		result += CountPassesFrom0(previousPosition, direction, distance)
	}
	return result
}

func CountPassesFrom0(previousPosition int, direction Direction, distance int) int {
	if distance == 0 {
		return 0
	}

	var firstHit int
	if direction == Right {
		firstHit = (100 - previousPosition) % 100
	} else {
		firstHit = previousPosition % 100
	}
	if firstHit == 0 {
		firstHit = 100
	}

	if firstHit > distance {
		return 0
	}

	return 1 + (distance-firstHit)/100
}

func main() {
	utils.Solution{
		Day:   1,
		Part1: Part1,
		Part2: Part2,
	}.Run()
}
