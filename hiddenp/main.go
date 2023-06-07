// hiddenp
// Instructions

// Write a program named hiddenp that takes two string and that, if the first string is hidden in the second one, displays 1 followed by a newline ('\n'), otherwise it displays 0 followed by a newline.

// Let s1 and s2 be string. It is considered that s1 is hidden in s2 if it is possible to find each character from s1 in s2, in the same order as they appear in s1.

// If s1 is an empty string, it is considered hidden in any string.

// If the number of arguments is different from 2, the program displays nothing.
// Usage

// $ go run . "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6" | cat -e
// 1$
// $ go run . "abc" "2altrb53c.sse" | cat -e
// 1$
// $ go run . "abc" "btarc" | cat -e
// 0$
// $ go run . "DD" "DABC" | cat -e
// 0$
// $ go run .
//$
// 0	"os.Args"
// 1	"github.com/01-edu/z01.PrintRune"
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
		temp := 0
		for i := 0; i < len(args[0]); i++ {
			for j := temp; j < len(args[1]); j++ {
				if args[0][i] == args[1][j] {
					res = res + string(args[0][i])
					temp = j + 1
					break
				}
			}
		}
		if len(res) == len(args[0]) {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	} else {
// 		result := ""
// 		temp := 0
// 		for i := 0; i < len(args[0]); i++ {
// 			for j := temp; j < len(args[1]); j++ {
// 				if args[0][i] == args[1][j] {
// 					temp = j + 1
// 					result = result + string(args[0][i])
// 					break
// 				}
// 			}
// 		}
// 		if len(args[0]) == len(result) {
// 			z01.PrintRune('1')
// 		} else {
// 			z01.PrintRune('0')
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
