package recursion

import "fmt"

func sumOfNaturalNumbers(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sumOfNaturalNumbers(n-1)
	}
}

func SumOfNaturalNumbers() {
	fmt.Println("Sum of Natural Numbers in Go!")

	sum := sumOfNaturalNumbers(5)
	fmt.Printf("Sum of first 5 natural numbers: %d\n", sum)
}
