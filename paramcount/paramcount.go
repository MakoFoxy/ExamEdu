// paramcount
// Instructions
// Write a program that displays the number of arguments passed to it. This number will be followed by a newline ('\n').

// If there is no argument, the program displays 0 followed by a newline.

// Usage
// student@ubuntu:~/piscine-go/paramcount$ go build
// student@ubuntu:~/piscine-go/paramcount$ ./paramcount 1 2 3 5 7 24
// 6
// student@ubuntu:~/piscine-go/paramcount$ ./paramcount 6 12 24 | cat -e
// 3$
// student@ubuntu:~/piscine-go/paramcount$ ./paramcount | cat -e
// 0$
// student@ubuntu:~/piscine-go/paramcount$

// allowedFunctions
// 0	"os.Args"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"strconv.Itoa"
// 3	"--allow-builtin"

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// if len(args) == 0 {
	// 	z01.PrintRune(48)
	// }

	// arr := []rune{}

	// for i := len(args); i > 0; i = i / 10 {
	// 	arr = append(arr, rune(i%10+48))
	// }

	// for v := len(arr) - 1; v >= 0; v-- {
	// 	z01.PrintRune(arr[v])
	// }
	// z01.PrintRune('\n')

	res := strconv.Itoa(len(args))

	for _, w := range res {
		z01.PrintRune(w)
	}
	z01.PrintRune('\n')

	// res := ""

	// for i := len(args); i > 0; i = i / 10 {
	// 	res = string(rune(i%10+48)) + res
	// }

	// for _, w := range res {
	// 	z01.PrintRune(w)
	// }
	// z01.PrintRune('\n')
}

// func main() {
// args := os.Args[1:]
// arr := []rune{}
// if len(args) == 0 {
// 	z01.PrintRune(48)
// }

// for i := len(args); i > 0; i = i/10 {
// 	arr = append(arr, rune(i%10+48))
// }

// for v := len(arr) - 1; v >=0; v-- {
// 	z01.PrintRune(arr[v])
// }
// z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	res := []rune{}

// 	if len(args) == 0 {
// 		z01.PrintRune(48)
// 	}
// 	for i := len(args); i > 0; i = i / 10 {
// 		res = append(res, rune(i%10+48))
// 	}

// 	for v := len(res) - 1; v >= 0; v-- {
// 		z01.PrintRune(res[v])
// 	}
// 	z01.PrintRune('\n')

// }

// func main() {
// args := os.Args[1:]
// res := []rune{}

// if len(args) == 0 {
// 	z01.PrintRune(48)
// }
// for i := len(args); i > 0; i = i / 10 {
// 	res = append(res, rune(i%10+48))
// }

// for v := len(res) - 1; v >= 0; v-- {
// 	z01.PrintRune(res[v])
// }
// z01.PrintRune('\n')

// args := os.Args[1:]

// res := ""
// if len(args) == 0 {
// 	z01.PrintRune(48)
// }
// for i := len(args); i > 0; i = i / 10 {
// 	res = string(rune(i%10+48)) + res
// }

// for _, w := range res {
// 	z01.PrintRune(w)
// }

// z01.PrintRune('\n')

// PrintNbr(len(args))
// res := ""
// if arg == 0 {
// 	z01.PrintRune(48)
// }
// for i := arg; i > 0; i = i / 10 {
// 	res = string(i%10+48) + res
// }
// for _, w := range res {
// 	z01.PrintRune(w)
// }

// z01.PrintRune('\n')

// func PrintNbr(n int) {
// 	for n != 0 {
// 		if n >= 0 {
// 			PrintNbr(n / 10)
// 		}
// 		z01.PrintRune(rune(n%10 + 48))
// 	}
// }

// func main() {
// 	args := os.Args[1:]

// 	arg := len(args)

// 	// res := []rune{}

// 	res := ""

// 	if arg == 0 {
// 		z01.PrintRune(rune(48))
// 	}

// 	// for i := arg; i > 0; i = i / 10 {
// 	// 	res = append(res, rune(i%10+48))
// 	// }
// 	for i := arg; i > 0; i = i / 10 {
// 		res = string(i%10+48) + res
// 	}
// 	// for v := len(res) - 1; v >= 0; v-- {
// 	// 	z01.PrintRune(rune(res[v]))
// 	// }
// 	for _, w := range res {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	res := []rune{}
// 	arg := len(args)

