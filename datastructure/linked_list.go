package datastructure

import (
	"errors"
)

// LinkedListNode specified a node lin linked list
type LinkedListNode struct {
	Data interface{}
	Next *LinkedListNode
}

// LinkedList specified a unidirection linked list
type LinkedList struct {
	len int
	Tail, Head *LinkedListNode
}

// NewLinkedList create a linked list instance
func NewLinkedList() * LinkedList {
	return &LinkedList{
		len: 0,
		Head: nil,
		Tail: nil,
	}
}

// Append insert a LinkedListNode to the tail of the list
func (l *LinkedList) Append(data interface{}) {
	newNode := &LinkedListNode{
		Data: data,
		Next: nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}

	l.len++
}

// Remove delete a LinkedListNode at given index in the list
func (l *LinkedList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= l.len {
		return 0, errors.New("index is out of range")
	}

	i := 0
	found := l.Head
	pre := l.Head
	for i < index {
		i++
		pre = found
		found = found.Next
	}
	pre.Next = found.Next
	found.Next = nil
	return found.Data, nil
}

// Len return the length of the list
func (l *LinkedList) Len() int {
	return l.len
}