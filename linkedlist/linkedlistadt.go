package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	n := new(Node)

	n.value = value
	n.next = nil

	return n
}

type LinkedListADT struct {
	startNode *Node
}

func CreateLinkedListADT() *LinkedListADT {
	ll := new(LinkedListADT)

	ll.startNode = nil

	return ll
}

func (ll *LinkedListADT) Display() {
	var cursor *Node
	cursor = ll.startNode

	for cursor != nil {
		fmt.Printf("%d, ", cursor.value)
		cursor = cursor.next
	}

	fmt.Println()
}

func (ll *LinkedListADT) Append(value int) {
	newNode := NewNode(value)

	if ll.startNode == nil {
		ll.startNode = newNode
		return
	}

	var cursor *Node
	cursor = ll.startNode

	for cursor.next != nil {
		cursor = cursor.next
	}

	cursor.next = newNode
}

func (ll *LinkedListADT) AppendArray(values []int) {
	if len(values) == 0 {
		return
	}

	idx := 0
	if ll.startNode == nil {
		ll.startNode = NewNode(values[idx])
		idx++
	}

	var cursor *Node
	cursor = ll.startNode

	for cursor.next != nil {
		cursor = cursor.next
	}

	for i := idx; i < len(values); i++ {
		cursor.next = NewNode(values[i])
		cursor = cursor.next
	}
}

func RunLinkedListADT(run bool) {
	if !run {
		return
	}

	ll := CreateLinkedListADT()
	ll.Display()
}
