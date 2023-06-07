// addprimesum
// Instructions

// Write a program that takes a positive integer as argument and displays the sum of all prime numbers inferior or equal to it followed by a newline ('\n').

//     If the number of arguments is different from 1, or if the argument is not a positive number, the program displays 0 followed by a newline.

// Usage

// $ go run . 5
// 10
// $ go run . 7
// 17
// $ go run . -2
// 0
// $ go run . 0
// 0
// $ go run .
// 0
// $ go run . 5 7
// 0
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.*"
// 2	"--allow-builtin"

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || Atoi(args[0]) <= 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	nb := Atoi(args[0])
	str := Addprimesum(nb)

	res := Itoa(str)
	for _, w := range res {
		z01.PrintRune(w)
	}

	z01.PrintRune('\n')
}

func Addprimesum(nb int) int {
	for nb <= 1 {
		return 0
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return Addprimesum(nb - 1)
		}
	}
	return nb + Addprimesum(nb-1)
}

func Atoi(args string) int {
	s := 0
	temp := 1
	if args[0] == '-' {
		temp = -1
		args = args[1:]
	} else if args[0] == '+' {
		args = args[1:]
	}
	for _, w := range args {
		if w < '0' || w > '9' {
			os.Exit(0)
		}
		s = s * 10
		s = s + int(w-48)
	}
	return s * temp
}

func Itoa(n int) string {
	res := ""
	if n == 0 {
		res = ""
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

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	n := Atoi(args[0])
// 	fmt.Print(Addprimesum(n))
// 	z01.PrintRune('\n')
// }

// func Addprimesum(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}

// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			return Addprimesum(nb - 1)
// 		}
// 	}
// 	return nb + Addprimesum(nb-1) // 4+f(3)   f(2)=2+f(1)=
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
// 	for _, v := range args {
// 		if v < '0' || v > '9' {
// 			os.Exit(0)
// 		}
// 		s = s * 10
// 		s = s + int(v-48)
// 	}
// 	return s * temp
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 || Atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	n := Atoi(args[0])
// 	sum := 0
// 	for i := n; i >= 2; i-- {
// 		if Isprime(i) { //0+7 + 5 + 3 + 2 = 17
// 			sum = i + sum
// 		}
// 	}
// 	for _, v := range Itoa(sum) {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

// func Isprime(nb int) bool {
// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
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
// 	for _, v := range args {
// 		if v < '0' || v > '9' {
// 			os.Exit(0)
// 		}
// 		s = s * 10
// 		s = s + int(v-48)
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

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 || Atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	n := Atoi(args[0])
// 	sum := 0
// 	for i := n; i >= 2; i-- {
// 		if Isprime(i) {
// 			sum = sum + i
// 		}
// 	}

// 	for _, w := range Itoa(sum) {
// 		z01.PrintRune(w)
// 	}
// 	z01.PrintRune('\n')
// }

// func Isprime(n int) bool {
// 	for i := 2; i <= n/2; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	args := os.Args[1:]

// 	if len(args) != 1 || Atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}

// 	n := Atoi(args[0])

// 	sum := 0

// 	for n := n; n >= 2; n-- { // 7
// 		if Isprime(n) {
// 			sum = sum + n // 0 + 7 + 5 + 3 + 2
// 		}
// 	}
// 	for _, v := range Itoa(sum) {
// 		z01.PrintRune(v) // 7
// 	}
// 	z01.PrintRune('\n')
// }

// func Isprime(n int) bool {
// 	for i := 2; i <= n/2; i++ { // 7

// 		if n%i == 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 || Atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	n := (Atoi(os.Args[1]))
// 	sum := 0
// 	for i := n; i >= 2; i-- {
// 		if Isprime(i) {
// 			sum = sum + i
// 		}
// 	}
// 	for _, v := range Itoa(sum) {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune(10)
// }

// func Isprime(n int) bool {
// 	for a := 2; a <= n/2; a++ {
// 		if n%a == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func Atoi(args string) int {
// 	s := 0
// 	temp := 1
// 	for _, v := range args {
// 		if v == '-' {
// 			temp = -1
// 			args = args[1:]
// 		}
// 	}
// 	for _, w := range args {
// 		if w < '0' && w > '9' {
// 			os.Exit(0)
// 		}
// 		s = s * 10
// 		s = s + int(w-48)

// 	}
// 	return s * temp
// }

// func Itoa(n int) string {
// 	res := ""

// 	if n == 0 {
// 		res = ""
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

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 || Atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 		return
// 	}
// 	n := (Atoi(os.Args[1]))
// 	sum := 0
// 	for i := n; i >= 2; i-- {
// 		if Isprime(i) {
// 			sum = sum + i
// 		}
// 	}
// 	for _, v := range Itoa(sum) {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune(10)
// }

// func Isprime(n int) bool {
// 	for a := 2; a <= n/2; a++ {
// 		if n%a == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 || atoi(args[0]) <= 1 {
// 		z01.PrintRune('0')
// 		z01.PrintRune('\n')
// 	} else {
// 		for _, w := range itoa(addprimesum(atoi(args[0]))) {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func itoa(n int) string {
// 	res := ""
// 	for n != 0 {
// 		res = string(rune(n%10)+48) + res
// 		n = n / 10
// 	}
// 	return res
// }

// func addprimesum(n int) int {
// 	res := 0
// 	for i := 2; i <= n; i++ {
// 		if isPrime(i) {
// 			res += i
// 		}
// 	}
// 	return res
// }

// func isPrime(n int) bool {
// 	for i := 2; i <= n/2; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func atoi(s string) int {
// 	res := 0
// 	temp := 1
// 	if s[0] == '-' {
// 		temp = -1
// 		s = s[1:]
// 	}
// 	for _, w := range s {
// 		if w >= '0' && w <= '9' {
// 			res = res*10 + int(w-48)
// 		}
// 	}
// 	return res * temp
// }
