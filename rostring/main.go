// rostring
// Instructions

// Write a program that takes a string and displays this string after rotating it one word to the left.

// Thus, the first word becomes the last, and others stay in the same order.

// A word is a sequence of alphanumerical characters.

// Words will be separated by only one space in the output.

// If the number of arguments is different from 1, the program displays a newline.
// Usage

// $ go run . "abc   " | cat -e
// abc$
// $ go run . "Let there     be light"
// there be light Let
// $ go run . "     AkjhZ zLKIJz , 23y"
// zLKIJz , 23y AkjhZ
// $ go run . | cat -e
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
		if args[0] == "" {
			z01.PrintRune('\n')
			return
		}
		res := ""
		arr := []string{}

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

		str := ""

		for _, v := range arr[1:] {
			str = str + v + " "
		}
		str1 := str + arr[0]

		for _, z := range str1 {
			z01.PrintRune(z)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 1 {
// 		var res []string
// 		word := ""
// 		for _, v := range os.Args[1] {
// 			if v != ' ' {
// 				word = word + string(v)
// 			} else if v == ' ' && word != "" {
// 				res = append(res, word)
// 				word = ""
// 			} else {
// 				continue
// 			}
// 		}
// 		if word != "" {
// 			res = append(res, word)
// 		}
// 		str := ""
// 		for _, w := range res[1:] {
// 			str = str + w + " "
// 		}
// 		str1 := str + res[0]
// 		for _, n := range str1 {
// 			z01.PrintRune(n)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune('\n')
// 		return
// 	}
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		// var res []string
// 		var res []string
// 		word := ""
// 		for _, v := range os.Args[1] {
// 			if v != ' ' {
// 				word = word + string(v)
// 			} else if v == ' ' && word != "" {
// 				res = append(res, word)
// 				word = ""
// 			} else {
// 				continue
// 			}
// 		}
// 		if word != "" {
// 			res = append(res, word)
// 		}
// 		str := ""
// 		//  cycle := len(res) - 1

// 		// for i := 1; i < len(res); i++ {
// 		// 	str = str + res[i] + " "
// 		// }
// 		// str1 := str + res[0]
// 		for _, w := range res[1:] {
// 			str = str + w + " "
// 		}
// 		str1 := str + res[0]
// 		// for i := len(res) - 1; i > 0; i-- {
// 		// 	val = res + string(' ') + val
// 		// }
// 		// val = val + res
// 		for _, n := range str1 {
// 			z01.PrintRune(n)
// 		}
// 		z01.PrintRune('\n')
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		var res []string
// 		word := ""
// 		for _, v := range os.Args[1] {
// 			if v != ' ' {
// 				word = word + string(v)
// 			} else if v == ' ' && word != "" {
// 				res = append(res, word)
// 				word = ""
// 			} else {
// 				continue
// 			}
// 		}
// 		if word != "" {
// 			res = append(res, word)
// 		}
// 		val := ""
// 		for i := len(res) - 1; i > 0; i-- {
// 			val = res[i] + " " + val
// 		}
// 		val += res[0]
// 		for _, v := range val {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		z01.PrintRune('\n')
// 		return
// 	} else {
// 		for _, w := range rostring(args[0]) {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func rostring(s string) string {
// 	a := split(s)
// 	result := ""

// 	for i := 1; i < len(a); i++ {
// 		result = result + a[i] + " "
// 	}
// 	result = result + a[0]
// 	return result
// }

// func split(s string) []string {
// 	result := []string{}
// 	res := ""
// 	for i, w := range s {
// 		if w != ' ' {
// 			res = res + string(w)
// 		}
// 		if len(res) != 0 && (w == ' ' || i == len(s)-1) {
// 			result = append(result, res)
// 			res = ""
// 		}
// 	}
// 	return result
// }
