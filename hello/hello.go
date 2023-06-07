// hello
// Instructions
// Write a program that displays "Hello World!" followed by a newline ('\n').

// Usage
// student@ubuntu:~/piscine-go/test$ go build
// student@ubuntu:~/piscine-go/test$ ./test
// Hello World!
// student@ubuntu:~/piscine-go/test$
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import "github.com/01-edu/z01"

func main() {
	ch := "Hello World!"

	// for _, w := range ch {
	// 	z01.PrintRune(w)
	// }
	a := []rune{}
	for i := len(ch) - 1; i >= 0; i-- {
		a = append(a, rune(ch[i]))
	}
	for _, w := range a {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	ch := "Hello World!"

// 	// for _, w := range ch {
// 	// 	z01.PrintRune(w)
// 	// }
// 	for i := len(ch) - 1; i >= 0; i-- {
// 		z01.PrintRune(rune(ch[i]))
// 	}

// 	z01.PrintRune('\n')
// }

// func main() {
// 	ch := "Hello World!"
// 	for _, v := range ch {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	ch := "Hello World!"

// 	for _, v := range ch {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }
// func main() {
// 	ch := "Hello World!"
// 	a := []rune{}
// 	// res :=""
// 	for i := len(ch) - 1; i >= 0; i-- {
// 		a = append(a, rune(ch[i]))
// 		// z01.PrintRune(rune(ch[i]))
// 	}
// 	for _, w := range a {
// 		z01.PrintRune(rune(w))
// 	}
// 	z01.PrintRune('\n')
// }
