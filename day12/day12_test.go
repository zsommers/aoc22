package day12

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsommers/aoc22/point"
	"github.com/zsommers/aoc22/util"
)

var rawInput = `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 31, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 29, B(input))
}

func Test_canVisit(t *testing.T) {
	visited := makeVisited(input)
	visited[1][1] = true
	tests := []struct {
		current, destination *point.Point
		want                 bool
	}{
		{point.New(0, 0), point.New(-1, 0), false},
		{point.New(0, 0), point.New(0, -1), false},
		{point.New(7, 7), point.New(7, 8), false},
		{point.New(7, 7), point.New(8, 7), false},
		{point.New(0, 0), point.New(1, 1), false},
		{point.New(0, 0), point.New(0, 2), false},
		{point.New(0, 0), point.New(2, 0), false},
		{point.New(0, 0), point.New(0, 1), true},
		{point.New(1, 0), point.New(1, 1), false},
		{point.New(4, 2), point.New(5, 2), true},
		{point.New(5, 1), point.New(5, 2), false},
		{point.New(5, 3), point.New(5, 2), false},
		{point.New(6, 2), point.New(5, 2), false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("(%d,%d)-(%d,%d)", tt.current.X, tt.current.Y, tt.destination.X, tt.destination.Y), func(t *testing.T) {
			assert.Equal(t, tt.want, canVisit(tt.current, tt.destination, input, visited))
		})
	}
}

func Test_visit(t *testing.T) {
	visited := makeVisited(input)
	distance := makeDistance(input)
	distance[0][0] = 0

	visit(point.New(0, 0), input, visited, distance)
	assert.True(t, visited[0][0])
	assert.Equal(t, 1, distance[1][0])
	assert.Equal(t, 1, distance[0][1])

	visit(point.New(1, 0), input, visited, distance)
	assert.True(t, visited[0][1])
	assert.Equal(t, 0, distance[0][0])
	assert.Equal(t, 2, distance[0][2])
	assert.Equal(t, 2, distance[1][1])

	visit(point.New(0, 1), input, visited, distance)
	assert.True(t, visited[1][0])
	assert.Equal(t, 0, distance[0][0])
	assert.Equal(t, 2, distance[2][0])
	assert.Equal(t, 2, distance[1][1])

	visit(point.New(2, 0), input, visited, distance)
	assert.True(t, visited[0][2])
	assert.Equal(t, 1, distance[0][1])
	assert.Equal(t, util.MaxInt-1, distance[0][3])
	assert.Equal(t, 3, distance[1][2])

	visit(point.New(2, 1), input, visited, distance)
	assert.True(t, visited[1][2])
	assert.Equal(t, 2, distance[0][2])
	assert.Equal(t, 2, distance[1][1])
	assert.Equal(t, util.MaxInt-1, distance[1][3])
	assert.Equal(t, 4, distance[2][2])
	/*
		Sabqponm
		abcryxxl
		accszExk
		acctuvwj
		abdefghi
	*/
}
