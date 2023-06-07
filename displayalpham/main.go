// displayalpham
// Instructions

// Write a program that displays the alphabet, with even letters in uppercase, and odd letters in lowercase, followed by a newline ('\n').
// Usage

// student@ubuntu:~/piscine-go/displayalpham$ go build
// student@ubuntu:~/piscine-go/displayalpham$ ./displayalpham | cat -e
// aBcDeFgHiJkLmNoPqRsTuVwXyZ$
// student@ubuntu:~/piscine-go/displayalpham$
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import "github.com/01-edu/z01"

func main() {
	z := 'a'

	for i := 'A'; i <= 'Z'; i++ {
		if i%2 == 0 {
			z01.PrintRune(i)
		} else {
			z01.PrintRune(z)
		}
		z++
	}
	z01.PrintRune('\n')
}

// func main() {
// 	v := 'a'
// 	for i := 'A'; i <= 'Z'; i++ {
// 		if i%2 == 0 {
// 			z01.PrintRune(i)
// 		} else {'\n
// 			z01.PrintRune(v)
// 		}
// 		v++
// 	}
// 	z01.PrintRune('\n')
// }
