package day1

import (
	"sort"

	"github.com/zsommers/aoc22/util"
)

func makeElfCals(input []string) []int {
	elfCals := []int{}
	elfSum := 0
	for _, line := range input {
		if line == "" {
			elfCals = append(elfCals, elfSum)
			elfSum = 0
			continue
		}

		elfSum += util.Atoi(line)
	}
	if elfSum != 0 {
		elfCals = append(elfCals, elfSum)
	}

	return elfCals
}

func A(input []string) int {
	return util.Max(makeElfCals(input)...)
}

func B(input []string) int {
	elfCals := makeElfCals(input)
	sort.Ints(elfCals)
	return util.SumInts(elfCals[len(elfCals)-3:]...)
}
