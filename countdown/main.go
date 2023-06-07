// countdown
// Instructions
// Write a program that displays all digits in descending order, followed by a newline ('\n').

// Usage
// student@ubuntu:~/piscine-go/test$ go build
// student@ubuntu:~/piscine-go/test$ ./test
// 9876543210
// student@ubuntu:~/piscine-go/test$
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"--allow-builtin"

package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	for i := '9'; i >= '0'; i-- {
// 		z01.PrintRune(i)
// 	}
// 	z01.PrintRune('\n')
// }
