package main

import (
	"fmt"

	"github.com/This-Is-Prince/godsa/array"
	"github.com/This-Is-Prince/godsa/linkedlist"
	"github.com/This-Is-Prince/godsa/recursion"
	"github.com/This-Is-Prince/godsa/stringadt"
	"github.com/This-Is-Prince/godsa/trees"
)

func main() {
	fmt.Println("Data Structures and Algorithms in Go!")

	recursion.Start(false)
	array.Start(false)
	stringadt.Start(false)
	linkedlist.Start(false)
	trees.Start(true)
}
