package recursion

import "fmt"

func power1(x, n float64) float64 {
	if x == 0.0 {
		return 0.0
	} else if n == 0.0 {
		return 1.0
	} else {
		p := power1(x, n-1.0)
		return p * x
	}
}

func factorial1(n float64) float64 {
	if n == 0.0 {
		return 1.0
	} else {
		f := factorial1(n - 1.0)
		return f * n
	}
}

func taylorSeries(x, n float64) float64 {
	if n == 0 {
		return 1
	} else {
		r := taylorSeries(x, n-1.0)
		p := power1(x, n)
		f := factorial1(n)

		return r + (p / f)
	}
}

var p *float64
var f *float64

func init() {
	p = new(float64)
	f = new(float64)
	*p = 1.0
	*f = 1.0
}

func taylorSeries1(x, n float64) float64 {
	if n == 0 {
		return 1
	} else {
		r := taylorSeries1(x, n-1.0)
		_p := *p
		*p = _p * x
		power := *p
		_f := *f
		*f = _f * n
		factorial := *f

		return r + (power / factorial)
	}
}

func TaylorSeries() {
	fmt.Println("Taylor Series in Go!")

	result := taylorSeries(2.0, 5.0)
	fmt.Printf("Taylor series result for x=2 and n=5: %f\n", result)
	result = taylorSeries(3.0, 4.0)
	fmt.Printf("Taylor series result for x=3 and n=4: %f\n", result)

	result = taylorSeries1(2.0, 5.0)
	fmt.Printf("Taylor series result for x=2 and n=5 using iterative approach: %f\n", result)
	*p = 1.0
	*f = 1.0
	result = taylorSeries1(3.0, 4.0)
	fmt.Printf("Taylor series result for x=3 and n=4 using iterative approach: %f\n", result)
}
