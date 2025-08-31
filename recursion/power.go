package recursion

import "fmt"

func power(m int, n int) int {
	if m == 0 {
		return 0
	} else if n == 0 {
		return 1
	} else {
		return power(m, n-1) * m
	}
}

func Power(run bool) {
	if !run {
		return
	}

	fmt.Println("Power in Go!")

	result := power(5, 3)
	fmt.Printf("5 raised to the power of 3 is: %d\n", result)
}
