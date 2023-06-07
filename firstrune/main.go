// firstrune
// Instructions
// Write a function that returns the first rune of a string.

// Expected function
// func FirstRune(s string) rune {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"github.com/01-edu/z01"

// 	"piscine"
// )

// func main() {
// 	z01.PrintRune(piscine.FirstRune("Hello!"))
// 	z01.PrintRune(piscine.FirstRune("Salut!"))
// 	z01.PrintRune(piscine.FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }
// And its output :

// $ go run .
// HSO
// $
// allowedFunctions
// 0	"--allow-builtin"

package main

import (
	"github.com/01-edu/z01"
)

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(s string) rune {
	arr := []rune(s)
	return arr[0]
	// var arr []rune
	// for _, w := range s {
	// 	arr = append(arr, w)
	// }

	// return arr[0]
}

// func FirstRune(s string) rune {
// 	var arr []rune
// 	res := ""
// 	// res := ""
// 	for _, w := range s {
// 		arr = append(arr, w)
// 	}
// 	// return arr[0]
// 	for _, x := range arr {
// 		if arr[0] != x {
// 			res = res + string(arr[x])
// 		}
// 	}
// 	return arr[0]
// }

// func FirstRune(s string) rune {
// 	arr := []rune(s)
// 	return (arr[0])
// }

// func FirstRune(s string) rune {
// 	arr := []rune(s)
// 	return (arr[0])
// }

// func FirstRune(s string) rune {
// 	var result []rune
// 	for _, v := range s {
// 		result = append(result, v)
// 	}
// 	return result[0]
// }
