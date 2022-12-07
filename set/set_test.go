// Adapted from https://github.com/golang-collections/collections for generics
package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := New[int]()

	s.Insert(5)

	assert.Equal(t, 1, s.Len())
	assert.True(t, s.Has(5))

	s.Remove(5)

	assert.Zero(t, s.Len())
	assert.False(t, s.Has(5))

	// Difference
	s1 := New(1, 2, 3, 4, 5, 6)
	s2 := New(4, 5, 6)
	s3 := s1.Difference(s2)

	assert.Equal(t, 3, s3.Len())
	assert.True(t, s3.Has(1))
	assert.True(t, s3.Has(2))
	assert.True(t, s3.Has(3))

	// Intersection
	s3 = s1.Intersection(s2)

	assert.Equal(t, 3, s3.Len())
	assert.True(t, s3.Has(4))
	assert.True(t, s3.Has(5))
	assert.True(t, s3.Has(6))

	// Union
	s4 := New(7, 8, 9)
	s3 = s2.Union(s4)

	assert.Equal(t, 6, s3.Len())
	assert.True(t, s3.Has(7))

	// Subset
	assert.True(t, s1.SubsetOf(s1))

	// Proper Subset
	assert.False(t, s1.ProperSubsetOf(s1))
}
