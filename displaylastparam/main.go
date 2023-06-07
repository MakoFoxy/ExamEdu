// displaylastparam
// Instructions
// Write a program that displays its last argument, if there is one.

// Usage
// $ go run . hello there
// there
// $ go run . "hello there" how are you
// you
// $ go run . "hello there"
// hello there
// $ go run .
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.*"
// 2	"--allow-builtin"

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		return
	}
	arg := []rune(args[len(args)-1])

	for _, w := range arg {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	args := os.Args
// 	arg := []rune(args[len(args)-1])
// 	for _, v := range arg {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args
// 	arg := []rune(args[len(args)-1])
// 	for _, w := range arg {
// 		if w != 0 {
// 			z01.PrintRune(w)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args

// 	arg := []rune(args[len(args)-1])

// 	for _, v := range arg {
// 		if v != 0 {
// 			z01.PrintRune(v)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
