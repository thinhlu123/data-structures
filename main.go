package main

import (
	"data_struct/hash_table"
	"data_struct/linked_list"
	"data_struct/stack_queue"
	"data_struct/tree"
	"fmt"
)

func main(){
	//testLinkedList()

	//testStack()

	//testQueue()

	testBSTree()

	//testHashTable()
}

// test linked list ===============================================
func testLinkedList() {
	link := linked_list.List{}
	link.Push(5)
	link.Push(9)
	link.Push(13)
	link.Push(22)
	link.Push(28)
	link.Push(36)

	link.Pop(28)

	fmt.Print(link.GetNode(1))

	fmt.Println(link.UpdateNode(13, 10))

	link.PrintList()
}

// test stack =======================================================
func testStack() {
	s := stack_queue.NewStack()
	s.Push(5)
	s.Push(9)
	s.Push(13)
	s.Push(22)
	s.Push(28)
	s.Push(36)

	fmt.Println(s.GetTop())
	s.Pop()
	fmt.Println(s.GetTop())
}

// test queue ======================================================
func testQueue() {
	q := stack_queue.NewQueue()
	q.EnQueue(5)
	q.EnQueue(9)
	q.EnQueue(13)
	q.EnQueue(22)
	q.EnQueue(28)
	q.EnQueue(36)

	fmt.Println(q.Peek())
	q.DeQueue()
	fmt.Println(q.Peek())
}

// test bst tree ===================================================
func testBSTree() {

	bst := &tree.BSTree{}

	bst.Insert(8)
	bst.Insert(4)
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(6)
	bst.Insert(1)
	bst.Insert(3)
	bst.Insert(5)
	bst.Insert(7)
	bst.Insert(9)
	bst.Insert(11)

	bst.String()

	TestInOrderTraverse(bst)
	TestPostOrderTraverse(bst)
	TestPreOrderTraverse(bst)
	TestMax(bst)
	TestMin(bst)
	TestSearch(bst)
	TestRemove(bst)
}

func isSameSlice(a, b []int) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestInOrderTraverse(bst *tree.BSTree) {
	var result []int
	bst.InOrderTraverse(func(i interface{}) {
		result = append(result, i.(int))
	})
	if !isSameSlice(result, []int{1,2,3,4,5,6,7,8,9,10,11}) {
		fmt.Println("Traversal order incorrect, got %v", result)
	}
}

func TestPreOrderTraverse(bst *tree.BSTree) {
	var result []int
	bst.PreOrderTraverse(func(i interface{}) {
		result = append(result, i.(int))
	})
	if !isSameSlice(result, []int{8,4,2,1,3,6,5,7,10,9,11}) {
		fmt.Println("Traversal order incorrect, got %v instead of %v", result, []string{"8", "4", "2", "1", "3", "6", "5", "7", "10", "9", "11"})
	}
}

func TestPostOrderTraverse(bst *tree.BSTree) {
	var result []int
	bst.PostOrderTraverse(func(i interface{}) {
		result = append(result, i.(int))
	})
	if !isSameSlice(result, []int{1,3,2,5,7,6,4,9,11,10,8}) {
		fmt.Println("Traversal order incorrect, got %v instead of %v", result, []string{"1", "3", "2", "5", "7", "6", "4", "9", "11", "10", "8"})
	}
}

func TestMin(bst *tree.BSTree) {
	if fmt.Sprintf("%s", bst.Min()) != "1" {
		fmt.Println("min should be 1")
	}
}

func TestMax(bst *tree.BSTree) {
	if fmt.Sprintf("%s", bst.Max()) != "11" {
		fmt.Println("max should be 11")
	}
}

func TestSearch(bst *tree.BSTree) {
	if !bst.Search(1) || !bst.Search(8) || !bst.Search(11) {
		fmt.Println("search not working")
	}
}

func TestRemove(bst *tree.BSTree) {
	bst.Remove(1)
	if fmt.Sprintf("%s", bst.Min()) != "2" {
		fmt.Println("min should be 2")
	}
}

// hash table ===================================================
func testHashTable() {

	ht := &hash_table.HashTable{}

	ht.Insert(1)
	fmt.Print(ht.Size())
	ht.Insert(2)
	fmt.Print(ht.Size())
	ht.Insert(100)

	fmt.Print(ht.IsExist(2))
	fmt.Print(ht.Size())

}