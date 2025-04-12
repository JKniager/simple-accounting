package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := New[int]()
		_, ok := q.Pop()
		assert.False(t, ok)
	})

	t.Run("single item queue", func(t *testing.T) {
		q := New[int]()
		q.Push(5)
		i, ok := q.Pop()
		require.True(t, ok)
		assert.Equal(t, 5, i)
		_, ok = q.Pop()
		assert.False(t, ok)
	})

	t.Run("multi item queue", func(t *testing.T) {
		q := New[int]()
		q.Push(5)
		q.Push(3)
		q.Push(1)
		i, ok := q.Pop()
		require.True(t, ok)
		assert.Equal(t, 5, i)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 3, i)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 1, i)
		_, ok = q.Pop()
		assert.False(t, ok)
	})

	t.Run("push and pop", func(t *testing.T) {
		q := New[int]()
		q.Push(5)
		i, ok := q.Pop()
		require.True(t, ok)
		assert.Equal(t, 5, i)
		_, ok = q.Pop()
		require.False(t, ok)
		q.Push(4)
		q.Push(3)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 4, i)
		q.Push(2)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 3, i)
		q.Push(1)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 2, i)
		i, ok = q.Pop()
		require.True(t, ok)
		assert.Equal(t, 1, i)
		_, ok = q.Pop()
		require.False(t, ok)
	})
}
