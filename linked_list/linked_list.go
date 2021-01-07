package linked_list

import "fmt"

type Node struct {
	next *Node
	data  interface{}
}

type List struct {
	head *Node
	tail *Node
}

// Push to head
func (l *List) Push (data interface{}) {
	node := &Node{
		data: data,
		next: l.head,
	}

	l.head = node
}

// Pop node
func (l *List) Pop (data interface{}) *Node {
	i := l.head
	var prev *Node
	for i.data != data {
		prev = i
		i = i.next
	}

	prev.next = i.next

	return prev
}

func (l *List) PrintList () {
	fmt.Println("List: ")
	i := l.head
	for i != nil {
		fmt.Println(i.data)
		i = i.next
	}
}

func (l *List) GetNode (data interface{}) *Node {
	i := l.head
	for i.data != data {
		if i.next == nil {
			return nil
		}

		i = i.next
	}

	return i
}

func (l *List) UpdateNode (filter, updater interface{}) *Node {
	node := l.GetNode(filter)

	if node != nil {
		node.data = updater
	}

	return node
}