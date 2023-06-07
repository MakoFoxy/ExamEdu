// printbits
// Instructions

// Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

//     If the the argument is not a number, the program should print 00000000.

// Usage :

// $ go run . 1
// 00000001$
// $ go run . 192
// 11000000$
// $ go run . a
// 00000000$
// $ go run . 1 1
// $ go run .
// $

// allowedFunctions
// 0	"strconv.Atoi"
// 1	"os.Args"
// 2	"github.com/01-edu/z01.PrintRune"
// 3	"--allow-builtin"

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		n, _ := strconv.Atoi(args[0])
		arr := make([]int, 8)
		for i := len(arr) - 1; i >= 0; i-- {
			arr[i] = n % 2
			n = n / 2
		}
		for _, w := range arr {
			z01.PrintRune(rune(w + 48))
		}
		z01.PrintRune('\n')
	} else {
		return
	}
}

// func Atoi(n string) int {
// 	s := 0
// 	temp := 1
// 	if n[0] == '-' {
// 		s = -1
// 		n = n[1:]
// 	}

// 	for w := range n {
// 		s = s * 10
// 		s = s + int(n[w]-48)
// 	}
// 	return s * temp
// }

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	args := os.Args[1:]

// 	if len(args) == 1 {
// 		n, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			fmt.Println("00000000")
// 		} else {
// 			fmt.Println(toBin(n))
// 		}
// 	}
// }

// func toBin(n int) string {
// 	res := ""

// 	for n > 0 {
// 		res = string(rune(n%2+48)) + res
// 		n = n / 2
// 	}
// 	for len(res) < 8 {
// 		res = "0" + res
// 	}
// 	return res
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		arr := make([]int, 8)
// 		num, err := strconv.Atoi(os.Args[1])
// 		if err == nil {
// 			for i := len(arr) - 1; i >= 0; i-- {
// 				arr[i] = num % 2
// 				num = num / 2
// 			}
// 		}

// 		for _, v := range arr {
// 			z01.PrintRune(rune(v + 48))
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func Atoi(args string) int {
// 	z := 0
// 	temp := 1
// 	if args[0] == '-' {
// 		temp = -1
// 		args = args[1:]
// 	}
// 	for i := 0; i < len(args); i++ {
// 		z = z * 10
// 		z = z + int(args[i]-48)
// 	}
// 	return temp * z
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		slice := make([]int, 8)
// 		num, err := strconv.Atoi(os.Args[1])
// 		if err != nil {
// 			for _, val := range slice {
// 				z01.PrintRune(rune(val) + 48)
// 			}
// 			z01.PrintRune('\n')
// 			return
// 		}
// 		i := len(slice) - 1
// 		for num != 0 {
// 			slice[i] = num % 2
// 			num /= 2
// 			i--
// 		}
// 		for _, val := range slice {
// 			z01.PrintRune(rune(val) + 48)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
