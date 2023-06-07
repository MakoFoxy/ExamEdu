// reduceint
// Instructions
// The function should have as parameters a slice of integers a []int and a function f func(int, int) int. For each element of the slice, it should apply the function f func(int, int) int, save the result and then print it.

// Expected function
// func ReduceInt(a []int, f func(int, int) int) {

// }
// Usage
// Here is a possible program to test your function :

// package main

// func main() {
// 	mul := func(acc int, cur int) int {
// 		return acc * cur
// 	}
// 	sum := func(acc int, cur int) int {
// 		return acc + cur
// 	}
// 	div := func(acc int, cur int) int {
// 		return acc / cur
// 	}
// 	as := []int{500, 2}
// 	ReduceInt(as, mul)
// 	ReduceInt(as, sum)
// 	ReduceInt(as, div)
// }
// And its output :

// $ go run .
// 1000
// 502
// 250
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"strconv.Itoa"
// 2	"--allow-builtin"
package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2} // len = 4
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

func ReduceInt(a []int, f func(int, int) int) {
	z := a[0]

	for i := 1; i < len(a); i++ {
		z = f(z, a[i])
	}

	res := strconv.Itoa(z)

	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

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

// func Mul(a int, b int) int {
// 	return a * b
// }

// func Sum(a int, b int) int {
// 	return a + b
// }

// func Div(a int, b int) int {
// 	return a / b
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	z := a[0]

// 	for i := 1; i < len(a); i++ {
// 		z = f(z, a[i])
// 	}
// 	res := strconv.Itoa(z)

// 	for _, w := range res {
// 		z01.PrintRune(w)
// 	}

// 	z01.PrintRune('\n')
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	z := a[0]
// 	for i := 1; i < len(a); i++ {
// 		z = f(z, a[i])
// 	}
// 	res := strconv.Itoa(z)

// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func ReduceInt(a []int, f func(int, int) int) { //
// 	z := a[0]

// 	for i := 1; i < len(a); i++ { // i = 3
// 		z = f(z, a[i])
// 	}
// 	res := strconv.Itoa(z)

// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}

// 	z01.PrintRune('\n')

// 	// z := a[0] // хранит результат

// 	// for i := 1; i < len(a); i++ { // длина переменной as
// 	// 	z = f(z, a[i]) // передает в z функцию f acc*cur
// 	// }
// 	// res := strconv.Itoa(z)

// 	// for _, w := range res {
// 	// 	z01.PrintRune(w)
// 	// }
// 	// z01.PrintRune('\n')

// 	// var ab []int
// 	// z := a[0]
// 	// var arr []rune
// 	// for i := 1; i < len(a); i++ {
// 	// 	z = f(z, a[i])
// 	// }
// 	// res := strconv.Itoa(z)

// 	// for _, w := range res {
// 	// 	arr = append(arr, rune(w))
// 	// 	z01.PrintRune(rune(w))
// 	// }
// 	// z01.PrintRune('\n')

// 	// for _, d := range res {
// 	// 	ab = append(ab, int(d))
// 	// 	z01.PrintRune(d)
// 	// }
// 	// z01.PrintRune('\n')

// 	// for _, c := range ab {
// 	// 	z01.PrintRune(rune(c))
// 	// }
// 	// z01.PrintRune('\n')
// 	// for _, x := range ab {
// 	// 	z01.PrintRune(rune(x))
// 	// }
// 	// z01.PrintRune('\n')

// 	// for _, x := range arr {
// 	// 	z01.PrintRune(x)
// 	// }
// 	// z01.PrintRune('\n')
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	z := a[0] // объявили  массив для " 1 цикла а"
// 	for i := 0; i < len(a); i++ {
// 		z = f(z, a[i]) // передали в массив зед  в "1 цикл а"
// 	}
// 	res := strconv.Itoa(z) // передали  численный массив в функц Итоа
// 	for _, w := range res {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// func ReduceInt(a []int, f func(int, int) int) {
// 	result := a[0]
// 	for i := 1; i < len(a); i++ {
// 		result = f(result, a[i])
// 	}
// 	res := strconv.Itoa(result)
// 	for _, w := range res {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }
