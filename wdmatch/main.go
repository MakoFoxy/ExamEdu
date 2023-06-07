// wdmatch
// Instructions
// Write a program that takes two string and checks whether it is possible to write the first string with characters from the second string, while respecting the order in which these characters appear in the second string.

// If it is possible, the program displays the string followed by a newline ('\n'), otherwise it simply displays a newline.

// If the number of arguments is different from 2, the program displays a newline.

// Usage
// student@ubuntu:~/piscine-go/test$ go build
// student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbreda"
// faya
// student@ubuntu:~/piscine-go/test$ ./test "faya" "fgvvfdxcacpolhyghbred"

// student@ubuntu:~/piscine-go/test$ ./test "error" rrerrrfiiljdfxjyuifrrvcoojh

// student@ubuntu:~/piscine-go/test$ ./test "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
// quarante deux

// student@ubuntu:~/piscine-go/test$ ./test

// student@ubuntu:~/piscine-go/test$

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

	res := ""
	temp := 0

	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	} else {
		for i := 0; i < len(args[0]); i++ {
			for j := temp; j < len(args[1]); j++ {
				if args[0][i] == args[1][j] {
					temp = j + 1
					res = res + string(args[0][i])
					break
				}
			}
		}
		if len(res) == len(args[0]) {
			for _, w := range res {
				z01.PrintRune(w)
			}
		} else {
			return
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	temp := 0
// 	result := ""
// 	if len(args) != 2 {
// 		return
// 	}
// 	for i := 0; i < len(args[0]); i++ {
// 		for j := temp; j < len(args[1]); j++ {
// 			if args[0][i] == args[1][j] {
// 				temp = j + 1
// 				result = result + string(args[0][i])
// 				break
// 			}
// 		}
// 	}
// 	if len(result) == len(args[0]) {
// 		for _, letter := range result {
// 			z01.PrintRune(letter)
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	res := ""
// 	temp := 0
// 	if len(args) != 2 {
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	for i := 0; i < len(args[0]); i++ {
// 		for j := temp; j < len(args[1]); j++ {
// 			if args[0][i] == args[1][j] {
// 				j = j + 1
// 				res = res + string(args[0][i])
// 				break
// 			}
// 		}
// 	}
// 	if len(res) == len(args[0]) {
// 		for _, w := range res {
// 			z01.PrintRune(w)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	temp := 0
// 	result := ""

// 	if len(args) != 2 {
// 		return
// 	}
// 	for i := 0; i < len(args[0]); i++ {
// 		for j := temp; j < len(args[1]); j++ {
// 			if args[0][i] == args[1][j] {
// 				temp = j + 1
// 				result = result + string(args[0][i])
// 				break
// 			}
// 		}
// 	}
// 	if len(result) == len(args[0]) {
// 		for _, v := range result {
// 			z01.PrintRune(v)
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	temp := 0
// 	result := ""
// 	if len(args) != 2 {
// 		return
// 	}
// 	for i := 0; i < len(args[0]); i++ {
// 		for j := temp; j < len(args[1]); j++ {
// 			if args[0][i] == args[1][j] {
// 				temp = j + 1
// 				result = result + string(args[0][i])
// 				break
// 			}
// 		}
// 	}
// 	if len(result) == len(args[0]) {
// 		for _, letter := range result {
// 			z01.PrintRune(letter)
// 		}
// 	} else {
// 		return
// 	}
// 	z01.PrintRune('\n')
// }
