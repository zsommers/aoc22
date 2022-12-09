package day8

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var rawInput = `30373
25512
65332
33549
35390`

var input = strings.Split(rawInput, "\n")

var grid = [][]int{
	{3, 2, 6, 3, 3},
	{0, 5, 5, 3, 5},
	{3, 5, 3, 5, 3},
	{7, 1, 3, 4, 9},
	{3, 2, 2, 9, 0},
}

func TestA(t *testing.T) {
	assert.Equal(t, 21, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 8, B(input))
}

func Test_parseInput(t *testing.T) {
	assert.Equal(t, grid, parseInput(input))
}

func Test_isVisable(t *testing.T) {
	tests := []struct {
		x, y            int
		want, willPanic bool
	}{
		{-1, 2, false, true},
		{5, 2, false, true},
		{2, -1, false, true},
		{2, 5, false, true},
		{0, 2, true, false},
		{4, 2, true, false},
		{2, 0, true, false},
		{2, 4, true, false},
		{1, 1, true, false},
		{1, 2, true, false},
		{1, 3, false, false},
		{2, 1, true, false},
		{2, 2, false, false},
		{2, 3, true, false},
		{3, 1, false, false},
		{3, 2, true, false},
		{3, 3, false, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d", tt.x, tt.y), func(t *testing.T) {
			if tt.willPanic {
				require.Panics(t, func() { isVisable(tt.x, tt.y, grid) })
			} else {
				var result bool
				require.NotPanics(t, func() { result = isVisable(tt.x, tt.y, grid) })
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

func Test_scenicScore(t *testing.T) {
	tests := []struct {
		x, y int
		want int
	}{
		{2, 1, 4},
		{2, 3, 8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d-%d", tt.x, tt.y), func(t *testing.T) {
			assert.Equal(t, tt.want, scenicScore(tt.x, tt.y, grid))
		})
	}
}
