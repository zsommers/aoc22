package day1

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rawInput = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

var input = strings.Split(rawInput, "\n")

func TestA(t *testing.T) {
	assert.Equal(t, 24000, A(input))
}

func TestB(t *testing.T) {
	assert.Equal(t, 45000, B(input))
}
