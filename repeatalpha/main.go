// repeatalpha
// Instructions

// Write a program called repeat_alpha that takes a string and displays it repeating each alphabetical character as many times as its alphabetical index.

// The result must be followed by a newline ('\n').

// 'a' becomes 'a', 'b' becomes 'bb', 'e' becomes 'eeeee', etc...

// If the number of arguments is different from 1, the program displays nothing.
// Usage

// $ go run . abc | cat -e
// abbccc
// $ go run . Choumi. | cat -e
// CCChhhhhhhhooooooooooooooouuuuuuuuuuuuuuuuuuuuummmmmmmmmmmmmiiiiiiiii.$
// $ go run . "abacadaba 01!" | cat -e
// abbacccaddddabba 01!$
// $ go run .
// $ go run . "" | cat -e
// $
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
	if len(args) != 1 {
		return
	} else {
		res := ""
		for _, w := range args[0] {
			for i := 0; i < Times(w); i++ {
				res = res + string(w)
				z01.PrintRune(w)
			}
		}
		if res == "" {
			return
		}
		z01.PrintRune('\n')
	}
}

func Times(w rune) int {
	if w >= 'a' && w <= 'z' {
		return int(w-'a') + 1
	} else if w >= 'A' && w <= 'Z' {
		return int(w-'A') + 1
	} else {
		return 1
	}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for _, x := range args[0] {
// 			for z := 0; z < times(x); z++ {
// 				z01.PrintRune(x)
// 			}
// 		}
// 	}

// 	z01.PrintRune('\n')
// }

// func Times(r rune) int {
// 	if r >= 'a' && r <= 'z' {
// 		return int(r-'a') + 1
// 	} else if r >= 'A' && r <= 'Z' {
// 		return int(r-'A') + 1
// 	} else {
// 		return 1
// 	}
// }
