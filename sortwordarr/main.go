// sortwordarr
// Instructions

// Write a function SortWordArr that sorts by ascii (in ascending order) a string slice.
// Expected function

// func SortWordArr(a []string) {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
// 	piscine.SortWordArr(result)

// 	fmt.Println(result)
// }

// And its output :

// $ go run .
// [1 2 3 A B C a b c]
// $
// 0	"--allow-builtin"

package main

import "fmt"

func main() {
	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(result)

	fmt.Println(result)
}

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// func SortWordArr(a []string) {
// 	for i := 0; i < len(a); i++ {
// 		for j := i + 1; j < len(a); j++ {
// 			if a[j] < a[i] {
// 				a[i], a[j] = a[j], a[i]
// 			}
// 		}
// 	}
// }
