package stack_queue

import "errors"

type MyQueue struct {
	start *nodeq
	end *nodeq
	length int
}

type nodeq struct {
	data interface{}
	next *nodeq
}

func NewQueue() *MyQueue {
	return &MyQueue{nil, nil, 0}
}

func (q MyQueue) GetLength () int {
	return q.length
}

func (q *MyQueue) DeQueue () interface{} {
	if q.length == 0 {
		return errors.New("Queue don't have element")
	}

	n := q.start
	if q.length == 1 {
		q.start = nil
		q.end = nil
	} else {
		q.start = q.start.next
	}

	q.length --
	return n.data
}

func (q *MyQueue) EnQueue (data interface{}) {
	n := &nodeq{
		data: data,
		next: nil,
	}

	if q.length == 0 {
		q.start = n
		q.end = n
	} else {
		q.end.next = n
		q.end = n
	}

	q.length++
}

// get start element
func (q MyQueue) Peek () interface{} {
	return q.start.data
}