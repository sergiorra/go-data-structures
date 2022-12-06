package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stack(t *testing.T) {
	stack := New()

	assert.True(t, stack.IsEmpty(), "IsEmpty() should return true")

	first, last := 0, 10

	for i := first; i < last; i++ {
		stack.Push(i)
		assert.Equal(t, i-first+1, stack.Len(), "values should be the same")
	}

	for i := last - 1; i >= first; i-- {
		assert.Equal(t, i, stack.Pop(), "values should be the same")
	}

	assert.True(t, stack.IsEmpty(), "IsEmpty() should return true")
}
