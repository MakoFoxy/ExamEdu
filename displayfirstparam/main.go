// displayfirstparam
// Instructions
// Write a program that displays its first argument, if there is one.

// Usage
// $ go run . hello there
// hello
// $ go run . "hello there" how are you
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
	args := os.Args

	if len(args) < 2 {
		return
	}

	for _, w := range args[1] {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	args := os.Args

// 	for _, v := range args[1] {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	arg := os.Args
// 	for _, v := range arg[2] {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }
