package tree

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)

// HashKey to convert interface to uint
func HashKey(key interface{}) (uint, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return 0, err
	}
	data := buf.Bytes()

	var a, b, c uint
	a, b = 0x9e3779b9, 0x9e3779b9
	c = 0
	i := 0

	for i = 0; i < len(data)-12; {
		a += uint(data[i]) | uint(data[i+1]<<8) | uint(data[i+2]<<16) | uint(data[i+3]<<24)
		i += 4
		b += uint(data[i]) | uint(data[i+1]<<8) | uint(data[i+2]<<16) | uint(data[i+3]<<24)
		i += 4
		c += uint(data[i]) | uint(data[i+1]<<8) | uint(data[i+2]<<16) | uint(data[i+3]<<24)

		a, b, c = mix(a, b, c)
	}

	c += uint(len(data))

	if i < len(data) {
		a += uint(data[i])
		i++
	}
	if i < len(data) {
		a += uint(data[i]) << 8
		i++
	}
	if i < len(data) {
		a += uint(data[i]) << 16
		i++
	}
	if i < len(data) {
		a += uint(data[i]) << 24
		i++
	}

	if i < len(data) {
		b += uint(data[i])
		i++
	}
	if i < len(data) {
		b += uint(data[i]) << 8
		i++
	}
	if i < len(data) {
		b += uint(data[i]) << 16
		i++
	}
	if i < len(data) {
		b += uint(data[i]) << 24
		i++
	}

	if i < len(data) {
		c += uint(data[i]) << 8
		i++
	}
	if i < len(data) {
		c += uint(data[i]) << 16
		i++
	}
	if i < len(data) {
		c += uint(data[i]) << 24
		i++
	}

	a, b, c = mix(a, b, c)
	return c, nil
}

func mix(a, b, c uint) (uint, uint, uint) {
	a -= b
	a -= c
	a ^= (c >> 13)
	b -= c
	b -= a
	b ^= (a << 8)
	c -= a
	c -= b
	c ^= (b >> 13)
	a -= b
	a -= c
	a ^= (c >> 12)
	b -= c
	b -= a
	b ^= (a << 16)
	c -= a
	c -= b
	c ^= (b >> 5)
	a -= b
	a -= c
	a ^= (c >> 3)
	b -= c
	b -= a
	b ^= (a << 10)
	c -= a
	c -= b
	c ^= (b >> 15)
	return a, b, c
}

type Node struct {
	key   uint
	value interface{}
	left  *Node
	right *Node
}

// root tree
type BSTree struct {
	root *Node
	lock sync.RWMutex
}

// Insert inserts the Item t in the tree
func (bst *BSTree) Insert(value interface{}) {
	key, err := HashKey(value)
	if err != nil {
		return
	}

	bst.lock.Lock()
	defer bst.lock.Unlock()
	n := &Node{key, value, nil, nil}
	if bst.root == nil {
		bst.root = n
	} else {
		insertNode(bst.root, n)
	}
}

func insertNode(node, newNode *Node) {
	if newNode.key < node.key {
		if node.left == nil {
			node.left = newNode
		} else {
			insertNode(node.left, newNode)
		}
	} else {
		if node.right == nil {
			node.right = newNode
		} else {
			insertNode(node.right, newNode)
		}
	}
}

// InOrderTraverse visits all nodes with in-order traversing
func (bst *BSTree) InOrderTraverse(f func(interface{})) {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	inOrderTraverse(bst.root, f)
}

func inOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		inOrderTraverse(n.left, f)
		f(n.value)
		inOrderTraverse(n.right, f)
	}
}

// PreOrderTraverse visits all nodes with pre-order traversing
func (bst *BSTree) PreOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	preOrderTraverse(bst.root, f)
}

func preOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		f(n.value)
		preOrderTraverse(n.left, f)
		preOrderTraverse(n.right, f)
	}
}

// PostOrderTraverse visits all nodes with post-order traversing
func (bst *BSTree) PostOrderTraverse(f func(interface{})) {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	postOrderTraverse(bst.root, f)
}

func postOrderTraverse(n *Node, f func(interface{})) {
	if n != nil {
		postOrderTraverse(n.left, f)
		postOrderTraverse(n.right, f)
		f(n.value)
	}
}

// Min returns the node with min value stored in the tree
func (bst *BSTree) Min() interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.left == nil {
			return n.value
		}
		n = n.left
	}
}

func (bst *BSTree) Max() interface{} {
	bst.lock.RLock()
	defer bst.lock.RUnlock()
	n := bst.root
	if n == nil {
		return nil
	}
	for {
		if n.right == nil {
			return n.value
		}
		n = n.right
	}
}

// Search returns true if the node t exists in the tree
func (bst *BSTree) Search(value interface{}) bool {
	key, err := HashKey(value)
	if err != nil {
		return false
	}

	bst.lock.RLock()
	defer bst.lock.RUnlock()
	return search(bst.root, key)
}

func search(n *Node, key uint) bool {
	if n == nil {
		return false
	}
	if key < n.key {
		return search(n.left, key)
	}
	if key > n.key {
		return search(n.right, key)
	}
	return true
}

// Remove removes the node with key `key` from the tree
func (bst *BSTree) Remove(value interface{}) {
	key, err := HashKey(value)
	if err != nil {
		return
	}

	bst.lock.Lock()
	defer bst.lock.Unlock()
	remove(bst.root, key)
}

func remove(node *Node, key uint) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = remove(node.left, key)
		return node
	}
	if key > node.key {
		node.right = remove(node.right, key)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.key, node.value = leftmostrightside.key, leftmostrightside.value
	node.right = remove(node.right, node.key)
	return node
}

// String print tree
func (bst *BSTree) String() {
	bst.lock.Lock()
	defer bst.lock.Unlock()
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%v\n", n.value)
		stringify(n.right, level)
	}
}