// gcd
// Instructions

// Write a program that takes two string representing two strictly positive integers that fit in an int.

// The program displays their greatest common divisor followed by a newline ('\n').

// If the number of arguments is different from 2, the program displays nothing.

// All arguments tested will be positive int values.
// Usage

// $ go run . 42 10 | cat -e
// 2$
// $ go run . 42 12
// 6
// $ go run . 14 77
// 7
// $ go run . 17 3
// 1
// $ go run .
// $ go run . 50 12 4
// $
// allowedFunctions
// 0	"fmt.*"
// 1	"os.Args"
// 2	"strconv.Atoi"
// 3	"--allow-builtin"

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	} else {
		n1, _ := strconv.Atoi(args[0])
		n2, _ := strconv.Atoi(args[1])

		if n1 <= 0 && n2 <= 0 {
			return
		}
		if n1 <= 0 || n2 <= 0 {
			return
		}
		if n1 > 0 && n2 > 0 {
			for n1 != n2 {
				if n1 > n2 {
					n1 = n1 - n2
				} else {
					n2 = n2 - n1
				}
			}
			fmt.Println(n1)
		}
	}
}

// func main() {
// 	if len(os.Args) == 3 {
// 		num1, err1 := strconv.Atoi(os.Args[1])
// 		num2, err2 := strconv.Atoi(os.Args[2])
// 		if err1 != nil && err2 != nil {
// 			return
// 		}
// 		res := 0
// 		for i := 1; i <= num1 && i <= num2; i++ {
// 			if num1%i == 0 && num2%i == 0 {
// 				res = i
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

// func main() {
// 	if len(os.Args) != 3 {
// 		return
// 	}
// 	f, _ := strconv.Atoi(os.Args[1])
// 	s, _ := strconv.Atoi(os.Args[2])
// 	var min int
// 	if f > s {
// 		min = s
// 	} else {
// 		min = f
// 	}
// 	for i := min; i >= 1; i-- {
// 		if f%i == 0 && s%i == 0 {
// 			fmt.Println(i)
// 			return
// 		}
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 2 {
// 		return
// 	} else {
// 		n1, _ := strconv.Atoi(args[0])
// 		n2, _ := strconv.Atoi(args[1])
// 		if n1 > 0 && n2 > 0 {
// 			for n1 != n2 {
// 				if n1 > n2 {
// 					n1 = n1 - n2
// 				} else {
// 					n2 = n2 - n1
// 				}
// 			}
// 			fmt.Println(n1)
// 		}
// 	}
// }
