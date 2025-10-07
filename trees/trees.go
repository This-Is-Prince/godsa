package trees

import (
	"fmt"

	binarysearchtrees "github.com/This-Is-Prince/godsa/trees/binary_search_trees"
	binarytrees "github.com/This-Is-Prince/godsa/trees/binary_trees"
)

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("Trees in Go!")
	binarytrees.RunBinaryTreesADT(true)
	binarysearchtrees.RunBinarySearchTrees(false)
}
