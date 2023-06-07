// strlen
// Instructions
// Write a function that counts the characters of a string and that returns that count.
// Expected function
// func StrLen(str string) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	piscine ".."
// )

// func main() {
// 	str := "Hello World!"
// 	nb := piscine.StrLen(str)
// 	fmt.Println(nb)
// }
// And its output :

// student@ubuntu:~/piscine-go/test$ go build
// student@ubuntu:~/piscine-go/test$ ./test
// 12
// student@ubuntu:~/piscine-go/test$

// allowedFunctions
// 0	"--allow-builtin"
package main

import "fmt"

func main() {
	v := StrLen("Hello World!")
	fmt.Println(v)
}

func StrLen(s string) int {
	return len([]rune(s))
	// return len(s)
	// arr := []rune(s)

	// return len(arr)
	// var arr []rune

	// for _, w := range s {
	// 	arr = append(arr, w)
	// }
	// return len(arr)
}

// func StrLen(s string) int {
// 	// var arr []rune
// 	// // for _, w := range s {
// 	// // 	arr = append(arr, w)
// 	// // }
// 	return len([]rune(s))
// }

// func StrLen(s string) int {
// 	return len([]rune(s))

// 	// var arr []rune
// 	// for _, v := range s {
// 	// 	arr = append(arr, v)
// 	// }
// 	// return len(arr)
// }

// func StrLen(s string) int {
// 	return len([]rune(s))
// }

// package main

// import "fmt"

// func main() {
// 	l := StrLen("Hello World!")
// 	fmt.Println(l)
// }

// func StrLen(s string)int{
// 	return len([]rune(s))
// }
