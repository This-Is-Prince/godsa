package stringadt

import (
	"bytes"
	"fmt"
)

type String struct {
	S []rune
}

func CreateString(size uint) *String {
	str := new(String)

	str.S = make([]rune, 0, size)

	return str
}

func (str *String) Display() {
	if str == nil || str.S == nil {
		return
	}

	fmt.Println(string(str.S))
}

func (str *String) Length() int {
	if str == nil || str.S == nil {
		return 0
	}

	return len(str.S)
}

func (str *String) Append(s string) {
	if str == nil || str.S == nil {
		return
	}

	str.S = append(str.S, []rune(s)...)
}

func (str *String) LowerCase() {
	if str == nil || str.S == nil {
		return
	}

	for i, v := range str.S {
		if v >= 65 && v <= 90 {
			str.S[i] = 97 + (v - 65)
		}
	}
}

func (str *String) UpperCase() {
	if str == nil || str.S == nil {
		return
	}

	for i, v := range str.S {
		if v >= 97 && v <= 122 {
			str.S[i] = 65 + (v - 97)
		}
	}
}

func IsVowel(c rune) bool {
	vowels := []byte("aeiouAEIOU")

	return bytes.ContainsRune(vowels, c)
}

func (str *String) VowelCount() int {
	if str == nil || str.S == nil {
		return 0
	}

	count := 0

	for _, v := range str.S {
		if IsVowel(v) {
			count++
		}
	}

	return count
}

func (str *String) WordCount() int {
	if str == nil || str.S == nil || str.Length() == 0 {
		return 0
	}

	count := 0

	for i, v := range str.S {
		if v == ' ' && i > 0 && str.S[i-1] != v {
			count++
		}
	}

	l := str.Length()
	if l > 0 && str.S[l-1] != ' ' {
		count++
	}

	return count
}

func (str *String) IsValidString() bool {
	if str == nil || str.S == nil {
		return false
	}

	if str.Length() == 0 {
		return true
	}

	for _, v := range str.S {
		if !((v >= 65 && v <= 90) || (v >= 97 && v <= 122) || (v >= 48 && v <= 57)) {
			return false
		}
	}

	return false
}

func (str *String) Reverse() {
	if str == nil || str.S == nil || str.Length() == 0 {
		return
	}

	i, j := 0, str.Length()-1

	for i < j {
		tmp := str.S[i]
		str.S[i] = str.S[j]
		str.S[j] = tmp
		i++
		j--
	}
}

func (str *String) IsEqual(o *String) bool {
	if str == nil || str.S == nil || o == nil || o.S == nil {
		return false
	}

	i, j := 0, 0

	for i < str.Length() && j < o.Length() {
		if str.S[i] != o.S[j] {
			return false
		}

		i++
		j++
	}

	return i == str.Length() && j == o.Length()
}

func (str *String) IsPalindrome() bool {
	if str == nil || str.S == nil {
		return false
	}

	if str.Length() == 0 {
		return true
	}

	i, j := 0, str.Length()-1

	for i < j {
		if str.S[i] != str.S[j] {
			return false
		}

		i++
		j++
	}

	return true
}

func (str *String) IsAllUpperCaseAlpha() bool {
	if str == nil || str.S == nil {
		return false
	}

	if str.Length() == 0 {
		return true
	}

	for _, v := range str.S {
		if !(v >= 65 && v <= 90) {
			return false
		}
	}

	return true
}

func (str *String) IsAllLowerCaseAlpha() bool {
	if str == nil || str.S == nil {
		return false
	}

	if str.Length() == 0 {
		return true
	}

	for _, v := range str.S {
		if !(v >= 97 && v <= 122) {
			return false
		}
	}

	return true
}

func (str *String) IsAllNumeric() bool {
	if str == nil || str.S == nil {
		return false
	}

	if str.Length() == 0 {
		return true
	}

	for _, v := range str.S {
		if !(v >= 48 && v <= 57) {
			return false
		}
	}

	return true
}

func (str *String) DuplicatesCharacters() []rune {
	if str == nil || str.S == nil {
		return nil
	}

	arr := make([]rune, 0, str.Length())

	if str.Length() == 0 {
		return arr
	}

	H := int64(0)
	X := int64(1)

	for _, v := range str.S {
		X = int64(1)
		X = X << (v - 97)

		if (H & X) > 0 {
			arr = append(arr, v)
		} else {
			H = H | X
		}
	}

	return arr
}

func (str *String) Duplicates() map[rune]int {
	if str == nil || str.S == nil {
		return nil
	}

	m := map[rune]int{}

	if str.Length() == 0 {
		return m
	}

	for _, v := range str.S {
		m[v]++
	}

	return m
}

func (str *String) IsAnagram(o *String) bool {
	if str == nil || str.S == nil || o == nil || o.S == nil {
		return false
	}

	if str.Length() != o.Length() {
		return false
	}

	H := [26]rune{}

	for _, v := range str.S {
		H[(v-97)]++
	}

	for _, v := range o.S {
		H[(v-97)]--
	}

	for _, v := range H {
		if v > 0 || v < 0 {
			return false
		}
	}

	return true
}

func RunStringADT(run bool) {
	if !run {
		return
	}

	str := CreateString(10)
	str.Append("   I     Love     Golang    ")
	str.Display()

	str.LowerCase()
	str.Display()
	str.UpperCase()
	str.Display()

	fmt.Println("Vowel Count is: ", str.VowelCount())
	fmt.Println("Word Count is: ", str.WordCount())

	str.Reverse()
	str.Display()
}

func Start(run bool) {
	if !run {
		return
	}

	fmt.Println("String in Go!")

	RunStringADT(true)
}
