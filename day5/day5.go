package day5

import (
	"bytes"
	"strings"

	"github.com/zsommers/aoc22/stack"
	"github.com/zsommers/aoc22/util"
)

func buildStacks(stackCount int, input []string) []*stack.Stack[byte] {
	stacks := make([]*stack.Stack[byte], 0, stackCount)
	for i := 0; i < stackCount; i++ {
		stacks = append(stacks, stack.New[byte]())
	}

	for _, line := range input {
		for i, s := range stacks {
			if c := line[4*i+1]; c != ' ' {
				s.Push(c)
			}
		}
	}

	return stacks
}

func initState(input []string) ([]*stack.Stack[byte], []string) {
	startingCrates := []string{}
	var stacks []*stack.Stack[byte]
	var moves []string
	for i, line := range input {
		if line[:3] == " 1 " {
			stacks = buildStacks(util.Atoi(line[len(line)-2:len(line)-1]), startingCrates)
			moves = input[i+2:]
			break
		}
		// Invert starting state to insert from bottom in stack
		startingCrates = append([]string{line}, startingCrates...)
	}
	return stacks, moves
}

func moveCrates(crates []*stack.Stack[byte], moves []string) {
	for _, m := range moves {
		parts := strings.Split(m, " ")
		count := util.Atoi(parts[1])
		source := util.Atoi(parts[3]) - 1
		destination := util.Atoi(parts[5]) - 1

		for i := 0; i < count; i++ {
			crate, ok := crates[source].Pop()
			if !ok {
				panic("Bad crate!")
			}
			crates[destination].Push(crate)
		}
	}
}

func multiMoveCrates(crates []*stack.Stack[byte], moves []string) {
	for _, m := range moves {
		parts := strings.Split(m, " ")
		count := util.Atoi(parts[1])
		source := util.Atoi(parts[3]) - 1
		destination := util.Atoi(parts[5]) - 1

		moveStack := stack.New[byte]()
		for i := 0; i < count; i++ {
			crate, ok := crates[source].Pop()
			if !ok {
				panic("Bad crate!")
			}
			moveStack.Push(crate)
		}
		for i := 0; i < count; i++ {
			crate, ok := moveStack.Pop()
			if !ok {
				panic("Bad crate!")
			}
			crates[destination].Push(crate)
		}
	}
}

func A(input []string) string {
	crates, moves := initState(input)
	moveCrates(crates, moves)
	result := bytes.NewBufferString("")
	for _, stack := range crates {
		c, _ := stack.Pop()
		result.WriteByte(c)
	}
	return result.String()
}

func B(input []string) string {
	crates, moves := initState(input)
	multiMoveCrates(crates, moves)
	result := bytes.NewBufferString("")
	for _, stack := range crates {
		c, _ := stack.Pop()
		result.WriteByte(c)
	}
	return result.String()
}
