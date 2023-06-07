// split
// Instructions

// Write a function that receives a string and a separator and returns a slice of strings that results of splitting the string s by the separator sep.
// Expected function

// func Split(s, sep string) []string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "HelloHAhowHAareHAyou?"
// 	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
// }

// And its output :

// $ go run .
// []string{"Hello", "how", "are", "you?"}
// $
// allowedFunctions
// 0	"make"
// 1	"--allow-builtin"

package main

import (
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	slice := 0
	arr := []string{}

	for w := range s {
		if w < len(s)-len(sep) {
			if s[w:w+len(sep)] == sep {
				arr = append(arr, s[slice:w])
				slice = w + len(sep)
			}
		}
		// if w == len(s)-1 {
		// 	arr = append(arr, s[slice:])
		// }
		if w == len(s)-1 {
			arr = append(arr, s[slice:])
		}

	}

	return arr
}

// func Split(s, sep string) []string {
// 	slice := 0
// 	result := []string{}
// 	for i := range s {
// 		if i < len(s)-len(sep) {
// 			if s[i:i+len(sep)] == sep {
// 				result = append(result, s[slice:i])
// 				slice = i + len(sep)
// 			}
// 		}
// 		if i == len(s)-1 {
// 			result = append(result, s[slice:])
// 		}
// 	}
// 	return result
// }

// [hello, how, are, you?]
