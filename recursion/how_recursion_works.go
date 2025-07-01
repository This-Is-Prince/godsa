package recursion

import "fmt"

func printFirstAndMoveForward(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		printFirstAndMoveForward(n - 1)
	}
}

func moveForwardFirstAndPrint(n int) {
	if n > 0 {
		moveForwardFirstAndPrint(n - 1)
		fmt.Print(n, " ")
	}
}

func HowRecursionWorks() {
	fmt.Println("How Recursion Works in Go!")
	fmt.Println("Print First and Move Forward:")
	printFirstAndMoveForward(5)
	fmt.Println("\nMove Forward First and Print:")
	moveForwardFirstAndPrint(5)
	fmt.Println()
}
