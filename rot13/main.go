// rot13
// Instructions
// Write a program that takes a string and displays it, replacing each of its letters by the letter 13 spots ahead in alphabetical order.

// 'z' becomes 'm' and 'Z' becomes 'M'. The case of the letter stays the same.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program displays nothing.

// Usage
// $ go run . "abc"
// nop
// $ go run . "hello there"
// uryyb gurer
// $ go run . "HELLO, HELP"
// URYYB, URYC
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

	res := []rune{}

	if len(args) != 1 {
		return
	} else {
		for _, w := range args[0] {
			if w >= 'a' && w <= 'z' {
				if w+13 >= 'z' {
					res = append(res, rune(w+13-26))
				} else {
					res = append(res, rune(w+13))
				}
			} else if w >= 'A' && w <= 'Z' {
				if w+13 >= 'Z' {
					res = append(res, rune(w+13-26))
				} else {
					res = append(res, rune(w+13))
				}
			} else {
				res = append(res, rune(w))
			}
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	res := []rune{}

// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for _, v := range args[0] {
// 			if v >= 'a' && v <= 'z' {
// 				if v+13 >= 'z' {
// 					res = append(res, rune(v+13-26))
// 				} else {
// 					res = append(res, rune(v+13))
// 				}
// 			} else if v >= 'A' && v <= 'Z' {
// 				if v+13 >= 'Z' {
// 					res = append(res, rune(v+13-26))
// 				} else {
// 					res = append(res, rune(v+13))
// 				}
// 			} else {
// 				res = append(res, rune(v))
// 			}
// 		}
// 		for _, w := range res {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		for _, ch := range os.Args[1] {
// 			if ch >= 'A' && ch <= 'Z' {
// 				if ch >= 'A'+13 {
// 					ch -= 13
// 				} else {
// 					ch += 13
// 				}
// 			}
// 			if ch >= 'a' && ch <= 'z' {
// 				if ch >= 'a'+13 {
// 					ch -= 13
// 				} else {
// 					ch += 13
// 				}
// 			}
// 			z01.PrintRune(ch)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	result := []rune{}
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for _, letter := range args[0] {
// 			if letter >= 'a' && letter <= 'z' {
// 				if letter+13 >= 'z' {
// 					result = append(result, letter+13-26)
// 				} else {
// 					result = append(result, letter+13)
// 				}
// 			} else if letter >= 'A' && letter <= 'Z' {
// 				if letter+13 >= 'Z' {
// 					result = append(result, letter+13-26)
// 				} else {
// 					result = append(result, letter+13)
// 				}
// 			} else {
// 				result = append(result, letter)
// 			}
// 		}
// 		for _, w := range string(result) {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
