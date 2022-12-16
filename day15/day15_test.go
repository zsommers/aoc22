package day15

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zsommers/aoc22/point"
)

var rawInput = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`

var input = strings.Split(rawInput, "\n")

var sensors = []*sensor{
	{
		loc:    &point.Point{X: 2, Y: 18},
		beacon: &point.Point{X: -2, Y: 15},
		dist:   7,
	},
	{
		loc:    &point.Point{X: 9, Y: 16},
		beacon: &point.Point{X: 10, Y: 16},
		dist:   1,
	},
	{
		loc:    &point.Point{X: 13, Y: 2},
		beacon: &point.Point{X: 15, Y: 3},
		dist:   3,
	},
	{
		loc:    &point.Point{X: 12, Y: 14},
		beacon: &point.Point{X: 10, Y: 16},
		dist:   4,
	},
	{
		loc:    &point.Point{X: 10, Y: 20},
		beacon: &point.Point{X: 10, Y: 16},
		dist:   4,
	},
	{
		loc:    &point.Point{X: 14, Y: 17},
		beacon: &point.Point{X: 10, Y: 16},
		dist:   5,
	},
	{
		loc:    &point.Point{X: 8, Y: 7},
		beacon: &point.Point{X: 2, Y: 10},
		dist:   9,
	},
	{
		loc:    &point.Point{X: 2, Y: 0},
		beacon: &point.Point{X: 2, Y: 10},
		dist:   10,
	},
	{
		loc:    &point.Point{X: 0, Y: 11},
		beacon: &point.Point{X: 2, Y: 10},
		dist:   3,
	},
	{
		loc:    &point.Point{X: 20, Y: 14},
		beacon: &point.Point{X: 25, Y: 17},
		dist:   8,
	},
	{
		loc:    &point.Point{X: 17, Y: 20},
		beacon: &point.Point{X: 21, Y: 22},
		dist:   6,
	},
	{
		loc:    &point.Point{X: 16, Y: 7},
		beacon: &point.Point{X: 15, Y: 3},
		dist:   5,
	},
	{
		loc:    &point.Point{X: 14, Y: 3},
		beacon: &point.Point{X: 15, Y: 3},
		dist:   1,
	},
	{
		loc:    &point.Point{X: 20, Y: 1},
		beacon: &point.Point{X: 15, Y: 3},
		dist:   7,
	},
}

func TestA(t *testing.T) {
	assert.Equal(t, 0, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 0, B(input))
}

func Test_parseInput(t *testing.T) {
	for i, s := range parseInput(input) {
		s2 := sensors[i]
		t.Run(s2.loc.ToString(), func(t *testing.T) {
			assert.True(t, s.loc.Equal(s2.loc))
			assert.True(t, s.beacon.Equal(s2.beacon))
			assert.Equal(t, s2.dist, s.dist)
		})
	}
}

func Test_countCovered(t *testing.T) {
	assert.Equal(t, 26, countCovered(10, sensors).Len())
}

func Test_countCoveredLimit(t *testing.T) {
	assert.Equal(t, 20, countCoveredLimit(10, 20, sensors).Len())
	assert.Equal(t, 19, countCoveredLimit(11, 20, sensors).Len())
}

func Test_findDistressBeacon(t *testing.T) {
	expected := point.New(14, 11)
	got := findDistressBeacon(20, sensors)
	fmt.Println(got)
	assert.True(t, expected.Equal(got))
}
