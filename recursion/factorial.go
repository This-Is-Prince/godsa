package recursion

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return factorial(n-1) * n
	}
}

func Factorial() {
	fmt.Println("Factorial in Go!")

	fact := factorial(5)
	fmt.Printf("Factorial of 5: %d\n", fact)
}
