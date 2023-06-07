// ispowerof2
// Instructions

// Write a program that determines if a given number is a power of 2.

// This program must print true if the given number is a power of 2, otherwise it has to print false.

//     If there is more than one or no argument, the program should print nothing.

//     When there is only one argument, it will always be a positive valid int.

// Usage :

// $ go run . 2 | cat -e
// true$
// $ go run . 64
// true
// $ go run . 513
// false
// $ go run .
// $ go run . 64 1024
// $
// "0	"github.com/01-edu/z01.PrintRune"
// 1	"os.*"
// 2	"strconv.Atoi"

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

//"strconv"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	n, _ := strconv.Atoi(args[0])
	// if err == nil {
	// for i := 0; i < len(args[0]); i++ {
	if n&(n-1) == 0 && n > 0 {
		res := "true"
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	} else {
		res := "false"
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
	//}
	//}
}

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		n := Atoi(args[0])
// 		if n&(n-1) == 0 && n > 0 {
// 			result := "true"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		} else {
// 			result := "false"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }
func Atoi(args string) int {
	s := 0
	temp := 1

	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}

	for v := range args {
		s = s * 10
		s = s + int(args[v]-48)
	}
	return s * temp
}

// func Atoi(args string) int {
// 	s := 0
// 	temp := 1

// 	if args[0] == '-' {
// 		temp = -1
// 		args = args[1:]
// 	}

// 	for i := 0; i < len(args); i++ {

// 		s = s * 10
// 		s = s + int(args[i]-48)

// 	}
// 	return temp * s
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		n, _ := strconv.Atoi(args[0])
// 		if n&(n-1) == 0 && n > 0 {
// 			result := "true"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		} else {
// 			result := "false"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func Atoi(arg string) int {
// 	s := 0
// 	temp := 1
// 	if arg[0] == '-' {
// 		temp = -1
// 		arg = arg[1:]
// 	}

// 	for i := 0; i < len(arg); i++ {
// 		s = s * 10
// 		s = s + int(arg[i]-'0')

// 	}
// 	return s * temp
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		return
// 	} else {
// 		n := Atoi(args[0])
// 		if n&(n-1) == 0 && n > 0 {
// 			result := "true"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		} else {
// 			result := "false"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func Atoi(args string) int {
// 	s := 0
// 	sign := 1

// 	if args[0] == '-' {
// 		sign = -1
// 		args = args[1:]
// 	}
// 	for i := 0; i < len(args); i++ {
// 		s = s * 10
// 		s = s + int(args[i]-'0')
// 	}
// 	return sign * s
// }

// func Atoi(args string) int {
// 	s := 0
// 	for i := 0; i < len(args); i++ {
// 		s = s * 10
// 		s = s + int(args[i]-'0')
// 	}
// 	return s
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		n, _ := strconv.Atoi(args[0])
// 		if n&(n-1) == 0 && n > 0 {
// 			result := "true"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		} else {
// 			result := "false"
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func ispowerof2(n int) bool {
// 	return n&(n-1) == 0 && n > 0
// }
