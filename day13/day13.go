package day13

import (
	"fmt"
	"strings"

	"github.com/zsommers/aoc22/util"
)

func A(input []string) int {
	sum := 0
	for i := 0; i < len(input)/3; i++ {
		l := stripBrackets(input[3*i])
		r := stripBrackets(input[3*i+1])
		if compareLists(l, r) == 1 {
			sum += i + 1
		}
	}
	return sum
}

func B(input []string) int {
	packets := []string{
		"[2]",
		"[6]",
	}
	for _, l := range input {
		if l == "" {
			continue
		}
		packets = append(packets, stripBrackets(l))
	}
	sorted := mergeSortPackets(packets)
	var two, six int
	for i, l := range sorted {
		switch l {
		case "[2]":
			two = i + 1
		case "[6]":
			six = i + 1
		}
	}
	return two * six
}

func mergeSortPackets(packets []string) []string {
	if len(packets) < 2 {
		return packets
	}
	mid := len(packets) / 2
	a := mergeSortPackets(packets[:mid])
	b := mergeSortPackets(packets[mid:])
	return merge(a, b)
}

func merge(a, b []string) []string {
	var aIdx, bIdx int
	sorted := []string{}
	for aIdx < len(a) || bIdx < len(b) {
		if aIdx == len(a) {
			sorted = append(sorted, b[bIdx])
			bIdx++
			continue
		}
		if bIdx == len(b) {
			sorted = append(sorted, a[aIdx])
			aIdx++
			continue
		}
		switch compareLists(a[aIdx], b[bIdx]) {
		case 1:
			sorted = append(sorted, a[aIdx])
			aIdx++
		case -1:
			sorted = append(sorted, b[bIdx])
			bIdx++
		case 0:
			sorted = append(sorted, a[aIdx], b[bIdx])
			aIdx++
			bIdx++
		default:
			panic("shouldn't reach here")
		}
	}
	return sorted
}

func stripBrackets(s string) string {
	if s[0] == '[' && s[len(s)-1] == ']' {
		return s[1 : len(s)-1]
	}
	return s
}

// Assumes first '[' is index 0
func matchBracket(s string) int {
	bracketCount := 1
	for i := 1; i < len(s); i++ {
		switch s[i] {
		case ']':
			bracketCount--
			if bracketCount == 0 {
				return i
			}
		case '[':
			bracketCount++
		}

	}
	panic(fmt.Sprintf("Bad format: %s", s))
}

// returns token, remaining string, and whether it's a list
func nextToken(s string) (string, string, bool) {
	if s[0] == '[' {
		end := matchBracket(s)
		remainder := util.Min(len(s), end+2)
		return s[1:end], s[remainder:], true
	}

	end := strings.Index(s, ",")
	if end == -1 {
		end = strings.Index(s, "]")
	}
	if end == -1 {
		return s, "", false
	}
	remainder := util.Min(len(s), end+1)
	return s[:end], s[remainder:], false
}

// Assumes outer brackets are stripped
// return 1 means correct, -1 incorrent, 0 inconclusive
func compareLists(left, right string) int {
	var (
		lToken, rToken string
		lList, rList   bool
	)
	for {
		// Check for end of list
		switch {
		case left == "" && right == "":
			return 0
		case left == "":
			return 1
		case right == "":
			return -1
		}
		lToken, left, lList = nextToken(left)
		rToken, right, rList = nextToken(right)
		result := 0
		switch {
		case lList || rList:
			result = compareLists(lToken, rToken)
		case lList:
			result = compareLists(lToken, fmt.Sprintf("[%s]", rToken))
		case rList:
			result = compareLists(fmt.Sprintf("[%s]", lToken), rToken)
		default: // Both integers
			l := util.Atoi(lToken)
			r := util.Atoi(rToken)
			switch {
			case l > r:
				result = -1
			case l < r:
				result = 1
			default:
				result = 0
			}
		}
		if result != 0 {
			return result
		}
	}
	panic("shouldn't reach here!")
}
