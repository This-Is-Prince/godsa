package recursion

import "fmt"

func treeRecursion(n int) {
	if n > 0 {
		fmt.Printf("%d ", n)
		treeRecursion(n - 1)
		treeRecursion(n - 1)
	}
}

func TreeRecursion(run bool) {
	if !run {
		return
	}

	fmt.Println("How Tree Recursion Works in Go!")

	treeRecursion(3)
}
