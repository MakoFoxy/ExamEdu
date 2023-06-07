// atoibase
// Instructions

// Write a function that takes two arguments:

//     s: a numeric string in a given base.
//     base: a string representing all the different digits that can represent a numeric value.

// And return the integer value of s in the given base.

// If the base is not valid it returns 0.

// Validity rules for a base :

//     A base must contain at least 2 characters.
//     Each character of a base must be unique.
//     A base should not contain + or - characters.

// String number must contain only elements that are in base.

// Only valid string numbers will be tested.

// The function does not have to manage negative numbers.
// Expected function

// func AtoiBase(s string, base string) int {

// }

// Usage

// Here is a possible program to test your function :

package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

// And its output :

// $ go run .
// 125
// 125
// 125
// 125
// 0
// $

func AtoiBase(s, base string) int {
	m := make(map[rune]int)

	for i, w := range base {
		if w == '+' || w == '-' {
			return 0
		}
		m[w] = i
	}

	var n int

	for _, v := range s {
		n = n*len(base) + m[v]
	}
	return n
}

// func AtoiBase(s, base string) int {
// 	m := make(map[rune]int)
// 	for i, v := range base {
// 		if v == '+' || v == '-' {
// 			return 0
// 		}
// 		m[v] = i
// 	}
// 	var num int
// 	for _, v2 := range s {
// 		num = num*len(base) + m[v2]
// 	}
// 	return num
// }

// func AtoiBase(s, base string) int {
// 	if len(base) < 2 {
// 		return 0
// 	}
// 	for i, v := range base {
// 		for j, w := range base {
// 			if w == v && i != j || v == '-' || v == '+' {
// 				return 0
// 			}
// 		}
// 	}
// 	sum := 0
// 	for i := 0; i < len(s); i++ {
// 		sum += Index(base, string(s[i])) * Pow(len(base), len(s)-i-1)
// 	}
// 	return sum
// }

// func Index(base, nbr string) int {
// 	for i := range base {
// 		if string(base[i]) == nbr {
// 			return i
// 		}
// 	}
// 	return 0
// }

// func Pow(nbr, power int) int {
// 	result := 1
// 	for i := 0; i < power; i++ {
// 		result = result * nbr
// 	}
// 	return result
// }
