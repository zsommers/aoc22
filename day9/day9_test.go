package day9

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsommers/aoc22/set"
)

var rawInput = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

var rawInput2 = `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`

var input = strings.Split(rawInput, "\n")
var input2 = strings.Split(rawInput2, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 13, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 1, B(input))
	assert.Equal(t, 36, B(input2))
}

func Test_parseInput(t *testing.T) {
	expected := []intPair{
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{0, 1},
		{0, 1},
		{0, 1},
		{0, 1},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{0, -1},
		{1, 0},
		{1, 0},
		{1, 0},
		{1, 0},
		{0, -1},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{-1, 0},
		{1, 0},
		{1, 0},
	}
	assert.Equal(t, expected, parseInput(input))
}

func Test_intPair_distance(t *testing.T) {
	tests := []struct {
		b    intPair
		want int
	}{
		{intPair{0, 0}, 0},
		{intPair{1, 0}, 1},
		{intPair{0, 1}, 1},
		{intPair{-1, 0}, 1},
		{intPair{0, -1}, 1},
		{intPair{1, 1}, 2},
		{intPair{1, -1}, 2},
		{intPair{-1, -1}, 2},
		{intPair{-1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.b), func(t *testing.T) {
			assert.Equal(t, tt.want, (&intPair{}).distance(&tt.b))
		})
	}
}

func Test_move(t *testing.T) {
	tests := []struct {
		head *intPair
		want *intPair
	}{
		{&intPair{-2, -1}, &intPair{-1, -1}},
		{&intPair{-2, 0}, &intPair{-1, 0}},
		{&intPair{-2, 1}, &intPair{-1, 1}},
		{&intPair{-1, -2}, &intPair{-1, -1}},
		{&intPair{-1, -1}, nil},
		{&intPair{-1, 0}, nil},
		{&intPair{-1, 1}, nil},
		{&intPair{-1, 2}, &intPair{-1, 1}},
		{&intPair{0, -2}, &intPair{0, -1}},
		{&intPair{0, -1}, nil},
		{&intPair{0, 0}, nil},
		{&intPair{0, 1}, nil},
		{&intPair{0, 2}, &intPair{0, 1}},
		{&intPair{1, -2}, &intPair{1, -1}},
		{&intPair{1, -1}, nil},
		{&intPair{1, 0}, nil},
		{&intPair{1, 1}, nil},
		{&intPair{1, 2}, &intPair{1, 1}},
		{&intPair{2, -1}, &intPair{1, -1}},
		{&intPair{2, 0}, &intPair{1, 0}},
		{&intPair{2, 1}, &intPair{1, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.head), func(t *testing.T) {
			tail := &intPair{}
			trail := set.New[intPair]()
			moveTail(tt.head, tail, trail)
			if tt.want != nil {
				assert.True(t, tt.want.equal(tail))
				assert.Equal(t, 1, trail.Len())
				assert.True(t, trail.Has(*tt.want))
			} else {
				assert.Equal(t, intPair{}, *tail)
				assert.Zero(t, trail.Len())
			}
		})
	}
}
