package array

import "fmt"

type ArrayADT struct {
	A []int64
}

func CreateArray(size uint) *ArrayADT {
	arrADT := new(ArrayADT)

	arrADT.A = make([]int64, 0, size)

	return arrADT
}

func (a *ArrayADT) Length() uint {
	return uint(len(a.A))
}

func (a *ArrayADT) Size() uint {
	return uint(cap(a.A))
}

func (a *ArrayADT) Display() {
	for idx, v := range a.A {
		if uint(idx) == a.Length()-1 {
			fmt.Printf("%d\n", v)
		} else {
			fmt.Printf("%d, ", v)
		}
	}
}

func (a *ArrayADT) Append(elms ...int64) {
	a.A = append(a.A, elms...)
}

func (a *ArrayADT) FillArray(raw []int64) {
	a.Append(raw...)
}

func (a *ArrayADT) Get(index uint) (int64, bool) {
	if index >= a.Length() {
		return 0, false
	}

	return a.A[index], true
}

func (a *ArrayADT) Insert(index uint, elm int64) {
	if index > a.Length() {
		return
	}

	if index == a.Length() {
		a.Append(elm)
		return
	}

	a.A = append(a.A[:index+1], a.A[index:]...)
	a.A[index] = elm
}

func (a *ArrayADT) Delete(index uint) {
	if index >= a.Length() {
		return
	}

	a.A = append(a.A[:index], a.A[index+1:]...)
}

func RunArrayADT(run bool) {
	if !run {
		return
	}

	arr := CreateArray(10)

	arr.FillArray([]int64{2, 3, 4, 5, 6})

	arr.Display()

	arr.Append(7, 8, 9, 10, 11, 12, 13, 14, 15)

	arr.Display()

	arr.Insert(0, 100)
	arr.Insert(15, 120)
	arr.Insert(10, 130)

	arr.Display()

	arr.Delete(17)
	arr.Delete(16)
	arr.Delete(10)
	arr.Delete(0)

	arr.Display()
}
