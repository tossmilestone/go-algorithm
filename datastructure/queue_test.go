package datastructure_test

import (
	"testing"
	ds "github.com/tossmilestone/go-algorithm/datastructure"
	assert "github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	q := ds.NewQueue(5)
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, 5, q.Capacity)
	assert.Equal(t, true, q.IsEmpty())
	assert.Equal(t, false, q.IsFull())

	q = ds.NewQueue(10)
	assert.NotNil(t, q)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, 10, q.Capacity)
	assert.Equal(t, true, q.IsEmpty())
	assert.Equal(t, false, q.IsFull())
}

func TestQueue_Enqueue(t *testing.T) {
	q := ds.NewQueue(5)
	err := q.Enqueue(1)
	assert.Nil(t, err)
	err = q.Enqueue(2)
	assert.Nil(t, err)
	e, err := q.Dequeue()
	assert.Equal(t, 1, e)
	e, err = q.Dequeue()
	assert.Equal(t, 2, e)

	q = ds.NewQueue(2)
	err = q.Enqueue(3)
	assert.Nil(t, err)
	err = q.Enqueue(4)
	assert.Nil(t, err)
	err = q.Enqueue(5)
	assert.NotNil(t, err)
}

func TestQueue_Dequeue(t *testing.T) {
	q := ds.NewQueue(8)
	q.Enqueue(3)
	q.Enqueue(2)
	e, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 3, e)
	e, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, 2, e)
	e, err = q.Dequeue()
	assert.NotNil(t, err)
}

func TestQueue_Size(t *testing.T) {
	q := ds.NewQueue(10)
	assert.Equal(t, 0, q.Size())
	q.Enqueue(5)
	q.Enqueue(3)
	q.Enqueue(3)
	assert.Equal(t, 3, q.Size())
	q.Dequeue()
	assert.Equal(t, 2, q.Size())
	q.Dequeue()
	assert.Equal(t, 1, q.Size())
	q.Dequeue()
	assert.Equal(t, 0, q.Size())
}

func TestQueue_IsEmpty(t *testing.T) {
	q := ds.NewQueue(10)
	assert.Equal(t, true, q.IsEmpty())
	q.Enqueue(5)
	assert.Equal(t, false, q.IsEmpty())
	q.Enqueue(3)
	assert.Equal(t, false, q.IsEmpty())
	q.Enqueue(3)
	assert.Equal(t, false, q.IsEmpty())
	q.Dequeue()
	assert.Equal(t, false, q.IsEmpty())
	q.Dequeue()
	assert.Equal(t, false, q.IsEmpty())
	q.Dequeue()
	assert.Equal(t, true, q.IsEmpty())
}

func TestQueue_IsFull(t *testing.T) {
	q := ds.NewQueue(3)
	assert.Equal(t, false, q.IsFull())
	q.Enqueue(5)
	assert.Equal(t, false, q.IsFull())
	q.Enqueue(3)
	assert.Equal(t, false, q.IsFull())
	q.Enqueue(3)
	assert.Equal(t, true, q.IsFull())
	q.Dequeue()
	assert.Equal(t, false, q.IsFull())
	q.Dequeue()
	assert.Equal(t, false, q.IsFull())
	q.Dequeue()
	assert.Equal(t, false, q.IsFull())
}