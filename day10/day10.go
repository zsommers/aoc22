package day10

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/util"
)

func A(input []string) int {
	register := []int{1, 1}
	for _, l := range input {
		register = execute(l, register)
	}

	sum := 0
	for i := 20; i <= 220; i += 40 {
		sum += i * register[i]
	}
	return sum
}

func B(input []string) int {
	register := []int{1}
	for _, line := range input {
		register = execute(line, register)
	}

	for l := 0; l < 6; l++ {
		for i := 0; i < 40; i++ {
			v := register[i+40*l]
			if v-1 <= i && v+1 >= i {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	return 0
}

func execute(instruction string, register []int) []int {
	parts := strings.Split(instruction, " ")
	tail := register[len(register)-1]
	switch parts[0] {
	case "noop":
		register = append(register, tail)
	case "addx":
		register = append(register, tail, tail+util.Atoi(parts[1]))
	default:
		panic(fmt.Sprintf("Unknown instruction: '%s'", parts[0]))
	}
	return register
}
