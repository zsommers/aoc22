package day6

import "github.com/zsommers/aoc22/set"

func findUniqueCharSeq(count int, s string) int {
	for i := count - 1; i < len(s); i++ {
		nums := set.New[byte]()
		for j := i; j > i-count; j-- {
			nums.Insert(s[j])
		}
		if nums.Len() == count {
			return i + 1
		}
	}
	return 0
}

func A(input []string) int {
	return findUniqueCharSeq(4, input[0])
}

func B(input []string) int {
	return findUniqueCharSeq(14, input[0])
}
