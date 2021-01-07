package stack_queue

import "errors"

type MyStack struct {
	top *nodes
	length int
}

type nodes struct {
	data interface{}
	prev *nodes
}

func NewStack() *MyStack {
	return &MyStack{nil, 0}
}

func (s MyStack) isEmpty() bool {
	return s.length == 0
}

func (s MyStack) GetLength() int {
	return s.length
}

func (s *MyStack) Push (data interface{}) {
	new := &nodes{
		data: data,
		prev: s.top,
	}
	s.top = new
	s.length ++
}

func (s *MyStack) Pop () interface{} {

	if s.isEmpty() {
		return errors.New("Stack don't have element")
	}

	rt := s.top

	s.top = rt.prev
	s.length--

	return rt.data
}

func (s MyStack) GetTop () interface{} {
	if s.isEmpty() {
		return errors.New("Stack don't have element")
	}

	return s.top.data
}