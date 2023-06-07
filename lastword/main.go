// lastword
// Instructions
// Write a program that takes a string and displays its last word, followed by a newline ('\n').

// A word is a section of string delimited by spaces or by the start/end of the string.

// The output will be followed by a newline ('\n').

// If the number of arguments is different from 1, or if there are no word, the program displays nothing.

// Usage
// $ go run . "FOR PONY" | cat -e
// PONY$
// $ go run . "this        ...       is sparta, then again, maybe    not" | cat -e
// not$
// $ go run . "  "
// $ go run . "a" "b"
// $ go run . "  lorem,ipsum  " | cat -e
// lorem,ipsum$
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

	res := ""

	if len(args) != 1 {
		return
	} else {
		for i := len(args[0]) - 1; i > 0; i-- {
			if args[0][i] != ' ' {
				res = string(args[0][i]) + res
			} else if args[0][i] == ' ' && res != "" {
				break
			}
		}
		if res == "" {
			return
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]

// 	result := ""

// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for i := len(args[0]) - 1; i >= 0; i-- {
// 			if args[0][i] != ' ' {
// 				result = string(args[0][i]) + result
// 			} else if args[0][i] == ' ' && result != "" {
// 				break
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	result := ""

// 	if len(args) != 1 {
// 		return
// 	} else {

// 		for i := len(args[0]) - 1; i >= 0; i-- {
// 			if args[0][i] != ' ' {
// 				result = string(args[0][i]) + result
// 			} else if args[0][i] == ' ' && result != "" {
// 				break
// 			}
// 		}
// 		// for _, w := range args[0] {
// 		// 	if w != ' ' {
// 		// 		result = result + string(args[0])
// 		// 	} else if w == ' ' && result != "" {
// 		// 		break
// 		// 	}
// 		// }
// 		// for i := 0; i < len(args[0]); i++ {
// 		// 	if args[0][i] != ' ' {
// 		// 		result = result + string(args[0])
// 		// 	} else if args[0][i] == ' ' && result != "" {
// 		// 		break
// 		// 	}
// 		// }
// 		for _, v := range string(result) {
// 			z01.PrintRune(v)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	result := ""

// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for i := len(args[0]) - 1; i >= 0; i-- {
// 			if args[0][i] != ' ' {
// 				result = string(args[0][i]) + result
// 			} else if args[0][i] == ' ' && result != "" {
// 				break
// 			}
// 		}
// 		for _, v := range result {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	result := ""

// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for i := len(args[0]) - 1; i >= 0; i-- {
// 			if args[0][i] != ' ' {
// 				result = string(args[0][i]) + result
// 			} else if args[0][i] == ' ' && result != "" {
// 				break
// 			}
// 		}
// 		for v := len(string(result)) - 1; v >= 0; v-- {
// 			z01.PrintRune(rune(v) + 48)
// 		}
// 		for x := len(string(result)) - 1; x >= 0; x-- {
// 			z01.PrintRune(rune(result[x]))
// 		}
// 		for i := 0; i < len(result); i++ {
// 			z01.PrintRune(rune(result[i]))
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
