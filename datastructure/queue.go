package datastructure

import "errors"

// Queue specifies a FIFO queue
type Queue struct {
	data                 []int
	head, tail, Capacity int
}

// NewQueue create a Queue instance
func NewQueue(size int) *Queue {
	return &Queue{
		data:     make([]int, size+1),
		head:     0,
		tail:     0,
		Capacity: size,
	}
}

// Enqueue eject an object from the head of the queue
func (q *Queue) Enqueue(e int) error {
	tail := (q.tail + 1) % (q.Capacity + 1)
	if tail == q.head {
		return errors.New("queue is full")
	}
	q.data[q.tail] = e
	q.tail = tail
	return nil
}

// Dequeue push an object into the tail of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.head == q.tail {
		return 0, errors.New("queue is empty")
	}
	head := q.head
	q.head = (q.head + 1) % (q.Capacity + 1)
	return q.data[head], nil
}

// Size returns the size of a Queue
func (q *Queue) Size() int {
	return (q.tail - q.head + q.Capacity + 1) % (q.Capacity + 1)
}

// IsEmpty returns if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

// IsFull returns if the queue is full
func (q *Queue) IsFull() bool {
	return q.head == (q.tail+1)%(q.Capacity+1)
}
