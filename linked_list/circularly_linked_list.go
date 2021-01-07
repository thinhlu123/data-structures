package linked_list

type NodeCircle struct {
	next *NodeCircle
	data  interface{}
}

type ListCircle struct {
	head *NodeCircle
	tail *NodeCircle
	cur  *NodeCircle
}

// Push to head
func (l *ListCircle) Push (data interface{}) {
	node := &NodeCircle{
		data: data,
		next: l.head,
	}

	l.head = node
}


