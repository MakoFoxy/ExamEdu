// max
// Instructions
// Write a function, Max, that returns the maximum value in a slice of integers.

// Expected function
// func Max(arr []int) int {

// }
// Usage
// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	piscine ".."
// )

// func main() {
// 	arrInt := []int{23, 123, 1, 11, 55, 93}
// 	max := piscine.Max(arrInt)
// 	fmt.Println(max)
// }

// And its output :

// student@ubuntu:~/piscine-go/test$ go build
// student@ubuntu:~/piscine-go/test$ ./test
// 123
// student@ubuntu:~/piscine-go/test$
// allowedFunctions
// 0	"--allow-builtin"

package main

import (
	"fmt"
)

func main() {
	arrMax := []int{45, 785, 654, 8521, 546, 854, 89, 458}
	max := Max(arrMax)
	fmt.Println(max)
}

func Max(s []int) int {
	max := 0
	for _, w := range s {
		if max < w {
			max = w
		}
	}
	return max
}

// func Max(s []int) int {
// 	max := 0
// 	for _, w := range s {
// 		if max < w {
// 			max = w
// 		}
// 	}
// 	return max
// }

// func Max(arr []int) int {
// 	//var qwe []int
// 	max := 0
// 	for _, w := range arr {
// 		if max < w {
// 			max = w
// 			//qwe = append(qwe, max)
// 		}
// 	}
// 	return max

// 	// x := 0
// 	// for _, x := range qwe {
// 	// 	return x
// 	// }
// 	// return x
// }

// func Max(s []int) int {
// 	max := 0
// 	for _, v := range s {
// 		if max < v {
// 			max = v
// 		}
// 	}
// 	return max
// }

// func Max(s []int) int {
// 	max := 0
// 	for v := range s {
// 		if max < s[v] {
// 			max = s[v]
// 		}
// 	}
// 	return max
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	arrInt := []int{23, 123, 1, 11, 55, 93}
// 	max := Max(arrInt)
// 	fmt.Println(max)
// }

// func Max(a []int) int {
// 	max := 0
// 	for i := range a {
// 		if max < a[i] { // сравинивает
// 			max = a[i] // записывает в макс
// 		}
// 	}
// 	return max
// }
