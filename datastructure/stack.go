package datastructure

import "errors"

// Stack provides a LIFO container
type Stack struct {
	data []int
	top int
}

// NewStack creates a Stack instance
func NewStack(capacity int) *Stack {
	return &Stack{
		data: make([]int, capacity),
		top: -1,
	}
}

// Push insert an object into the stack
func (s *Stack)Push(e int) error {
	if s.top >= len(s.data) {
		return errors.New("stack is full")
	}

	s.top++
	s.data[s.top] = e
	return nil
}

// Pop get and remove an object from the stack
func (s *Stack)Pop() (int, error) {
	if s.top < 0 {
		return 0, errors.New("stack is empty")
	}
	top := s.top
	s.top--
	return s.data[top], nil
}

// Peek get the top object from the stack
func (s *Stack)Peek() (int, error) {
	return s.data[s.top], nil
}

// Size return the size of the stack
func (s *Stack)Size() int {
	return s.top+1
}

// IsEmpty return if the stack is empty
func (s *Stack)IsEmpty() bool {
	return s.top == -1
}

// IsFull return if the stack is full
func (s *Stack)IsFull() bool {
	return s.top == len(s.data)
}

