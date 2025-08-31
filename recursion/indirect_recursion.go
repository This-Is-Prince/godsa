package recursion

import "fmt"

func a(n int) {
	if n > 0 {
		fmt.Printf("a=%d ", n)
		b(n - 1)
	}
}

func b(n int) {
	if n > 0 {
		fmt.Printf("b=%d ", n)
		a(n - 1)
	}
}

func IndirectRecursion(run bool) {
	if !run {
		return
	}

	fmt.Println("Indirect Recursion in Go!")
	a(3)
}
