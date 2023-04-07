package queue

import (
	"errors"
	"fmt"
)

// Que Errors
var (
	ErrQueueIsFull  = errors.New("queue is already full")
	ErrQueueIsEmpty = errors.New("queue is empty")
)

// queue is the go implementation of queue data type
type queue[k any] struct {
	front int // Front pointer to front of the Queue
	rear  int // Rear pointer to back of the queue
	size  int // size for defining the size of the queue
	items []k // Items where actual queue items are stored
}

// NewQueue creates new queue of the described size
func NewQueue[k any](size int) queue[k] {
	i := make([]k, size)
	return queue[k]{
		front: -1,
		rear:  -1,
		size:  size,
		items: i,
	}
}

// IsEmpty checks if the given queue is emtpy or not
func (q *queue[k]) IsEmpty() bool {
	return q.front == -1
}

// IsFull checks if the given queue is full or not
func (q *queue[k]) IsFull() bool {
	return q.rear+1 == q.size
}

// Enqueue will add the element to the back of the queue
func (q *queue[k]) Enqueue(e k) (err error) {
	if q.IsFull() {
		return ErrQueueIsFull
	}
	// if queue is empty set the front to be 1st element or 0 index
	if q.IsEmpty() {
		q.front = 0
	}
	q.rear += 1         // increase the rear element by 1
	q.items[q.rear] = e // set the rear element to be the new element
	return nil
}

func (q *queue[k]) Dequeue() (err error) {
	if q.IsEmpty() {
		return ErrQueueIsEmpty
	}
	var zero k
	q.items[q.front] = zero
	q.front += 1 // set the new front

	// if this is the last element
	if q.size-1 == q.front && q.front == q.rear {
		q.front, q.rear = -1, -1
	}
	return nil
}

func (q *queue[k]) Peek() k {
	var zero k
	if q.IsEmpty() {
		return zero
	}
	return q.items[q.rear]
}

// String to implement custom format for queue struct
func (q queue[k]) String() string {
	return fmt.Sprintf("FRONT = %d REAR = %d SIZE = %d ITEMS = %v", q.front, q.rear, q.size, q.items)
}

var _ fmt.Stringer = (*queue[int])(nil)
