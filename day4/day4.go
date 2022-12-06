package day4

import (
	"strings"

	"github.com/zsommers/aoc22/util"
)

type section struct {
	start, end int
}

func newSection(s string) section {
	ss := strings.Split(s, "-")
	return section{
		util.Atoi(ss[0]),
		util.Atoi(ss[1]),
	}
}

func getSections(s string) (section, section) {
	ss := strings.Split(s, ",")
	return newSection(ss[0]), newSection(ss[1])
}

func checkContainedSection(a, b section) bool {
	if a.start <= b.start && a.end >= b.end {
		return true
	}
	if a.start >= b.start && a.end <= b.end {
		return true
	}
	return false
}

func checkIntersectSection(a, b section) bool {
	return !(a.end < b.start || a.start > b.end)
}

func A(input []string) int {
	count := 0
	for _, s := range input {
		if checkContainedSection(getSections(s)) {
			count++
		}
	}
	return count
}

func B(input []string) int {
	count := 0
	for _, s := range input {
		if checkIntersectSection(getSections(s)) {
			count++
		}
	}
	return count
}
