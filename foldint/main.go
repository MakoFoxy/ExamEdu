// foldint
// Instructions
// The function should have as parameters a function, f func(int, int) int a slice of integers, slice []int and an int acc int. For each element of the slice, it should apply the arithmetic function, save the result and print it. The function will be tested with our own functions Add, Sub, and Mul.

// Expected function
// func FoldInt(f func(int, int) int, a []int, n int) {

// }
// Usage
// Here is a possible program to test your function:

// package main

// func main() {
// 	table := []int{1, 2, 3}
// 	ac := 93
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// 	fmt.Println()

// 	table = []int{0}
// 	FoldInt(Add, table, ac)
// 	FoldInt(Mul, table, ac)
// 	FoldInt(Sub, table, ac)
// }
// And its output :

// $ go run .
// 99
// 558
// 87

// 93
// 0
// 93
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	table := []int{1, 2, 3}
	ac := 93
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
	fmt.Println()

	table = []int{0}
	FoldInt(Add, table, ac)
	FoldInt(Mul, table, ac)
	FoldInt(Sub, table, ac)
}

func FoldInt(f func(int, int) int, a []int, n int) {
	for i := 0; i < len(a); i++ {
		n = f(n, a[i])
	}
	res := Atoi(n)

	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

func Add(a, b int) int {
	return a + b
}

func Mul(a, b int) int {
	return a * b
}

func Sub(a, b int) int {
	return a - b
}

func Atoi(n int) string {
	res := ""

	if n == 0 {
		res = "0"
	} else if n < 0 {
		for n != 0 {
			res = string(rune(-(n%10 + 48))) + res
			n = n / 10
		}
		res = string('-') + res
	} else {
		for n != 0 {
			res = string(rune(n%10+48)) + res
			n = n / 10
		}
	}
	return res
}

// func FoldInt(f func(int, int) int, a []int, n int) {
// 	for i := 0; i < len(a); i++ {
// 		n = f(n, a[i])
// 	}
// 	for _, w := range Itoa(n) {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// func Add(a int, b int) int {
// 	return a + b
// }

// func Mul(a int, b int) int {
// 	return a * b
// }

// func Sub(a int, b int) int {
// 	return a - b
// }

// func Itoa(n int) string {
// 	result := ""
// 	if n == 0 {
// 		result = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			result = string(rune(-(n%10))+48) + result
// 			n = n / 10
// 		}
// 		result = string('-') + result
// 	} else {
// 		for n != 0 {
// 			result = string(rune(n%10)+48) + result
// 			n = n / 10
// 		}
// 	}
// 	return result
// }

// func Itoa(n int) string {
// 	result := ""
// 	if n == 0 {
// 		z01.PrintRune(rune(48))
// 	}
// 	for n != 0 {
// 		result = string(rune(n/10+48)) + result
// 		n = n % 10
// 	}
// 	return result
// }

// func FoldInt(f func(int, int) int, a []int, n int) {
// 	for i := 0; i < len(a); i++ {
// 		n = f(n, a[i])
// 	}

// 	for _, w := range Itao(n) {
// 		z01.PrintRune(w)
// 	}

// 	z01.PrintRune('\n')
// }

// func Add(a int, b int) int {
// 	return a + b
// }

// func Mul(a int, b int) int {
// 	return a * b
// }

// func Sub(a int, b int) int {
// 	return a - b
// }

// func Itao(s int) string {
// 	res := ""
// 	if s == 0 {
// 		z01.PrintRune(rune(48))
// 	}
// 	for s != 0 {
// 		res = string(rune(s%10+48)) + res
// 		s = s / 10
// 	}
// 	return res
// }

// func FoldInt(f func(int, int) int, a []int, n int) {
// 	for i := 0; i < len(a); i++ {
// 		n = f(n, a[i])
// 	}
// 	for _, w := range Itoa(n) {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// func Itoa(n int) string {
// 	result := ""
// 	if n == 0 {
// 		z01.PrintRune('0')
// 	}
// 	for n != 0 {
// 		result = string(rune(n%10)+48) + result
// 		n = n / 10
// 	}
// 	return result
// }

// func Add(a int, b int) int {
// 	return a + b
// }

// func Mul(a int, b int) int {
// 	return a * b
// }

// func Sub(a int, b int) int {
// 	return a - b
// }

// func FoldInt(f func(int, int) int, a []int, n int) {
// 	result := n
// 	for _, v := range a {
// 		result = f(result, v)
// 	}
// 	fmt.Println(result)
// }
