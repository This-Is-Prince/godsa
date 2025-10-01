package trees

import (
	"fmt"

	binarytrees "github.com/This-Is-Prince/godsa/trees/binary_trees"
)

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("Trees in Go!")
	binarytrees.RunTreesADT(true)
}
