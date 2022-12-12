package day11

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 10605, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 2713310158, B(input))
}

func Test_parseItems(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{
			"Starting items: 79, 98",
			[]int{79, 98},
		},
		{
			"Starting items: 54, 65, 75, 74",
			[]int{54, 65, 75, 74},
		},
		{
			"Starting items: 74",
			[]int{74},
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, parseItems(tt.input))
		})
	}
}

func Test_parseOperation(t *testing.T) {
	tests := []struct {
		input string
		want  func(int) int
	}{
		{
			"Operation: new = old * 19",
			func(i int) int { return i * 19 },
		},
		{
			"Operation: new = old + 6",
			func(i int) int { return i + 6 },
		},
		{
			"Operation: new = old * old",
			func(i int) int { return i * i },
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			f := parseOperation(tt.input)
			assertFuncsEqual(t, tt.want, f)
		})
	}
}

func assertFuncsEqual(t *testing.T, want, f func(i int) int) bool {
	nums := []int{2, 8, 32, 128, 512}
	pass := true

	for _, a := range nums {
		pass = pass && assert.Equal(t, want(a), f(a))
	}

	return pass
}

func Test_parseInput(t *testing.T) {
	tests := []struct {
		index int
		want  *monkey
	}{
		{
			0,
			&monkey{
				items:       []int{79, 98},
				operation:   func(i int) int { return i * 19 },
				testDivisor: 23,
				trueIndex:   2,
				falseIndex:  3,
			},
		},
		{
			1,
			&monkey{
				items:       []int{54, 65, 75, 74},
				operation:   func(i int) int { return i + 6 },
				testDivisor: 19,
				trueIndex:   2,
				falseIndex:  0,
			},
		},
		{
			2,
			&monkey{
				items:       []int{79, 60, 97},
				operation:   func(i int) int { return i * i },
				testDivisor: 13,
				trueIndex:   1,
				falseIndex:  3,
			},
		},
		{
			3,
			&monkey{
				items:       []int{74},
				operation:   func(i int) int { return i + 3 },
				testDivisor: 17,
				trueIndex:   0,
				falseIndex:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.index), func(t *testing.T) {
			got := parseInput(input[tt.index*7 : tt.index*7+6])[0]
			assert.Equal(t, tt.want.items, got.items)
			assertFuncsEqual(t, tt.want.operation, got.operation)
			assert.Equal(t, tt.want.testDivisor, got.testDivisor)
			assert.Equal(t, tt.want.trueIndex, got.trueIndex)
			assert.Equal(t, tt.want.falseIndex, got.falseIndex)
		})
	}
}
