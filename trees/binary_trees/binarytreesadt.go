package binarytrees

import (
	"fmt"
	"math"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func NewNode(value int) *Node {
	n := new(Node)

	n.left = nil
	n.right = nil
	n.value = value

	return n
}

type BinaryTree struct {
	root *Node
}

func NewBinaryTree() *BinaryTree {
	bt := new(BinaryTree)

	bt.root = nil

	return bt
}

func (bt *BinaryTree) preOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	fmt.Printf("%d, ", node.value)
	bt.preOrderTraversal(node.left)
	bt.preOrderTraversal(node.right)
}

func (bt *BinaryTree) PreOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("PreOrderTraversal!")
	bt.preOrderTraversal(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) IterativePreOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("IterativePreOrderTraversal!")

	stack := &Stack{}

	node := bt.root
	stack.Push(node)

	for !stack.IsEmpty() {
		node = stack.Pop()
		if node == nil {
			break
		}

		fmt.Printf("%d, ", node.value)
		if node.right != nil {
			stack.Push(node.right)
		}
		if node.left != nil {
			stack.Push(node.left)
		}
	}

	fmt.Println()
}

func (bt *BinaryTree) IterativePreOrderTraversal2() {
	if bt.root == nil {
		return
	}

	fmt.Println("IterativePreOrderTraversal! 2")

	stack := &Stack{}

	node := bt.root

	for node != nil || !stack.IsEmpty() {
		if node != nil {
			fmt.Printf("%d, ", node.value)
			stack.Push(node)
			node = node.left
		} else {
			node = stack.Pop()
			node = node.right
		}
	}

	fmt.Println()
}

func (bt *BinaryTree) inOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	bt.inOrderTraversal(node.left)
	fmt.Printf("%d, ", node.value)
	bt.inOrderTraversal(node.right)
}

func (bt *BinaryTree) InOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("InOrderTraversal!")
	bt.inOrderTraversal(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) IterativeInOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("IterativeInOrderTraversal!")

	stack := &Stack{}

	node := bt.root

	for node != nil || !stack.IsEmpty() {
		if node != nil {
			stack.Push(node)
			node = node.left
		} else {
			node = stack.Pop()
			fmt.Printf("%d, ", node.value)
			node = node.right
		}
	}

	fmt.Println()
}

func (bt *BinaryTree) postOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	bt.postOrderTraversal(node.left)
	bt.postOrderTraversal(node.right)
	fmt.Printf("%d, ", node.value)
}

func (bt *BinaryTree) PostOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("PostOrderTraversal!")
	bt.postOrderTraversal(bt.root)
	fmt.Println()
}

func (bt *BinaryTree) LevelOrderTraversal() {
	if bt.root == nil {
		return
	}

	fmt.Println("LevelOrderTraversal!")

	queue := &Queue{}

	node := bt.root
	queue.Enqueue(node)

	for !queue.IsEmpty() {
		node = queue.Dequeue()
		if node == nil {
			break
		}

		fmt.Printf("%d, ", node.value)
		if node.left != nil {
			queue.Enqueue(node.left)
		}
		if node.right != nil {
			queue.Enqueue(node.right)
		}
	}

	fmt.Println()
}

func (bt *BinaryTree) height(node *Node) int {
	if node == nil {
		return 0
	}

	lh := bt.height(node.left)
	rh := bt.height(node.right)

	return int(math.Max(float64(lh), float64(rh))) + 1
}

func (bt *BinaryTree) Height() int {
	return bt.height(bt.root)
}

func (bt *BinaryTree) count(node *Node) int {
	if node == nil {
		return 0
	}

	lc := bt.count(node.left)
	rc := bt.count(node.right)

	return lc + rc + 1
}

func (bt *BinaryTree) Count() int {
	return bt.count(bt.root)
}

func (bt *BinaryTree) leafNodeCount(node *Node) int {
	if node == nil {
		return 0
	}

	lLNC := bt.leafNodeCount(node.left)
	rLNC := bt.leafNodeCount(node.right)

	total := lLNC + rLNC

	if node.left == nil && node.right == nil {
		total++
	}

	return total
}

func (bt *BinaryTree) LeafNodeCount() int {
	return bt.leafNodeCount(bt.root)
}

/*
						[2]

			[4]						[6]

	[8]	    	[3]		  [9]					[5]

	arr = [2, 4, 6, 8, 3, 9, 5]
	idx =  0, 1, 2, 3, 4, 5, 6
*/

