package recursion

import "fmt"

func sumOfNaturalNumbers(n int) int {
	if n == 0 {
		return 0
	} else {
		return n + sumOfNaturalNumbers(n-1)
	}
}

func iSumOfNaturalNumbers(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s
}

func SumOfNaturalNumbers(run bool) {
	if !run {
		return
	}

	fmt.Println("Sum of Natural Numbers in Go!")

	sum := sumOfNaturalNumbers(5)
	fmt.Printf("Sum of first 5 natural numbers: %d\n", sum)

	// Iterative approach
	iterativeSum := iSumOfNaturalNumbers(5)
	fmt.Printf("Iterative sum of first 5 natural numbers: %d\n", iterativeSum)
}
