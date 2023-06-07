// printhex
// Instructions

// Write a program that takes a positive (or zero) number expressed in base 10, and that displays it in base 16 (with lowercase letters) followed by a newline ('\n').

//     If the number of arguments is different from 1, the program displays nothing.
//     Error cases have to be handled as shown in the example below.

// Usage

// $ go run . 10
// a
// $ go run . 255
// ff
// $ go run . 5156454
// 4eae66
// $ go run .
// $ go run . "123 132 1" | cat -e
// ERROR$
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.*"
// 2	"strconv.*"
// 3	"--allow-builtin"

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		a := "ERROR"

		for _, w := range a {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
		return
	}

	res := strconv.FormatInt(int64(n), 16)

	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		for _, v := range "ERROR" {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	res := strconv.FormatInt(int64(n), 16)
// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) == 1 {
// 		n, _ := strconv.Atoi(args[0])

// 		res := strconv.FormatInt(int64(n), 16)

// 		for _, v := range res {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 	} else if len(args) != 1 {
// 		err := "ERROR"

// 		for _, w := range err {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		for _, v := range "ERROR" {
// 			z01.PrintRune(v)
// 		}
// 		return
// 	}
// 	res := strconv.FormatInt(int64(n), 16)
// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }
