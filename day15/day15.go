package day15

import (
	"fmt"
	"strings"
	"time"

	"github.com/zsommers/aoc22/point"
	"github.com/zsommers/aoc22/set"
	"github.com/zsommers/aoc22/util"
)

type sensor struct {
	loc    *point.Point
	beacon *point.Point
	dist   int
}

func newSensor(l, b *point.Point) *sensor {
	return &sensor{
		loc:    l,
		beacon: b,
		dist:   l.Distance(b),
	}
}

func parseInput(input []string) []*sensor {
	sensors := []*sensor{}

	for _, l := range input {
		parts := strings.Split(l, " ")
		sx := util.Atoi(parts[2][2 : len(parts[2])-1])
		sy := util.Atoi(parts[3][2 : len(parts[3])-1])
		bx := util.Atoi(parts[8][2 : len(parts[8])-1])
		by := util.Atoi(parts[9][2:len(parts[9])])
		sensors = append(sensors, newSensor(point.New(sx, sy), point.New(bx, by)))
	}

	return sensors
}

func countCovered(row int, sensors []*sensor) *set.Set[int] {
	covered := set.New[int]()
	beacons := []int{}
	for _, s := range sensors {
		for i := 0; util.Abs(s.loc.Y-row)+i <= s.dist; i++ {
			covered.Insert(s.loc.X + i)
			covered.Insert(s.loc.X - i)
		}
		if s.beacon.Y == row {
			beacons = append(beacons, s.beacon.X)
		}
	}
	for _, b := range beacons {
		covered.Remove(b)
	}
	return covered
}

func countCoveredLimit(row, max int, sensors []*sensor) *set.Set[int] {
	covered := set.New[int]()
	for _, s := range sensors {
		for i := 0; util.Abs(s.loc.Y-row)+i <= s.dist; i++ {
			if x := s.loc.X + i; x >= 0 && x < max {
				covered.Insert(x)
			}
			if x := s.loc.X - i; x >= 0 && x < max {
				covered.Insert(x)
			}
		}
	}
	return covered
}

func pointCovered(p *point.Point, sensors []*sensor) bool {
	for _, s := range sensors {
		if s.loc.Distance(p) <= s.dist {
			return true
		}
	}
	return false
}

func findDistressBeacon(max int, sensors []*sensor) *point.Point {
	start := time.Now()
	checked := set.New[point.Point]()
	for sIdx, s := range sensors {
		elapsed := time.Since(start)
		remaining := time.Duration(int(elapsed) / (sIdx + 1) * (len(sensors) - sIdx))
		fmt.Printf("Checking sensor %2d of %2d - %6s elapsed - %6s remaining\n",
			sIdx+1, len(sensors), elapsed.Round(time.Second).String(), remaining.Round(time.Second).String(),
		)
		for i := -1 * (s.dist + 1); i <= s.dist+1; i++ {
			x := s.loc.X + i
			if x < 0 || x > max {
				continue
			}
			delta := s.dist + 1 - util.Abs(i)
			points := []*point.Point{
				point.New(x, s.loc.Y+delta),
				point.New(x, s.loc.Y-delta),
			}
			for _, p := range points {
				if p.Y < 0 || p.Y > max {
					continue
				}
				if checked.Has(*p) {
					continue
				}
				if !pointCovered(p, sensors) {
					return p
				}
				checked.Insert(*p)
			}
		}
	}
	panic("Shouldn't reach here!")
}

func A(input []string) int {
	sensors := parseInput(input)
	return countCovered(2_000_000, sensors).Len()
}

func B(input []string) int {
	sensors := parseInput(input)
	beacon := findDistressBeacon(4_000_000, sensors)
	return beacon.X*4_000_000 + beacon.Y
}
