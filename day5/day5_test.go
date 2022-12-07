package day5

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zsommers/aoc22/stack"
)

var rawInput = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, "CMZ", A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, "MCD", B(input))
}

func expectedStacks() []*stack.Stack[byte] {
	stackA := stack.New[byte]()
	stackA.Push('Z')
	stackA.Push('N')
	stackB := stack.New[byte]()
	stackB.Push('M')
	stackB.Push('C')
	stackB.Push('D')
	stackC := stack.New[byte]()
	stackC.Push('P')
	return []*stack.Stack[byte]{
		stackA, stackB, stackC,
	}
}

func Test_initState(t *testing.T) {
	expectedMoves := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	var stacks []*stack.Stack[byte]
	var moves []string
	require.NotPanics(t, func() { stacks, moves = initState(input) })
	assert.Equal(t, expectedStacks(), stacks)
	assert.Equal(t, expectedMoves, moves)
}

func Test_buildStacks(t *testing.T) {
	var stacks []*stack.Stack[byte]
	state := []string{
		input[2],
		input[1],
		input[0],
	}
	require.NotPanics(t, func() { stacks = buildStacks(3, state) })
	for i, eStack := range expectedStacks() {
		assert.True(t, stack.Equal(eStack, stacks[i]))
	}
}
