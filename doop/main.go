// doop
// Instructions
// Write a program that is called doop.

// The program has to be used with three arguments:

// A value
// An operator, one of : +, -, /, *, %
// Another value
// In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.

// The program has to handle the modulo and division operations by 0 as shown on the output examples below.

// Usage
// $ go run .
// $ go run . 1 + 1 | cat -e
// 2$
// $ go run . hello + 1
// $ go run . 1 p 1
// $ go run . 1 / 0 | cat -e
// No division by 0$
// $ go run . 1 % 0 | cat -e
// No modulo by 0$
// $ go run . 9223372036854775807 + 1
// $ go run . -9223372036854775809 - 3
// $ go run . 9223372036854775807 "*" 3
// $ go run . 1 "*" 1
// 1
// $ go run . 1 "*" -1
// -1
// $

// allowedFunctions
// 0	".."
// 1	"os.*"
// 2	"fmt.Printf"
// 3	"--allow-builtin"

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 {
		args := os.Args[1:]
		a, aErr := Atoi(args[0])
		b, bErr := Atoi(args[2])
		if !aErr && !bErr {
			switch args[1] {
			case "+":
				c := a + b
				if (c > a) == (b > 0) {
					fmt.Printf("%d\n", c)
				}
			case "-":
				c := a - b
				if (c < a) == (b > 0) {
					fmt.Printf("%d\n", c)
				}
			case "*":
				c := a * b
				if c/a == b {
					fmt.Printf("%d\n", c)
				}
			case "/":
				if b == 0 {
					fmt.Printf("No division by 0\n")
					return
				}
				fmt.Printf("%d\n", a/b)
			case "%":
				if b == 0 {
					fmt.Printf("No modulo by 0\n")
					return
				}
				fmt.Printf("%d\n", a%b)
			default:
				return
			}
		}
	}
}

func Atoi(args string) (int, bool) {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for _, w := range args {
		if w < '0' || w > '9' {
			os.Exit(0) // return
		}
		x := s * 10
		y := int(w - 48)
		t := x + y
		if (t > x) == (y > 0) {
			s = t
		} else {
			return -1, true
		}
	}
	return s * temp, false
}

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 3 {
// 		return
// 	}

// 	a := Atoi(args[0])
// 	b := Atoi(args[2])
// 	op := args[1]

// 	maxInt := int(^uint(0) >> 1)
// 	minInt := -maxInt - 1
// 	// if len(args[0]) >= 19 || len(args[2]) >= 19 {
// 	// 	return
// 	// }

// 	// if (a > 0 && b > 0 && a+b < 0) || (a < 0 && b < 0 && a+b > 0) {
// 	// 	// os.Exit(0)
// 	// 	return
// 	// }
// 	if (a > 0 && b > 0 && a+b < 0) || a > maxInt || b > maxInt || a < minInt || b < minInt {
// 		os.Exit(0) // return
// 	}
// 	switch op {
// 	case "+":
// 		c := a + b
// 		if (c > a) == (b > 0) {
// 			fmt.Printf("%d\n", c)
// 		}
// 	case "-":
// 		c := a - b
// 		if (c < a) == (b > 0) {
// 			fmt.Printf("%d\n", c)
// 		}
// 	case "*":
// 		c := a * b
// 		fmt.Printf("%d\n", c)
// 	case "/":
// 		if b == 0 {
// 			fmt.Printf("No division by 0\n")
// 			return
// 		}
// 		c := a / b
// 		fmt.Printf("%d\n", c)
// 	case "%":
// 		if b == 0 {
// 			fmt.Printf("No modulo by 0\n")
// 			return
// 		}
// 		c := a % b
// 		fmt.Printf("%d\n", c)
// 	default:
// 		return
// 	}
// }

// func Atoi(args string) int {
// 	s := 0
// 	temp := 1
// 	if args[0] == '-' {
// 		temp = -1
// 		args = args[1:]
// 	}

// 	if args[0] == '+' {
// 		args = args[1:]
// 	}

// 	for _, w := range args {
// 		if w < '0' || w > '9' {
// 			os.Exit(0)
// 		}
// 		s = s * 10
// 		s = s + int(w-48)
// 	}
// 	return s * temp
// }

// func main() {
// 	if len(os.Args) == 4 {
// 		args := os.Args[1:]
// 		a := Atoi(args[0])
// 		b := Atoi(args[2])
// 		if a > 0 && b > 0 && (a+b) < 0 || a < 0 && b < 0 && (a+b) > 0 {
// 			return // os.Exit(0)
// 		}
// 		switch args[1] {
// 		case "+":
// 			c := a + b
// 			if (c > a) == (b > 0) {
// 				fmt.Printf("%d\n", c)
// 			}
// 		case "-":
// 			c := a - b
// 			if (c < a) == (b > 0) {
// 				fmt.Printf("%d\n", c)
// 			}
// 		case "*":
// 			c := a * b
// 			if c/a == b {
// 				fmt.Printf("%d\n", c)
// 			}
// 		case "/":
// 			if b == 0 {
// 				fmt.Printf("No division by 0\n")
// 				return
// 			}
// 			fmt.Printf("%d\n", a/b)
// 		case "%":
// 			if b == 0 {
// 				fmt.Printf("No modulo by 0\n")
// 				return
// 			}
// 			fmt.Printf("%d\n", a%b)
// 		default:
// 			return
// 		}
// 	}
// }
// func Atoi(s string) int {
// 	var num int
// 	temp := 1

// 	if len(s) > 0 {
// 		if s[0] == '-' {
// 			temp = -1
// 			s = s[1:]
// 		}

// 		for _, w := range s {
// 			if w < '0' || w > '9' {
// 				os.Exit(0)
// 			}
// 			num = num*10 + int(w) - '0'
// 		}
// 	}
// 	return num * temp
// }

// func main() {
// 	if len(os.Args) != 4 {
// 		return
// 	}

// 	a := Atoi(os.Args[1])
// 	b := Atoi(os.Args[3])
// 	op := os.Args[2]

// 	maxInt := int(^uint(0) >> 1)
// 	minInt := -maxInt - 1

// 	if (a > 0 && b > 0 && a+b < 0) || a > maxInt || b > maxInt || a < minInt || b < minInt {
// 		os.Exit(0)
// 	}
// 	switch op {
// 	case "+":
// 		c := a + b
// 		fmt.Printf("%d\n", c)
// 	case "-":
// 		c := a - b
// 		fmt.Printf("%d\n", c)
// 	case "*":
// 		c := a * b
// 		fmt.Printf("%d\n", c)
// 	case "/":
// 		if b == 0 {
// 			fmt.Printf("No division by 0\n")
// 			return
// 		}
// 		c := a / b
// 		fmt.Printf("%d\n", c)
// 	case "%":
// 		if b == 0 {
// 			fmt.Printf("No modulo by 0\n")
// 			return
// 		}
// 		c := a % b
// 		fmt.Printf("%d\n", c)
// 	}
// }
