package array

import "fmt"

func TwoDArray(run bool) {
	if !run {
		return
	}

	fmt.Println("2D array")

	var A [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	for _, v := range A {
		for _, v1 := range v {
			fmt.Printf("%d,", v1)
		}
	}
	fmt.Println()
}
