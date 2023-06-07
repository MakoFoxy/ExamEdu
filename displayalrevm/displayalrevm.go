// displayalrevm
// Instructions
// Write a program that displays the alphabet in reverse, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').

// Usage
// student@ubuntu:~/piscine-go/displayalrevm$ go build
// student@ubuntu:~/piscine-go/displayalrevm$ ./displayalrevm | cat -e
// zYxWvUtSrQpOnMlKjIhGfEdCbA$
// // student@ubuntu:~/piscine-go/displayalrevm$
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import "github.com/01-edu/z01"

func main() {
	z := 'Z'

	for i := 'z'; i >= 'a'; i-- {
		if i%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(z)
		}
		z--
	}
	z01.PrintRune('\n')
}

// func main() {
// 	v := 'Z'
// 	for i := 'z'; i >= 'a'; i-- {
// 		if i%2 == 0 {
// 			z01.PrintRune(i)
// 		} else {
// 			z01.PrintRune(v)
// 		}
// 		v--
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	v := 'Z'
// 	for i := 'z'; i >= 'a'; i-- {
// 		if i%2 == 0 {
// 			z01.PrintRune(i)
// 		} else {
// 			z01.PrintRune(v)
// 		}
// 		v--
// 	}
// 	z01.PrintRune('\n')
// }
