package datastructure

import "errors"


type Queue struct {
	data []int
	head, tail, Capacity int
}

func NewQueue(size int) *Queue {
	return &Queue{
		data: make([]int, size + 1),
		head: 0,
		tail: 0,
		Capacity: size,
	}
}

func (q *Queue) Enqueue(e int) error {
	tail := (q.tail + 1) % (q.Capacity + 1)
	if tail == q.head {
		return errors.New("Queue is full!")
	}
	q.data[q.tail] = e
	q.tail = tail
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if q.head == q.tail {
		return 0, errors.New("Queue is empty!")
	}
	head := q.head
	q.head = (q.head + 1) % (q.Capacity + 1)
	return q.data[head], nil
}

func (q *Queue) Size() int {
	return (q.tail - q.head + q.Capacity + 1) % (q.Capacity + 1)
}

func (q *Queue) IsEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) IsFull() bool {
	return q.head == (q.tail + 1) % (q.Capacity + 1)
}
