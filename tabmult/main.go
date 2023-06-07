// tabmult
// Instructions
// Write a program that displays a number's multiplication table.

// The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.
// Usage
// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.*"
// 2	"strconv.Atoi"
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
		n, err := strconv.Atoi(args[0])
		if err == nil {
			for i := 1; i < 10; i++ {
				res := Itoa(i) + " x " + Itoa(n) + " = " + Itoa(i*n)

				for _, w := range res {
					z01.PrintRune(w)
				}
				z01.PrintRune('\n')
			}
		}
	} else {
		// z01.PrintRune('\n')
		return
	}
}

func Itoa(n int) string {
	res := ""

	if n == 0 {
		res = "0"
	} else if n < 0 {
		for n != 0 {
			res = string(rune(-(n%10 + 48))) + res
			n = n / 10
		}
		res = string('-') + res
	} else {
		for n != 0 {
			res = string(rune(n%10+48)) + res
			n = n / 10
		}
	}

	return res
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

// 	return s * temp
// }

// func Itoa(n int) string {
// 	res := ""
// 	if n == 0 {
// 		res = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			res = string(rune(-(n%10 + 48))) + res
// 			n = n / 10
// 		}
// 		res = string('-') + res
// 	} else {
// 		for n != 0 {
// 			res = string(rune(n%10+48)) + res
// 			n = n / 10
// 		}
// 	}

// 	return res
// }

// func Itoa(n int) string {
// 	result := ""

// 	if n == 0 {
// 		result = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			result = string(rune(-(n%10 + 48))) + result
// 			n = n / 10
// 		}
// 		result = string('-') + result
// 	} else {
// 		for n != 0 {
// 			result = string(rune(n%10+48)) + result
// 			n = n / 10
// 		}
// 	}
// 	return result
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 {
// 		return
// 	}
// 	s := Atoi(args[0])
// 	for i := 1; i < 10; i++ {
// 		result := Itoa(i) + " x " + Itoa(s) + " = " + Itoa(i*s)
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }
// func Atoi(args string) int {
// 	s := 0
// 	temp := 1

// 	for _, w := range args {
// 		if w == '-' {
// 			temp = -1
// 			args = args[1:]
// 		}
// 	}

// 	for i := range args {
// 		s = s * 10
// 		s = s + int(args[i]-48)
// 	}
// 	return s * temp
// }

// func Itoa(n int) string {
// 	result := ""

// 	if n == 0 {
// 		result = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			result = string(rune(-(n%10 + 48))) + result
// 			n = n / 10
// 		}
// 		result = string('-') + result
// 	} else {
// 		for n != 0 {
// 			result = string(rune(n%10+48)) + result
// 			n = n / 10
// 		}
// 	}
// 	return result
// }

// func Itoa(n int) string {
// 	result := ""
// 	if n == 0 {
// 		result = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			//-19
// 			result = string(rune(-(n%10))+48) + result
// 			n = n / 10
// 		}
// 		result = string('-') + result
// 	} else {
// 		for n != 0 {
// 			result = string(rune(n%10)+48) + result
// 			n = n / 10
// 		}
// 	}
// 	return result
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

// func Itoa(n int) string {
// 	result := ""
// 	if n == 0 {
// 		result = "0"
// 	} else if n < 0 {
// 		for n != 0 {
// 			//-19
// 			result = string(rune(-(n%10))+48) + result
// 			n = n / 10
// 		}
// 		result = string('-') + result
// 	} else {
// 		for n != 0 {
// 			result = string(rune(n%10)+48) + result
// 			n = n / 10
// 		}
// 	}
// 	return result
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	s, err := strconv.Atoi(args[0])
// 	if err == nil {
// 		for i := 1; i < 10; i++ {
// 			result := strconv.Itoa(i) + " x " + strconv.Itoa(s) + " = " + strconv.Itoa(i*s)
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	s, err := strconv.Atoi(args[0])

// 	if err == nil {
// 		for i := 1; i < 10; i++ {
// 			result := strconv.Itoa(i) + " x " + strconv.Itoa(s) + " = " + strconv.Itoa(i*s)
// 			for _, w := range result {
// 				z01.PrintRune(w)
// 			}
// 			z01.PrintRune('\n')
// 		}
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else {
// 		number, _ := strconv.Atoi(args[0])
// 		tabmul(number)
// 	}
// }

// func tabmul(z int) {
// 	for i := 1; i < 10; i++ {
// 		result := itoa(i) + " x " + itoa(z) + " = " + itoa(i*z)
// 		for _, w := range result {
// 			z01.PrintRune(rune(w))
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func itoa(j int) string {
// 	result := ""
// 	for j != 0 {
// 		result = string(rune(j%10)+48) + result
// 		j = j / 10
// 	}
// 	return result
// }
