package binarytrees

type Queue []*Node

func (q *Queue) Enqueue(node *Node) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() *Node {
	if len(*q) == 0 {
		return nil
	}

	node := (*q)[0]

	*q = (*q)[1:]
	return node
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
