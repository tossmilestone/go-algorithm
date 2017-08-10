package datastructure_test

import (
	"testing"
	ds "github.com/tossmilestone/go-algorithm/datastructure"
	assert "github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	s := ds.NewStack(5)
	assert.NotNil(t, s)

	s = ds.NewStack(1)
	assert.NotNil(t, s)
}

func TestStack_Push(t *testing.T) {
	s := ds.NewStack(5)
	assert.Equal(t, 0, s.Size())
	err := s.Push(1)
	assert.Nil(t, err)
	assert.Equal(t, 1, s.Size())
	err = s.Push(2)
	assert.Nil(t, err)
	assert.Equal(t, 2, s.Size())
	s.Push(3)
	s.Push(4)
	s.Push(5)
	assert.Equal(t, 5, s.Size())
	err = s.Push(6)
	assert.NotNil(t, err)
}

func TestStack_Pop(t *testing.T) {
	s := ds.NewStack(5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	assert.Equal(t, 4, s.Size())
	n, err := s.Pop()
	assert.Equal(t, 4, n)
	assert.Equal(t, 3, s.Size())
	assert.Nil(t, err)
	n, err = s.Pop()
	assert.Equal(t, 3, n)
	assert.Equal(t, 2, s.Size())
	assert.Nil(t, err)
	n, err = s.Pop()
	assert.Equal(t, 2, n)
	assert.Equal(t, 1, s.Size())
	assert.Nil(t, err)
	n, err = s.Pop()
	assert.Equal(t, 1, n)
	assert.Equal(t, 0, s.Size())
	assert.Nil(t, err)
	n, err = s.Pop()
	assert.Equal(t, 0, n)
	assert.NotNil(t, err)
}

func TestStack_Peek(t *testing.T) {
	s := ds.NewStack(5)
	s.Push(1)
	n, err := s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, n)
	s.Push(2)
	n, err = s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 2, n)
	s.Push(3)
	n, err = s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
	s.Push(4)
	n, err = s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 4, n)

	s.Pop()
	n, err = s.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 3, n)
}

func TestStack_IsEmpty(t *testing.T) {
	s := ds.NewStack(5)
	assert.True(t, s.IsEmpty())
	s.Push(1)
	assert.False(t, s.IsEmpty())
	s.Pop()
	assert.True(t, s.IsEmpty())
}

func TestStack_IsFull(t *testing.T) {
	s := ds.NewStack(5)
	assert.False(t, s.IsFull())
	s.Push(1)
	assert.False(t, s.IsFull())
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	assert.True(t, s.IsFull())
	s.Pop()
	assert.False(t, s.IsFull())
}

func TestStack_Size(t *testing.T) {
	s := ds.NewStack(5)
	assert.Equal(t, 0, s.Size())
	s.Push(1)
	assert.Equal(t, 1, s.Size())
	s.Push(1)
	assert.Equal(t, 2, s.Size())
	s.Push(1)
	assert.Equal(t, 3, s.Size())
	s.Push(1)
	assert.Equal(t, 4, s.Size())
	s.Push(1)
	assert.Equal(t, 5, s.Size())
	s.Pop()
	assert.Equal(t, 4, s.Size())
	s.Pop()
	assert.Equal(t, 3, s.Size())
	s.Pop()
	assert.Equal(t, 2, s.Size())
	s.Pop()
	assert.Equal(t, 1, s.Size())
	s.Pop()
	assert.Equal(t, 0, s.Size())
}