// inter
// Instructions

// Write a program that takes two string and displays, without doubles, the characters that appear in both string, in the order they appear in the first one.

//     The display will be followed by a newline ('\n').

//     If the number of arguments is different from 2, the program displays nothing.

// Usage

// $ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// padinto
// $ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
// df6ewg4
// $

// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.Args"
// 2	"--allow-builtin"

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	} else {
		res := ""
		for _, w := range args[0] {
			for _, v := range args[1] {
				if w == v {
					if !Choice(res, w) && w != ' ' {
						res = res + string(w)
					}
				}
			}
		}

		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

func Choice(res string, w rune) bool {
	for _, z := range res {
		if z == w {
			return true
		}
	}
	return false
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	} else {
// 		result := ""
// 		for _, w := range args[0] {
// 			for _, l := range args[1] {
// 				if w == l {
// 					if !contain(result, w) && w != ' ' {
// 						result = result + string(w)
// 					}
// 				}
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func contain(result string, w rune) bool {
// 	for _, v := range result {
// 		if v == w {
// 			return true
// 		}
// 	}
// 	return false
// }