// 	if arg == 0 {
// 		z01.PrintRune(rune(48))
// 	}

// 	for i := arg; i > 0; i = i / 10 {
// 		res = append(res, rune(i%10+48))
// 	}
// 	for v := len(res) - 1; v >= 0; v-- {
// 		z01.PrintRune(res[v])
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	res := []rune{}
// 	args := os.Args[1:]
// 	arg := len(args)
// 	if arg == 0 {
// 		z01.PrintRune(rune(48))
// 	}
// 	for i := arg; i > 0; i = i / 10 {
// 		res = append(res, rune(i%10+48))
// 	}

// 	for v := len(res) - 1; v >= 0; v-- {
// 		z01.PrintRune(res[v])
// 	}
// 	// 	z01.PrintRune('\n')
// 	// 	// for _, w := range res {
// 	// 	// 	z01.PrintRune(w)
// 	// 	// }
// 	// 	// z01.PrintRune('\n')
// 	// 	// for i := 0; i <= arg; i = i / 10 {
// 	// 	// 	res = append(res, rune(i/10+48))
// 	// 	// 	z01.PrintRune(rune(i%10 + 48))
// 	// 	// }
// 	// 	// for _, w := range res {
// 	// 	// 	z01.PrintRune(w)
// 	// 	// }
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	result := []rune{}
// 	arg := len(args)
// 	if arg == 0 {
// 		z01.PrintRune(rune(48))
// 	}
// 	for i := arg; i > 0; i = i / 10 {
// 		result = append(result, rune(i%10+48))
// 	}
// 	for v := len(result) - 1; v >= 0; v-- {
// 		z01.PrintRune(result[v])
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	arg := len(args)
// 	result := []rune{}

// 	for i := arg; i > 0; i = i / 10 {
// 		result = append(result, rune(i%10+48))
// 	}
// 	for v := len(result) - 1; v >= 0; v-- {
// 		z01.PrintRune(result[v])
// 	}
// 	if arg == 0 {
// 		z01.PrintRune(rune(48))
// 	}
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args
// 	arg := len(args)

// 	runes := []rune{}

// 	for i := arg; i > 0; i = i / 10 {
// 		runes = append(runes, rune(i%10+48))
// 	}

// 	for v := len(runes) - 1; v >= 0; v-- {
// 		z01.PrintRune(runes[v])
// 	}
// 	// if arg < 9 {
// 	// 	z01.PrintRune(rune(arg/10 + 48))
// 	// 	z01.PrintRune(rune(arg%10 + 48))
// 	// 	z01.PrintRune('\n')

// 	// } else {
// 	// 	z01.PrintRune(rune(arg) + '0')
// 	// }
// 	z01.PrintRune('\n')
// }

// func main() {
// 	args := os.Args[1:]
// 	len := len(args)

// 	if len > 9 {
// 		z01.PrintRune(rune(len/10 + 48))
// 		z01.PrintRune(rune(len%10 + 48))
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune(rune(len) + '0')
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	arg := os.Args[1:]

// 	len := len(arg)

// 	if len > 9 {
// 		z01.PrintRune(rune(len/10 + 48))
// 		z01.PrintRune(rune(len%10 + 48))
// 		z01.PrintRune('\n')cd ../
// 	} else {
// 		z01.PrintRune(rune(len) + '0')
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	arg := os.Args[1:]
// 	lens := len(arg) // 243 243 / 10 = 24 / 10 = 2 % 10 = 2 2 / 10 = 0
// 	if len > 9 {
// 		num1 := len / 10
// 		num2 := len % 10
// 		z01.PrintRune(rune(num1 + 48))
// 		z01.PrintRune(rune(num2 + 48))
// 		z01.PrintRune('\n')
// 	} else {
// 		z01.PrintRune(rune(len) + '0')
// 		z01.PrintRune('\n')
// 	}
// 	runes := []rune{} // 3,4,2
// 	for i := lens; i > 0; i = i / 10 {
// 		runes = append(runes, rune(i%10+48))
// 	}
// 	for i := len(runes) - 1; i >= 0; i-- {
// 		z01.PrintRune(runes[i])
// 	}
// 	z01.PrintRune('\n')

// }
