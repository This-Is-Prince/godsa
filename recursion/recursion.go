package recursion

import "fmt"

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("Recursion in Go!")

	HowRecursionWorks(false)
	TreeRecursion(false)
	IndirectRecursion(false)
	NestedRecursion(false)
	SumOfNaturalNumbers(false)
	Factorial(false)
	Power(false)
	TaylorSeries(false)
}
