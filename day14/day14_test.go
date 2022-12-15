package day14

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsommers/aoc22/point"
)

var rawInput = `498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`

var input = strings.Split(rawInput, "\n")

func baseGrid() grid {
	return grid{
		{X: 498, Y: 4}: rock,
		{X: 498, Y: 5}: rock,
		{X: 498, Y: 6}: rock,
		{X: 497, Y: 6}: rock,
		{X: 496, Y: 6}: rock,
		{X: 503, Y: 4}: rock,
		{X: 502, Y: 4}: rock,
		{X: 502, Y: 5}: rock,
		{X: 502, Y: 6}: rock,
		{X: 502, Y: 7}: rock,
		{X: 502, Y: 8}: rock,
		{X: 502, Y: 9}: rock,
		{X: 501, Y: 9}: rock,
		{X: 500, Y: 9}: rock,
		{X: 499, Y: 9}: rock,
		{X: 498, Y: 9}: rock,
		{X: 497, Y: 9}: rock,
		{X: 496, Y: 9}: rock,
		{X: 495, Y: 9}: rock,
		{X: 494, Y: 9}: rock,
	}
}

func TestA(t *testing.T) {
	assert.Equal(t, 24, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 93, B(input))
}

func Test_parseInput(t *testing.T) {
	assert.Equal(t, baseGrid(), parseInput(input))
}

func Test_dropSand(t *testing.T) {
	test := baseGrid()
	expected := baseGrid()

	expected[point.Point{X: 500, Y: 8}] = sand
	assert.True(t, dropSand(test))
	assert.Equal(t, expected, test)

	expected[point.Point{X: 499, Y: 8}] = sand
	assert.True(t, dropSand(test))
	assert.Equal(t, expected, test)

	expected[point.Point{X: 501, Y: 8}] = sand
	expected[point.Point{X: 498, Y: 8}] = sand
	expected[point.Point{X: 500, Y: 7}] = sand
	for i := 0; i < 3; i++ {
		assert.True(t, dropSand(test))
	}
	assert.Equal(t, expected, test)

	expected[point.Point{X: 497, Y: 8}] = sand
	expected[point.Point{X: 498, Y: 7}] = sand
	expected[point.Point{X: 499, Y: 7}] = sand
	expected[point.Point{X: 501, Y: 7}] = sand
	expected[point.Point{X: 499, Y: 6}] = sand
	expected[point.Point{X: 500, Y: 6}] = sand
	expected[point.Point{X: 501, Y: 6}] = sand
	expected[point.Point{X: 499, Y: 5}] = sand
	expected[point.Point{X: 500, Y: 5}] = sand
	expected[point.Point{X: 501, Y: 5}] = sand
	expected[point.Point{X: 499, Y: 4}] = sand
	expected[point.Point{X: 500, Y: 4}] = sand
	expected[point.Point{X: 501, Y: 4}] = sand
	expected[point.Point{X: 499, Y: 3}] = sand
	expected[point.Point{X: 500, Y: 3}] = sand
	expected[point.Point{X: 501, Y: 3}] = sand
	expected[point.Point{X: 500, Y: 2}] = sand
	for i := 0; i < 17; i++ {
		assert.True(t, dropSand(test))
	}
	assert.Equal(t, expected, test)

	expected[point.Point{X: 497, Y: 5}] = sand
	expected[point.Point{X: 495, Y: 8}] = sand
	for i := 0; i < 2; i++ {
		assert.True(t, dropSand(test))
	}
	assert.Equal(t, expected, test)

	expected[point.Point{X: 493, Y: 10}] = sand
	assert.False(t, dropSand(test))
	assert.Equal(t, expected, test)
}
