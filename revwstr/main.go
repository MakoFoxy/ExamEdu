// revwstr
// Instructions

// Write a program that takes a string as a parameter, and prints its words in reverse, followed by a newline.

//     A word is a sequence of alphanumerical characters.

//     If the number of arguments is different from 1, the program will display nothing.

//     In the parameters that are going to be tested, there will not be any extra spaces. (meaning that there will not be additional spaces at the beginning or at the end of the string and that words will always be separated by exactly one space).

// Usage

// $ go run . "the time of contempt precedes that of indifference"
// indifference of that precedes contempt of time the
// $ go run . "abcdefghijklm"
// abcdefghijklm
// $ go run . "he stared at the mountain"
// mountain the at stared he
// $ go run . "" | cat-e
// $
// $
// allowedFunctions
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

	if len(args) != 1 {
		return
	} else {
		res := ""
		arr := []string{}

		// if args[0] == "" {
		// 	return
		// }
		for _, w := range args[0] {
			if w != ' ' {
				res = res + string(w)
			} else if w == ' ' && res != "" {
				arr = append(arr, res)
				res = ""
			}
		}
		if res != "" {
			arr = append(arr, res)
		}

		for i := len(arr) - 1; i >= 0; i-- {
			for _, v := range arr[i] {
				z01.PrintRune(v)
			}
			if i != 0 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	var res []string
// 	word := ""
// 	for _, v := range args[0] {
// 		if v != ' ' {
// 			word = word + string(v)
// 		}
// 		if v == ' ' {
// 			res = append(res, word)
// 			word = ""
// 		}
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	for i := len(res) - 1; i >= 0; i-- {
// 		for _, j := range res[i] {
// 			z01.PrintRune(j)
// 		}
// 		if i != 0 {
// 			z01.PrintRune(' ')
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	// if len(os.Args) == 2 {
// 	var res []string
// 	word := ""
// 	for _, v := range os.Args[1] {
// 		if v != ' ' {
// 			word = word + string(v)
// 		}
// 		if v == ' ' {
// 			res = append(res, word)
// 			word = ""
// 		}
// 	}
// 	if word != "" {
// 		res = append(res, word)
// 	}
// 	for i := len(res) - 1; i >= 0; i-- {
// 		for _, j := range res[i] {
// 			z01.PrintRune(j)
// 		}
// 		if i != 0 {
// 			z01.PrintRune(' ')
// 		}
// 	}
// 	z01.PrintRune('\n')
// 	// }
// }
// //
// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for _, w := range revwstr(args[0]) {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func revwstr(s string) string {
// 	a := split(s)
// 	result := ""
// 	for i := len(a) - 1; i >= 0; i-- {
// 		result += string(a[i]) + " "
// 	}
// 	return result
// }

// func split(s string) []string {
// 	result := []string{}
// 	res := ""
// 	for i, w := range s {
// 		if w != ' ' {
// 			res += string(w)
// 		}
// 		if len(res) != 0 && (w == ' ' || i == len(s)-1) {
// 			result = append(result, res)
// 			res = ""
// 		}
// 	}
// 	return result
// }
