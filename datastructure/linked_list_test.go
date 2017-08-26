package datastructure_test

import (
	"testing"
	ds "github.com/tossmilestone/go-algorithm/datastructure"
	assert "github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	l := ds.NewLinkedList()
	assert.NotNil(t, l)
	assert.Nil(t, l.Head)
	assert.Nil(t, l.Tail)
}

func TestLinkedList_Append(t *testing.T) {
	l := ds.NewLinkedList()
	l.Append(1)
	assert.Equal(t, l.Len(), 1)
	l.Append("test")
	assert.Equal(t, l.Len(), 2)
}

func TestLinkedList_Remove(t *testing.T) {
	l := ds.NewLinkedList()
	l.Append(1)
	l.Append("test")
	l.Append(33)
	d, err := l.Remove(-1)
	assert.NotNil(t, err)
	assert.Nil(t, d)
	l.Remove(3)
	assert.NotNil(t, err)
	assert.Nil(t, d)
	d, err = l.Remove(0)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, 1, d.(int))
	assert.Equal(t, 2, l.Len())
	d, err = l.Remove(1)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, 33, d.(int))
	assert.Equal(t, 1, l.Len())
	d, err = l.Remove(0)
	assert.Nil(t, err)
	assert.NotNil(t, d)
	assert.Equal(t, "test", d.(string))
	assert.Equal(t, 0, l.Len())
}

func TestLinkedList_Len(t *testing.T) {
	l := ds.NewLinkedList()
	l.Append(1)
	assert.Equal(t, 1, l.Len())
	l.Append("test")
	assert.Equal(t, 2, l.Len())
	l.Append(33)
	assert.Equal(t, 3, l.Len())
	l.Remove(2)
	assert.Equal(t, 2, l.Len())
	l.Append(123)
	assert.Equal(t, 3, l.Len())
}