// Adapted from https://github.com/golang-collections/collections for generics
package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := New[int]()

	assert.Zero(t, s.Len())

	s.Push(1)

	assert.Equal(t, 1, s.Len())

	i, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, i)

	i, ok = s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, i)

	assert.Zero(t, s.Len())

	s.Push(1)
	s.Push(2)

	assert.Equal(t, 2, s.Len())
	i, ok = s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 2, i)
}
