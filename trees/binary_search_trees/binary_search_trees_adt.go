package binarysearchtrees

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func NewNode(value int) *Node {
	node := new(Node)

	node.value = value
	node.left = nil
	node.right = nil

	return node
}

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	bst := new(BinarySearchTree)

	bst.root = nil

	return bst
}

func (bst *BinarySearchTree) InsertIterative(value int) {
	newNode := NewNode(value)

	parent := bst.root
	cursor := bst.root

	for cursor != nil {
		parent = cursor
		if cursor.value == value {
			return
		} else if cursor.value < value {
			cursor = cursor.right
		} else {
			cursor = cursor.left
		}
	}

	if parent == nil {
		bst.root = newNode
	} else {
		if parent.value < value {
			parent.right = newNode
		} else {
			parent.left = newNode
		}
	}
}

func (bst *BinarySearchTree) insertRecursive(node *Node, value int) *Node {
	if node == nil {
		return NewNode(value)
	}

	if node.value < value {
		node.right = bst.insertRecursive(node.right, value)
	} else {
		node.left = bst.insertRecursive(node.left, value)
	}

	return node
}

func (bst *BinarySearchTree) InsertRecursive(value int) {
	bst.root = bst.insertRecursive(bst.root, value)
}

func (bst *BinarySearchTree) inOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	bst.inOrderTraversal(node.left)
	fmt.Printf("%d, ", node.value)
	bst.inOrderTraversal(node.right)
}

func (bst *BinarySearchTree) height(node *Node) int {
	if node == nil {
		return 0
	}

	x := bst.height(node.left)
	y := bst.height(node.right)

	return int(math.Max(float64(x), float64(y))) + 1
}

func (bst *BinarySearchTree) Height() int {
	return bst.height(bst.root)
}

func (bst *BinarySearchTree) count(node *Node) int {
	if node == nil {
		return 0
	}

	x := bst.count(node.left)
	y := bst.count(node.right)

	return x + y + 1
}

func (bst *BinarySearchTree) Count() int {
	return bst.count(bst.root)
}

func (bst *BinarySearchTree) inOrderPredecessor(node *Node) *Node {
	for node != nil && node.right != nil {
		node = node.right
	}

	return node
}

func (bst *BinarySearchTree) inOrderSuccessor(node *Node) *Node {
	for node != nil && node.left != nil {
		node = node.left
	}

	return node
}

func (bst *BinarySearchTree) delete(deletedNode **Node, node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if node.value < value {
		node.right = bst.delete(deletedNode, node.right, value)
	} else if node.value > value {
		node.left = bst.delete(deletedNode, node.left, value)
	} else {
		*deletedNode = node
		if node.left == nil && node.right == nil {
			return nil
		}

		lH := bst.height(node.left)
		rH := bst.height(node.right)

		nDeletedNode := new(*Node)
		if lH > rH {
			pre := bst.inOrderPredecessor(node.left)
			nodeLeft := bst.delete(nDeletedNode, node.left, pre.value)
			nNode := *nDeletedNode
			if nNode != nil {
				nNode.left = nodeLeft
				nNode.right = node.right
				return nNode
			}
		} else {
			suc := bst.inOrderSuccessor(node.right)
			nodeRight := bst.delete(nDeletedNode, node.right, suc.value)
			nNode := *nDeletedNode
			if nNode != nil {
				nNode.left = node.left
				nNode.right = nodeRight
				return nNode
			}
		}
	}

	return node
}

func (bst *BinarySearchTree) Delete(value int) *Node {
	deletedNode := new(*Node)

	bst.root = bst.delete(deletedNode, bst.root, value)

	return *deletedNode
}

func (bst *BinarySearchTree) InOrderTraversal() {
	if bst.root == nil {
		return
	}

	fmt.Println("InOrderTraversal!")
	bst.inOrderTraversal(bst.root)
	fmt.Println()
}

func RunBinarySearchTrees(run bool) {
	if !run {
		return
	}

	bst1 := NewBinarySearchTree()
	bstArray1 := []int{9, 15, 5, 20, 16, 8, 12, 3, 6}
	for _, v := range bstArray1 {
		bst1.InsertIterative(v)
	}
	// bst1.InOrderTraversal()
	for _, v := range bstArray1 {
		bst1.InOrderTraversal()
		node := bst1.Delete(v)
		fmt.Printf("left = %t, right = %t, value = %d\n", node.left == nil, node.right == nil, node.value)
	}

	// bst2 := NewBinarySearchTree()
	// bstArray2 := []int{9, 15, 5, 20, 16, 8, 12, 3, 6}
	// for _, v := range bstArray2 {
	// 	bst2.InsertRecursive(v)
	// }
	// bst2.InOrderTraversal()
}
