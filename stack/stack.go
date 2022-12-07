// Adapted from https://github.com/golang-collections/collections for ints
package stack

type (
	Stack[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		value T
		prev  *node[T]
	}
)

// Create a new stack
func New[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

// Return the number of items in the stack
func (this *Stack[T]) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack[T]) Peek() (T, bool) {
	var t T
	if this.length == 0 {
		return t, false
	}
	return this.top.value, true
}

// Pop the top item of the stack and return it
func (this *Stack[T]) Pop() (T, bool) {
	var t T
	if this.length == 0 {
		return t, false
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value, true
}

// Push a value onto the top of the stack
func (this *Stack[T]) Push(value T) {
	n := &node[T]{value, this.top}
	this.top = n
	this.length++
}

func Equal[T comparable](a, b *Stack[T]) bool {
	if a.Len() != b.Len() {
		return false
	}

	for i := 0; i < a.Len(); i++ {
		aa, _ := a.Pop()
		bb, _ := b.Pop()
		if aa != bb {
			return false
		}
	}
	return true
}
