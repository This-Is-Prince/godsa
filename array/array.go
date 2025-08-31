package array

import "fmt"

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("Array in Go!")

	Declarations(false)
	TwoDArray(false)
	RunArrayADT(true)
}
