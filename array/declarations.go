package array

import "fmt"

func Declarations(run bool) {
	if !run {
		return
	}

	fmt.Println("How Array declared in Go!")

	var A [5]int
	for _, a := range A {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()

	var B [5]int = [5]int{1, 2, 3, 4, 5}
	for _, a := range B {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()

	var C = [5]int{1, 2, 3, 4, 5}
	for _, a := range C {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()

	var D = []int{1, 2, 3, 4, 5} // Slices
	for _, a := range D {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()

	E := [5]int{1, 2}
	for _, a := range E {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()

	F := []int{1, 2} // Slices
	for _, a := range F {
		fmt.Printf("%d, ", a)
	}
	fmt.Println()
}
