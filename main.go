package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zsommers/aoc22/day1"
	"github.com/zsommers/aoc22/day2"
	"github.com/zsommers/aoc22/day3"
	"github.com/zsommers/aoc22/day4"
	"github.com/zsommers/aoc22/day5"
	"github.com/zsommers/aoc22/day6"
	"github.com/zsommers/aoc22/day7"
	"github.com/zsommers/aoc22/day8"
	"github.com/zsommers/aoc22/day9"
	"github.com/zsommers/aoc22/util"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Expected argument specifying exercise (eg. '1a')")
		os.Exit(1)
	}

	input := util.ReadLines(fmt.Sprintf("day%s/input.txt", flag.Arg(0)[:len(flag.Arg(0))-1]))

	var result interface{}
	switch flag.Arg(0) {
	case "1a":
		result = day1.A(input)
	case "1b":
		result = day1.B(input)
	case "2a":
		result = day2.A(input)
	case "2b":
		result = day2.B(input)
	case "3a":
		result = day3.A(input)
	case "3b":
		result = day3.B(input)
	case "4a":
		result = day4.A(input)
	case "4b":
		result = day4.B(input)
	case "5a":
		result = day5.A(input)
	case "5b":
		result = day5.B(input)
	case "6a":
		result = day6.A(input)
	case "6b":
		result = day6.B(input)
	case "7a":
		result = day7.A(input)
	case "7b":
		result = day7.B(input)
	case "8a":
		result = day8.A(input)
	case "8b":
		result = day8.B(input)
	case "9a":
		result = day9.A(input)
	case "9b":
		result = day9.B(input)
	default:
		fmt.Printf("'%s' is not a valid day", flag.Arg(0))
		os.Exit(1)
	}

	fmt.Printf("Result for '%s': '%v'", flag.Arg(0), result)
}
