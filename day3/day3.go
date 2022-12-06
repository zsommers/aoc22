package day3

import (
	"fmt"
	"unicode/utf8"
)

func checkChars(s string) map[rune]bool {
	m := map[rune]bool{}
	for _, c := range s {
		m[c] = true
	}
	return m
}

func getWeight(r rune) int {
	bs := make([]byte, utf8.UTFMax)
	utf8.EncodeRune(bs, r)
	switch {
	case bs[0] >= 'A' && bs[0] <= 'Z':
		return int(bs[0]) - int('A') + 27
	case bs[0] >= 'a' && bs[0] <= 'z':
		return int(bs[0]) - int('a') + 1
	default:
		panic(fmt.Sprintf("%c is not valid", r))
	}
}

func A(input []string) int {
	weightSum := 0
	for _, line := range input {
		m := checkChars(line[:len(line)/2])
		for _, r := range line[len(line)/2:] {
			if _, ok := m[r]; ok {
				weightSum += getWeight(r)
				break
			}
		}
	}
	return weightSum
}

func B(input []string) int {
	weightSum := 0
	for i := 0; i < len(input); i += 3 {
		m1 := checkChars(input[i])
		m2 := checkChars(input[i+1])
		m3 := checkChars(input[i+2])
		for k := range m1 {
			_, ok2 := m2[k]
			_, ok3 := m3[k]
			if ok2 && ok3 {
				fmt.Printf("%c -> %d\n", k, getWeight(k))
				weightSum += getWeight(k)
				break
			}
		}
	}
	return weightSum
}
