package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Queue(t *testing.T) {
	queue := New()

	assert.True(t, queue.IsEmpty(), "IsEmpty() should return true")

	first, last := 0, 10

	for i := first; i < last; i++ {
		queue.Push(i)
		assert.Equal(t, i-first+1, queue.Len(), "values should be the same")
	}

	for i := first; i < last; i++ {
		assert.Equal(t, i, queue.Pop(), "values should be the same")
	}

	assert.True(t, queue.IsEmpty(), "IsEmpty() should return true")
}
