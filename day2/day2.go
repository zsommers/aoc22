package day2

import (
	"fmt"
)

var shapePoints = map[byte]int{
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var outcomePoints = map[byte]int{
	'X': 0,
	'Y': 3,
	'Z': 6,
}

func outcomeScore(aByte, bByte byte) int {
	a := int(aByte - 'A')
	b := int(bByte - 'X')
	switch b {
	case a:
		return 3
	case (a + 1) % 3:
		return 6
	case (a + 2) % 3:
		return 0
	default:
		panic(fmt.Sprintf("%c, %c is not valid", aByte, bByte))
	}
}

func calculateMoveScore(aByte, b byte) int {
	a := int(aByte - 'A')
	var move int
	switch b {
	case 'X': // Lose
		move = (a + 2) % 3
	case 'Y':
		move = a
	case 'Z':
		move = (a + 1) % 3
	default:
		panic(fmt.Sprintf("%c, %c is not valid", aByte, b))
	}
	return shapePoints[byte(move+'X')]
}

func A(input []string) int {
	sum := 0
	for _, line := range input {
		sum += (int(line[2]) - 87) + outcomeScore(line[0], line[2])
	}
	return sum
}

func B(input []string) int {
	sum := 0
	for _, line := range input {
		sum += calculateMoveScore(line[0], line[2]) + outcomePoints[line[2]]
	}
	return sum
}
