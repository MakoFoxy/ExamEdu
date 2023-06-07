// fprime
// Instructions

// Write a program that takes a positive int and displays its prime factors, followed by a newline ('\n').

//     Factors must be displayed in ascending order and separated by *.

//     If the number of arguments is different from 1, if the argument is invalid, or if the integer does not have a prime factor, the program displays nothing.

// Usage

// $ go run . 225225
// 3*3*5*5*7*11*13
// $ go run . 8333325
// 3*3*5*5*7*11*13*37
// $ go run . 9539
// 9539
// $ go run . 804577
// 804577
// $ go run . 42
// 2*3*7
// $ go run . a
// $ go run . 0
// $ go run . 1
// $

// allowedFunctions
// 0	"strconv.Atoi"
// 1	"os.Args"
// 2	"fmt.*"
// 3	"--allow-builtin"

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	} else {
		n, err := strconv.Atoi(args[0])
		if args[0] == " " {
			return
		}
		if err == nil {
			if n <= 1 {
				return
			}
			for i := 2; i < n*2; i++ {
				if n%i == 0 {
					fmt.Print(i)
					if n/i != 1 {
						fmt.Print("*")
					}
					n = n / i
					i = 1
				}
			}
			fmt.Println()
		} else {
			return
		}
	}
}

// func main() {
// 	// if len(os.Args) == 1 {
// 	// 	return
// 	// }
// 	args := os.Args[1:]

// 	if len(args) == 1 {
// 		n, _ := strconv.Atoi(args[0])
// 		if n <= 1 {
// 			return
// 		}
// 		for i := 2; i < n*2; i++ { // 24
// 			if n%i == 0 {
// 				fmt.Print(i) // 2*2*3
// 				if n/i != 1 {
// 					fmt.Print("*") // 2*2*3
// 				}
// 				n = n / i
// 				i = 1
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		return
// 	}
// 	n, _ := strconv.Atoi(args[0])

// 	for i := 2; i < n*2; i++ { // 24
// 		if n%i == 0 { // 12/2=6/2=3/3=1/2=%1
// 			fmt.Print(i) // 2*2*3
// 			if n/i != 1 {
// 				fmt.Print("*") // 2*2*3
// 			}
// 			n = n / i // 12/2=6/2=3/3=1
// 			i = 1
// 		}
// 	}
// 	fmt.Println()
// }

// 12 % 2 = 0
// 2 *
// 6 % 2 = 0
// 2 *
// 3 % 2 = 1
// -
// 1 % 2 = 1
// 1 / 2 =

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	n, _ := strconv.Atoi(args[0])
// 	if n <= 1 {
// 		return
// 	}

// 	for i := 2; i < n*2; i++ { //24
// 		if n%i == 0 {   //12/2 =6/2=3
// 			fmt.Print(i)//2*2*3
// 			if n/i != 1 {
// 				fmt.Print("*")//2*2*3
// 			}
// 			n = n / i
// 			i = 1

// 		}
// 	}
// 	fmt.Println()
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	num, _ := strconv.Atoi(args[0])
// 	if num <= 1 {
// 		return
// 	}
// 	Fprime(num)
// 	fmt.Println()
// }

// func Fprime(n int) {
// 	for i := 2; i < n*2; i++ {
// 		if n%i == 0 {
// 			fmt.Print(i)
// 			if n/i != 1 {
// 				fmt.Print("*")
// 			}
// 			Fprime(n / i)
// 			break
// 		}
// 	}
// }
