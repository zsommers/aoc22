package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var MaxInt = int(^uint(0) >> 1)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func ReadLines(path string) []string {
	f, err := os.Open(path)
	CheckErr(err)
	defer f.Close()

	s := bufio.NewScanner(f)
	ls := []string{}
	for s.Scan() {
		ls = append(ls, s.Text())
	}
	CheckErr(s.Err())

	return ls
}

func Atoi(s string) (i int) {
	var err error
	if i, err = strconv.Atoi(s); err != nil {
		panic(err.Error())
	}
	return
}

func Min(is ...int) int {
	m := MaxInt
	for _, i := range is {
		if i < m {
			m = i
		}
	}
	return m
}

func Max(is ...int) int {
	m := -MaxInt - 1
	for _, i := range is {
		if i > m {
			m = i
		}
	}
	return m
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ReadIntString(s string) []int {
	result := []int{}
	for _, n := range strings.Split(s, ",") {
		result = append(result, Atoi(n))
	}
	return result
}

func SumInts(is ...int) int {
	sum := 0
	for _, i := range is {
		sum += i
	}
	return sum
}

func HighestInts(count int, vals ...int) []int {
	if count > len(vals) || count < 0 {
		panic(fmt.Sprintf("Cannot select %d items from slice of length %d", count, len(vals)))
	}
	result := make([]int, count)
	insert := func(v, index int) {
		for i := len(result) - 1; i > index; i-- {
			result[i] = result[i-1]
		}
		result[index] = v
	}
	for _, v := range vals {
		for ri := 0; ri < len(result); ri++ {
			if v > result[ri] {
				insert(v, ri)
				break
			}
		}
	}
	return result
}
