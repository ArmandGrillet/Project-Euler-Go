package algo

import "math"

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Triangle(n int) int {
	return ((n * (n + 1)) / 2)
}

func Pentagonal(n int) int {
	return ((n * (3*n - 1)) / 2)
}

func IsPandigital(number int) bool {
	digits := digits(number)
	if len(digits) == 9 {
		for i := 1; i < 10; i++ {
			if !Contains(digits, i) {
				return false
			}
		}
		return true
	} else {
		return false
	}
}

func EqualSlices(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func Factorial(x int) int {
	if x == 0 {
		return 1
	} else {
		return x * Factorial(x-1)
	}
}

func Digits(number int) []int {
	digits := []int{}
	for number > 0 {
		digits = append(digits, number%10)
		number = number / 10
	}
	return digits
}

func Join(i, j int) int {
	e, tempJ := 0, j
	for tempJ > 0 {
		e++
		tempJ = tempJ / 10
	}
	return i*int(math.Pow10(e)) + j
}
