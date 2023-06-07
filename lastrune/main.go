// lastrune
// Instructions
// Write a function that returns the last rune of a string.

// Expected function
// func LastRune(s string) rune {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"github.com/01-edu/z01"

// 	"piscine"
// )

// func main() {
// 	z01.PrintRune(piscine.LastRune("Hello!"))
// 	z01.PrintRune(piscine.LastRune("Salut!"))
// 	z01.PrintRune(piscine.LastRune("Ola!"))
// 	z01.PrintRune('\n')
// }
// And its output :

// $ go run .
// !!!
// $

// allowedFunctions
// 0	"--allow-builtin"
package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}

func LastRune(s string) rune {
	arr := []rune(s)
	// var arr []rune
	// for _, w := range s {
	// 	arr = append(arr, w)
	// }

	return arr[len(arr)-1]
}

// func LastRune(s string) rune {
// 	a := []rune(s)
// 	return a[len(a)-1]
// }

// func LastRune(s string) rune {
// 	var arr []rune
// 	for _, v := range s {
// 		arr = append(arr, v)
// 	}
// 	return arr[len(arr)-1]
// }
