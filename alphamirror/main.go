// alphamirror
// Instructions
// Write a program called alphamirror that takes a string as argument and displays this string after replacing each alphabetical character with the opposite alphabetical character.

// The case of the letter remains unchanged, for example :

// 'a' becomes 'z', 'Z' becomes 'A' 'd' becomes 'w', 'M' becomes 'N'

// The final result will be followed by a newline ('\n').

// If the number of arguments is different from 1, the program prints a new line.

// Usage
// $ go run . "abc"
// zyx$
// $
// $ go run . "My horse is Amazing." | cat -e
// Nb slihv rh Znzarmt.$
// $
// $ go run .
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

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}
	res := ""
	for _, w := range args[0] {
		if w >= 'a' && w <= 'z' {
			res = res + string(rune('a'+'z'-w))
		} else if w >= 'A' && w <= 'Z' {
			res = res + string(rune('A'+'Z'-w))
		} else {
			res = res + string(rune(w))
		}
	}
	for _, v := range res {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	res := ""
// 	for _, w := range args[0] {
// 		if w >= 'a' && w <= 'z' {
// 			res = res + string('a'+'z'-w)
// 		} else if w >= 'A' && w <= 'Z' {
// 			res = res + string('A'+'Z'-w)
// 		} else {
// 			res = res + string(w)
// 		}
// 	}

// 	for _, v := range res {
// 		z01.PrintRune(v)
// 	}

// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) == 1 {
// 		arr := []rune(args[0])
// 		for i, w := range arr {
// 			if w >= 'a' && w <= 'z' {
// 				arr[i] = 'a' + 'z' - w
// 			} else if w >= 'A' && w <= 'Z' {
// 				arr[i] = 'A' + 'Z' - w
// 			}
// 		}
// 		for _, v := range string(arr) {
// 			z01.PrintRune(v)
// 		}
// 		z01.PrintRune('\n')
// 	} else {
// 		return
// 	}
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}
// 	args := os.Args[1:]
// 	arr := []rune(args[0])
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		for i, w := range arr {
// 			if w >= 'a' && w <= 'z' {
// 				arr[i] = 'a' + 'z' - w
// 			} else if w >= 'A' && w <= 'Z' {
// 				arr[i] = 'A' + 'Z' - w
// 			}
// 		}

// 		for _, w := range string(arr) {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		for _, v := range os.Args[1] {
// 			if v >= 'A' && v <= 'Z' {
// 				v = ('Z' - v + 'A')
// 			} else if v >= 'a' && v <= 'z' { // val >= 97 && val <= 122
// 				v = ('z' - v + 'a') // 122 - 98 + 97 =121
// 			}
// 			z01.PrintRune(v)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		pralph := []rune(args[0])
// 		for i, v := range pralph {
// 			if v >= 'a' && v <= 'z' {
// 				pralph[i] = 'a' + 'z' - v
// 			} else if v >= 'A' && v <= 'Z' {
// 				pralph[i] = 'A' + 'Z' - v
// 			}
// 		}
// 		for _, z := range string(pralph) {
// 			z01.PrintRune(z)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		pralph := []rune(args[0])
// 		for i, v := range pralph {
// 			if v >= 'a' && v <= 'z' {
// 				pralph[i] = 'a' + 'z' - v
// 			} else if v >= 'A' && v <= 'z' {
// 				pralph[i] = 'A' + 'z' - v
// 			}
// 		}
// 		for _, w := range string(pralph) {
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
// 	var result []rune
// 	var arr []rune
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		// result := []rune(args[0])
// 		for _, letter := range args[0] {
// 			if letter >= 'a' && letter <= 'z' {
// 				result = append(result, letter)
// 				// result[i] = 'a' + 'z' - letter
// 			} else if letter >= 'A' && letter <= 'Z' {
// 				result = append(result, letter)
// 				// result[i] = 'A' + 'Z' - letter
// 			}
// 		}
// 		for i := len(result) - 1; i >= 0; i-- {
// 			z01.PrintRune(result[i])
// 		}
// 		z01.PrintRune('\n')

// 		for j := 'z'; j >= 'a'; j-- {

// 			arr = append(arr, j)
// 			z01.PrintRune(j)
// 		}
// 		z01.PrintRune('\n')
// 		if len(result) <= len(arr) {
// 			for c := len(arr); c > 0; c-- {
// 				for p := len(result); p > 0; p-- {
// 					if arr[c] != result[p] || arr[c] == result[p] {
// 						z01.PrintRune('\n')

// 						z01.PrintRune(rune(c))
// 					}
// 				}
// 			}
// 			for _, x := range result {
// 				z01.PrintRune(x)
// 			}
// 			z01.PrintRune('\n')
// 			for _, z := range arr {
// 				z01.PrintRune(z)
// 			}
// 		}

// 		// for _, w := range string(result) {
// 		// 	z01.PrintRune(w)
// 		// }
// 		z01.PrintRune('\n')
// 	}
// }
