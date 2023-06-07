// union
// Instructions

// Write a program that takes two string and displays, without doubles, the characters that appear in either one of the string.

// The display will be in the same order that the characters appear on the command line and will be followed by a newline ('\n').

// If the number of arguments is different from 2, then the program displays a newline ('\n').
// Usage

// $ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
// zpadintoqefwjy$
// $
// $ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
// df6vewg4thras$
// $
// $ go run . rien "cette phrase ne cache rien" | cat -e
// rienct phas$
// $
// $ go run . | cat -e
// $
// $ go run . rien | cat -e
// $

// 0	"os.Args"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"--allow-builti

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}
	res := ""
	s := args[0] + args[1]
	for _, w := range s {
		if !Choice(res, w) {
			res = res + string(w)
		}
	}

	for _, v := range res {
		z01.PrintRune(v)
	}

	z01.PrintRune('\n')
}

func Choice(res string, w rune) bool {
	for _, n := range res {
		if n == w {
			return true
		}
	}
	return false
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 	} else {
// 		s := args[0] + args[1]
// 		result := ""
// 		for _, v := range s {
// 			if !contain(result, v) {
// 				result = result + string(v)
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
// package main

// import (
// 	"os"

// 	"github.com/01-edu/z01"
// )

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 	} else {
// 		s1 := args[0]
// 		s2 := args[1]
// 		result := ""
// 		for _, w := range s1 {
// 			if !contain(result, w) {
// 				result += string(w)
// 			}
// 		}
// 		for _, w := range s2 {
// 			if !contain(result, w) {
// 				result += string(w)
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func contain(s string, r rune) bool {
// 	for _, w := range s {
// 		if w == r {
// 			return true
// 		}
// 	}
// 	return false
// }

// func contain(result string, v rune) bool {
// 	for _, w := range result {
// 		if w == v {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 2 {
// 		s := args[0] + args[1]
// 		result := ""
// 		for _, w := range s {
// 			found := false
// 			for _, v := range result {
// 				if w == v {
// 					found = true
// 					break
// 				}
// 			}
// 			if !found {
// 				result = result + string(w)
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
