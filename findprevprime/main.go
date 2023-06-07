// findprevprime
// Instructions
// Write a function that returns the first prime number that is equal or inferior to the int passed as parameter.

// If there are no primes inferior to the int passed as parameter the function should return 0.

// Expected function
// func FindPrevPrime(nb int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"

// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.FindPrevPrime(5))
// 	fmt.Println(piscine.FindPrevPrime(4))
// }
// And its output :

// $ go run .
// 5
// 3
// $

package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	if nb <= 1 {
		return 0
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			nb = FindPrevPrime(nb - 1)
		}
	}
	return nb
}

// func PrimeSum(nb int) int {
// 	s := 0
// 	if nb == 1 {
// 		return 0
// 	}
// 	if nb == 2 {
// 		return 2
// 	}
// 	p := 0
// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			p = 1
// 		}
// 	}
// 	if p == 0 {
// 		s = s + FindPrevPrime(nb-1)
// 	}
// 	return s
// }

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	for i := 2; i < nb/2; i++ {
// 		if nb%2 == 0 {
// 			nb = FindPrevPrime(nb - 1)
// 		}
// 	}
// 	return nb
// }

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}

// 	for i := 2; i <= nb/2; i++ {
// 		if nb%i == 0 {
// 			nb = FindPrevPrime(nb - 1)
// 		}
// 	}

// 	return nb
// }

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	for i := 2; i <= nb/2; i++ { // 2 /5.5     3/5.5     4/5.5
// 		if nb%2 == 0 { // 5+5+(1)/ 2+2+(1)/  2(0)
// 			return FindPrevPrime(nb - 1)
// 		}
// 	}
// 	return nb
// }

// package main

// import "fmt"

// func FindPrevPrime(nb int) int {
// 	if nb <= 1 {
// 		return 0
// 	}
// 	for a := nb; ; a-- {
// 		if Isprime(a) {
// 			return a
// 		}
// 	}
// }

// func Isprime(nb int) bool {
// 	for i := 2; i < nb; i++ {
// 		if nb%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(FindPrevPrime(2))
// 	fmt.Println(FindPrevPrime(4))
// }
