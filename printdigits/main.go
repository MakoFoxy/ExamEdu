// printdigits
// Instructions

// Write a program that prints the decimal digits in ascending order (from 0 to 9) on a single line.

// A line is a sequence of characters preceding the end of line character ('\n').
// Usage

// student@ubuntu:~/piscine-go/printdigits$ go build
// student@ubuntu:~/piscine-go/printdigits$ ./printdigits
// 0123456789
// student@ubuntu:~/piscine-go/printdigits$
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	for i := '0'; i <= '9'; i++ {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	var a []int
// 	for i := 0; i <= 9; i++ {
// 		a = append(a, i)
// 		// fmt.Println(a)
// 	}
// 	fmt.Println(a)
// 	for i, v := range a {
// 		fmt.Println(i, v)
// 	}
// 	fmt.Println('\n')

// 	for x := len(a) - 1; x > 0; x-- {
// 		fmt.Println(x)
// 	}
// }

// func main() {
// 	for i := '0'; i <= '9'; i++ {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }

// a := []int{1, 2, 3}
// var numbers = []int {5, 1, 9, 8, 4}

// func main() {
// 	var a []int

// 	for i := 0; i <= 9; i++ {
// 		a = append(a, i)
// 	}
// 	fmt.Println(a)
// 	for _, n := range a {
// 		fmt.Print(n)
// 	}
// }

// func main() {
// 	var a []string

// 	for i := '0'; i <= '9'; i++ {
// 		// z01.PrintRune(i)
// 		a = append(a, string(i))
// 	}

// 	// z01.PrintRune('\n')
// 	fmt.Println(a)
// }
