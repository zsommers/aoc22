package day11

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/util"
)

type monkey struct {
	items       []int
	operation   func(int) int
	testDivisor int
	trueIndex   int
	falseIndex  int
	inspections int
}

func A(input []string) int {
	monkeys := parseInput(input)

	for round := 0; round < 20; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				item = m.operation(item)
				item /= 3
				if item%m.testDivisor == 0 {
					monkeys[m.trueIndex].items = append(monkeys[m.trueIndex].items, item)
				} else {
					monkeys[m.falseIndex].items = append(monkeys[m.falseIndex].items, item)
				}
				m.inspections++
			}
			m.items = []int{}
		}
	}

	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	highest := util.HighestInts(2, inspections...)
	return highest[0] * highest[1]
}

func B(input []string) int {
	monkeys := parseInput(input)

	lcm := 1
	for _, m := range monkeys {
		lcm *= m.testDivisor
	}

	for round := 0; round < 10000; round++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				item = m.operation(item)
				item %= lcm
				if item%m.testDivisor == 0 {
					monkeys[m.trueIndex].items = append(monkeys[m.trueIndex].items, item)
				} else {
					monkeys[m.falseIndex].items = append(monkeys[m.falseIndex].items, item)
				}
				m.inspections++
			}
			m.items = []int{}
		}
	}

	inspections := []int{}
	for _, m := range monkeys {
		inspections = append(inspections, m.inspections)
	}
	highest := util.HighestInts(2, inspections...)
	return highest[0] * highest[1]
}

func parseItems(s string) []int {
	items := []int{}
	for _, p := range strings.Split(s, " ")[2:] {
		items = append(items, util.Atoi(strings.TrimRight(p, ",")))
	}
	return items
}

func parseOperation(s string) func(int) int {
	parts := strings.Split(s, " ")
	return func(i int) int {
		b := i
		if parts[5] != "old" {
			b = util.Atoi(parts[5])
		}
		switch parts[4] {
		case "+":
			return i + b
		case "*":
			return i * b
		default:
			panic(fmt.Sprintf("Unrecognized operand: '%s'", parts[4]))
		}
	}
}

func parseInput(input []string) []*monkey {
	monkeys := []*monkey{}

	i := 0
	for {
		if input[i][:6] != "Monkey" {
			panic(fmt.Sprintf("Unexpected input: %s", input[i]))
		}

		m := monkey{}
		m.items = parseItems(strings.TrimSpace(input[i+1]))
		m.operation = parseOperation(strings.TrimSpace(input[i+2]))
		m.testDivisor = util.Atoi(strings.Split(strings.TrimSpace(input[i+3]), " ")[3])
		m.trueIndex = util.Atoi(strings.Split(strings.TrimSpace(input[i+4]), " ")[5])
		m.falseIndex = util.Atoi(strings.Split(strings.TrimSpace(input[i+5]), " ")[5])

		monkeys = append(monkeys, &m)

		i += 7
		if len(input) < i+6 {
			break
		}
	}

	return monkeys
}