func (bt *BinaryTree) GenerateTreeFromArray(treeArray []any) {
	H := map[int]*Node{}

	for i, v := range treeArray {
		if v == nil {
			continue
		}

		_v, ok := v.(int)
		if !ok {
			continue
		}

		if i == 0 {
			bt.root = NewNode(_v)
			H[0] = bt.root
		} else {
			parentIndex := ((i + 1) / 2) - 1
			parentLeftChild := (2 * (parentIndex + 1)) - 1
			parentRightChild := (2 * (parentIndex + 1))

			parentNode := H[parentIndex]
			if parentNode != nil {
				switch i {
				case parentLeftChild:
					parentNode.left = NewNode(_v)
					H[i] = parentNode.left
				case parentRightChild:
					parentNode.right = NewNode(_v)
					H[i] = parentNode.right
				}
			}
		}
	}

	for k := range H {
		delete(H, k)
	}
}

func (bt *BinaryTree) GenerateTreeFromArrayUsingQueue(treeArray []any) {
	if len(treeArray) == 0 {
		return
	}

	if treeArray[0] == nil {
		panic("root can't be nil")
	}

	idx := 0

	rootValue, ok := treeArray[idx].(int)
	if !ok {
		panic("root can't be nil")
	}

	idx++

	queue := &Queue{}
	rootNode := NewNode(rootValue)
	bt.root = rootNode
	queue.Enqueue(rootNode)

	for !queue.IsEmpty() {
		parentNode := queue.Dequeue()
		if parentNode == nil {
			break
		}

		if idx < len(treeArray) {
			if _leftChildValue, ok := treeArray[idx].(int); ok {
				node := NewNode(_leftChildValue)
				parentNode.left = node
				queue.Enqueue(node)
			}
			idx++
		}

		if idx < len(treeArray) {
			if _rightChildValue, ok := treeArray[idx].(int); ok {
				node := NewNode(_rightChildValue)
				parentNode.right = node
				queue.Enqueue(node)
			}
			idx++
		}
	}
}

func (bt *BinaryTree) generateTreeFromPreOrder(preOrder []int, preOrderIdx *int, inOrderLeft, inOrderRight int, inOrderMap map[int]int) *Node {
	if inOrderLeft > inOrderRight {
		return nil
	}

	rootVal := preOrder[*preOrderIdx]
	*preOrderIdx++

	node := NewNode(rootVal)

	inOrderRootIdx := inOrderMap[rootVal]

	node.left = bt.generateTreeFromPreOrder(preOrder, preOrderIdx, inOrderLeft, inOrderRootIdx-1, inOrderMap)

	node.right = bt.generateTreeFromPreOrder(preOrder, preOrderIdx, inOrderRootIdx+1, inOrderRight, inOrderMap)

	return node
}

func (bt *BinaryTree) GenerateTreeFromPreOrder(preOrder, inOrder []int) {
	if len(preOrder) == 0 || len(preOrder) != len(inOrder) {
		bt.root = nil
		return
	}

	inOrderMap := make(map[int]int, len(inOrder))
	for i, val := range inOrder {
		inOrderMap[val] = i
	}

	preOrderIdx := 0
	bt.root = bt.generateTreeFromPreOrder(preOrder, &preOrderIdx, 0, len(inOrder)-1, inOrderMap)
}

func RunBinaryTreesADT(run bool) {
	if !run {
		return
	}

	bt1 := NewBinaryTree()
	treeArray1 := []any{2, 4, 6, 8, 3, 9, 5}
	// bt1.GenerateTreeFromArray(treeArray1)
	bt1.GenerateTreeFromArrayUsingQueue(treeArray1)
	bt1.PostOrderTraversal()
	bt1.PreOrderTraversal()
	// bt1.IterativePreOrderTraversal()
	bt1.InOrderTraversal()
	bt1.LevelOrderTraversal()

	// bt2 := NewBinaryTree()
	// treeArray2 := []any{5, 8, 6, nil, 9, 3, nil, nil, nil, 4, 2}
	// // bt2.GenerateTreeFromArray(treeArray2)
	// bt2.GenerateTreeFromArrayUsingQueue(treeArray2)
	// bt2.PreOrderTraversal()
	// bt2.IterativePreOrderTraversal()
	// bt2.InOrderTraversal()
	// bt2.LevelOrderTraversal()

	fmt.Println()

	bt3 := NewBinaryTree()
	preOrder := []int{2, 4, 8, 3, 6, 9, 5}
	inOrder := []int{8, 4, 3, 2, 9, 6, 5}
	bt3.GenerateTreeFromPreOrder(preOrder, inOrder)
	bt3.PostOrderTraversal()
	bt3.PreOrderTraversal()
	bt3.InOrderTraversal()
	bt3.LevelOrderTraversal()

	fmt.Println()
}
