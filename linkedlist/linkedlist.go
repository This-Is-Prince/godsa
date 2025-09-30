package linkedlist

import "fmt"

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("LinkedList in Go!")

	RunLinkedListADT(true)
}
