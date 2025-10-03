package binarytrees

type Stack []*Node

func (s *Stack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *Stack) Pop() *Node {
	l := len(*s)
	if l == 0 {
		return nil
	}

	node := (*s)[l-1]

	(*s)[l-1] = nil

	*s = (*s)[:l-1]

	return node
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
