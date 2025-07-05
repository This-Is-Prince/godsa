package recursion

import "fmt"

func nestedRecursion(n int) int {
	if n > 100 {
		return n - 10
	} else {
		return nestedRecursion(nestedRecursion(n + 11))
	}
}

func NestedRecursion() {
	fmt.Println("Nested Recursion in Go!")
	result := nestedRecursion(95)
	fmt.Printf("Result of nested recursion: %d\n", result)
}
