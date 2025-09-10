package array

import (
	"fmt"
	"slices"
)

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

func (a *ArrayADT) LinearSearch(elm int64) int {
	for idx, v := range a.A {
		if v == elm {
			return idx
		}
	}

	return -1
}

func (a *ArrayADT) BinarySearch(elm int64) int {
	l, h := 0, int(a.Length()-1)

	for l <= h {
		m := (l + h) / 2

		if a.A[m] == elm {
			return m
		} else if a.A[m] > elm {
			h = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func (a *ArrayADT) Get(index uint) (int64, bool) {
	if index >= a.Length() {
		return 0, false
	}

	return a.A[index], true
}

func (a *ArrayADT) First() (int64, bool) {
	if a.Length() == 0 {
		return 0, false
	}

	return a.A[0], true
}

func (a *ArrayADT) Last() (int64, bool) {
	l := a.Length()

	if l == 0 {
		return 0, false
	}

	return a.A[l-1], true
}

func (a *ArrayADT) Set(index uint, value int64) {
	if index < a.Length() {
		a.A[index] = value
	}
}

func (a *ArrayADT) Max() (int64, bool) {
	var m int64

	if a.Length() == 0 {
		return m, false
	}

	m = a.A[0]

	for _, v := range a.A {
		if m < v {
			m = v
		}
	}

	return m, true
}

func (a *ArrayADT) Min() (int64, bool) {
	var m int64

	if a.Length() == 0 {
		return m, false
	}

	m = a.A[0]

	for _, v := range a.A {
		if m > v {
			m = v
		}
	}

	return m, true
}

func (a *ArrayADT) Sum() int64 {
	var s int64

	for _, v := range a.A {
		s += v
	}

	return s
}

func (a *ArrayADT) Avg() float64 {
	s := a.Sum()

	return float64(s) / float64(a.Length())
}

func (a *ArrayADT) Reverse() {
	var i, j uint = 0, a.Length() - 1
	var tmp int64

	for i < j {
		tmp = a.A[i]
		a.A[i] = a.A[j]
		a.A[j] = tmp
		i++
		j--
	}
}

func (a *ArrayADT) IsSorted() bool {
	for i, v := range a.A {
		if i < int(a.Length())-1 && v > a.A[i+1] {
			return false
		}
	}

	return true
}

func (a *ArrayADT) InsertedInSortedArray(elm int64) bool {
	if !a.IsSorted() {
		return false
	}

	if a.Length() == 0 {
		a.Append(elm)
		return true
	}

	idx := a.Length() - 1

	if a.A[idx] <= elm {
		a.Append(elm)
		return true
	} else {
		a.Append(a.A[idx])
	}

	for idx > 0 && a.A[idx-1] > elm {
		a.A[idx] = a.A[idx-1]
		idx--
	}

	a.A[idx] = elm
	return true
}

func (a *ArrayADT) NegativeOnLeftSide() {
	var i, j uint = 0, a.Length() - 1
	var tmp int64

	for i < j {
		if a.A[i] < 0 {
			i++
		} else if a.A[j] > 0 {
			j--
		} else {
			tmp = a.A[i]
			a.A[i] = a.A[j]
			a.A[j] = tmp
			i++
			j--
		}
	}
}

func (a *ArrayADT) LeftShift() {
	if a.Length() == 0 {
		return
	}

	i := 0

	for i < int(a.Length())-1 {
		a.A[i] = a.A[i+1]
		i++
	}

	a.A[i] = 0
}

func (a *ArrayADT) RightShift() {
	if a.Length() == 0 {
		return
	}

	var i = a.Length() - 1

	a.Append(a.A[i])

	for i > 0 {
		a.A[i] = a.A[i-1]
		i--
	}

	a.A[i] = 0
}

func (a *ArrayADT) Rotate() {
	if a.Length() <= 1 {
		return
	}

	tmp := a.A[0]
	a.LeftShift()
	a.A[a.Length()-1] = tmp
}

func (a *ArrayADT) Merge(b *ArrayADT) *ArrayADT {
	f := CreateArray(a.Size() + b.Size())

	if a.IsSorted() && b.IsSorted() {
		var i, j uint = 0, 0
		for i < a.Length() && j < b.Length() {
			if a.A[i] <= b.A[j] {
				f.Append(a.A[i])
				i++
			} else if a.A[i] > b.A[j] {
				f.Append(b.A[j])
				j++
			}
		}

		for i < a.Length() {
			f.Append(a.A[i])
			j++
		}

		for j < b.Length() {
			f.Append(b.A[j])
			j++
		}
	} else {
		f.Append(a.A...)
		f.Append(b.A...)
	}

	return f
}

func (a *ArrayADT) Union(b *ArrayADT) *ArrayADT {
	u := CreateArray(a.Size() + b.Size())

	if a.Length() == 0 {
		u.Append(b.A...)
	} else if b.Length() == 0 {
		u.Append(a.A...)
	} else if a.IsSorted() && b.IsSorted() {
		var i, j uint = 0, 0

		if a.A[i] <= b.A[j] {
			u.Append(a.A[i])
			i++
		} else {
			u.Append(b.A[j])
			j++
		}

		for i < a.Length() && j < b.Length() {
			if a.A[i] <= b.A[j] {
				if a.A[i] != u.A[u.Length()-1] {
					u.Append(a.A[i])
				}
				i++
			} else {
				if b.A[j] != u.A[u.Length()-1] {
					u.Append(b.A[j])
				}
				j++
			}
		}

		for i < a.Length() {
			if a.A[i] != u.A[u.Length()-1] {
				u.Append(a.A[i])
			}
			i++
		}

		for j < b.Length() {
			if b.A[j] != u.A[u.Length()-1] {
				u.Append(b.A[j])
			}
			j++
		}
	} else {
		isPresent := false

		for _, v := range a.A {
			isPresent = false
			for _, _v := range u.A {
				if v == _v {
					isPresent = true
				}
			}

			if !isPresent {
				u.Append(v)
			}
		}

		for _, v := range b.A {
			isPresent = false
			for _, _v := range u.A {
				if v == _v {
					isPresent = true
				}
			}

			if !isPresent {
				u.Append(v)
			}
		}
	}

	return u
}

func (a *ArrayADT) Intersection(b *ArrayADT) *ArrayADT {
	u := CreateArray(a.Size() + b.Size())

	if a.Length() == 0 {
		u.Append(b.A...)
	} else if b.Length() == 0 {
		u.Append(a.A...)
	} else if a.IsSorted() && b.IsSorted() {
		var i, j uint = 0, 0

		for i < a.Length() && j < b.Length() {
			if a.A[i] < b.A[j] {
				i++
			} else if a.A[i] > b.A[j] {
				j++
			} else {
				elm, ok := u.Last()

				if ok {
					if elm != a.A[i] {
						u.Append(a.A[i])
					}
				} else {
					u.Append(a.A[i])
				}

				i++
				j++
			}
		}
	} else {
		for _, v := range a.A {
			if slices.Contains(b.A, v) {
				if !slices.Contains(u.A, v) {
					u.Append(v)
				}
			}
		}
	}

	return u
}

func (a *ArrayADT) Difference(b *ArrayADT) *ArrayADT {
	u := CreateArray(a.Size() + b.Size())

	if a.Length() == 0 {
		u.Append(b.A...)
	} else if b.Length() == 0 {
		u.Append(a.A...)
	} else if a.IsSorted() && b.IsSorted() {
		var i, j uint = 0, 0

		for i < a.Length() && j < b.Length() {
			if a.A[i] < b.A[j] {
				elm, ok := u.Last()

				if ok {
					if elm != a.A[i] {
						u.Append(a.A[i])
					}
				} else {
					u.Append(a.A[i])
				}
				i++
			} else if a.A[i] > b.A[j] {
				j++
			} else {
				i++
				j++
			}
		}

		for i < a.Length() {
			elm, ok := u.Last()

			if ok {
				if elm != a.A[i] {
					u.Append(a.A[i])
				}
			} else {
				u.Append(a.A[i])
			}
			i++
		}
	} else {
		for _, v := range a.A {
			if !slices.Contains(b.A, v) && !slices.Contains(u.A, v) {
				u.Append(v)
			}
		}
	}

	return u
}

func (a *ArrayADT) FindingSingleMissingElementInSortedArray() (int64, bool) {
	if !a.IsSorted() || a.Length() < 2 {
		return 0, false
	}

	var i uint = 0
	var curr, next int64

	for i < a.Length()-1 {
		curr = a.A[i]
		next = a.A[i+1]

		if curr != next && curr+1 != next {
			return curr + 1, true
		}

		i++
	}

	return 0, false
}

func (a *ArrayADT) FindingMultipleMissingElementInSortedArray() (n *ArrayADT) {
	if !a.IsSorted() || a.Length() < 2 {
		return
	}

	var i uint = 0
	var curr, next int64

	for i < a.Length()-1 {
		curr = a.A[i]
		next = a.A[i+1]

		if curr != next && curr+1 != next {
			curr++
			for curr < next {
				if n == nil {
					n = CreateArray(a.Length())
				}

				n.Append(curr)
				curr++
			}
		}

		i++
	}

	return
}

func (a *ArrayADT) FindingDuplicatesElementsInSortedArray() (n *ArrayADT) {
	if !a.IsSorted() || a.Length() < 2 {
		return
	}

	var i uint = 0
	var curr, next, lastDuplicate int64

	lastDuplicate = a.A[i] - 1

	for i < a.Length()-1 {
		curr = a.A[i]
		next = a.A[i+1]

		if curr == next {
			if curr != lastDuplicate {
				if n == nil {
					n = CreateArray(a.Length())
				}

				n.Append(curr)
				lastDuplicate = curr
			}
		}

		i++
	}

	return
}

func (a *ArrayADT) FindingDuplicatesElementsInUnsortedArray() (n *ArrayADT) {
	if a.Length() < 2 {
		return
	}

	for i, v := range a.A {
		j := i + 1
		count := 0

		for j < int(a.Length()) {
			if v == a.A[j] && count == 0 {
				if n == nil {
					n = CreateArray(a.Size())
				}
				n.Append(v)
				count++
			}
			j++
		}
	}

	return
}

func (a *ArrayADT) FindingPairOfElementsWithSumKInSortedArray(k int64) (n *ArrayADT) {
	if a.Length() < 2 {
		return
	}

	var i, j uint = 0, a.Length() - 1

	for i < j {
		sum := a.A[i] + a.A[j]
		if sum == k {
			n = CreateArray(a.Length())
			n.Append(a.A[i], a.A[j])
			i++
			j--
		} else if sum < k {
			i++
		} else {
			j--
		}
	}

	return
}

func (a *ArrayADT) FindingPairOfElementsWithSumKInUnSortedArray(k int64) (n *ArrayADT) {
	if a.Length() < 2 {
		return
	}

	for i := range a.A {
		j := i + 1
		for j < int(a.Length()) {
			if a.A[i]+a.A[j] == k {
				n = CreateArray(2)
				n.Append(a.A[i], a.A[j])
				break
			}
			j++
		}
	}

	return
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

	fmt.Printf("\n\n")

	fmt.Println("7 is at index:", arr.LinearSearch(7))
	fmt.Println("1 is at index:", arr.LinearSearch(1))
	fmt.Println("100 is at index:", arr.LinearSearch(100))
	fmt.Println("15 is at index:", arr.LinearSearch(15))
	fmt.Println("7 is at index:", arr.BinarySearch(7))
	fmt.Println("1 is at index:", arr.BinarySearch(1))
	fmt.Println("100 is at index:", arr.BinarySearch(100))
	fmt.Println("15 is at index:", arr.BinarySearch(15))

	fmt.Printf("\n\n")

	elmAt2, ok := arr.Get(2)
	fmt.Println("Elm at index 2 is:", elmAt2, ok)
	max, ok := arr.Max()
	fmt.Println("Max:", max, ok)
	min, ok := arr.Min()
	fmt.Println("Min:", min, ok)
	fmt.Println("Sum:", arr.Sum())
	fmt.Println("Avg:", arr.Avg())
	isSorted := arr.IsSorted()
	fmt.Println("IsSorted:", isSorted)
	arr.Reverse()
	arr.Display()
	isSorted = arr.IsSorted()
	fmt.Println("IsSorted:", isSorted)

	fmt.Printf("\n\n")

	fmt.Println("Length of Array:", arr.Length())
	fmt.Println("Size of Array:", arr.Size())
}
