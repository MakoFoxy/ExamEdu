// searchreplace
// Instructions
// Write a program that takes 3 arguments, the first argument is a string in which a letter (the 2nd argument) will be replaced by another one (the 3rd argument).

// If the number of arguments is different from 3, the program displays nothing.

// If the second argument is not contained in the first one (the string) then the program rewrites the string followed by a newline ('\n').

// Usage
// $ go run . "hella there" "a" "o"
// hello there
// $ go run . "hallo thara" "a" "e"
// hello there
// $ go run . "abcd" "z" "l"
// abcd
// $ go run . "something" "a" "o" "b" "c"
// $
// allowedFunctions
// 0	"github.com/01-edu/z01.PrintRune"
// 1	"os.Args"
// 2	"--allow-builtin"

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	res := ""

	if len(args) != 3 {
		return
	} else {
		for _, w := range args[0] {
			for _, v := range args[1] {
				if w == v {
					res = res + args[2]
				} else {
					res = res + string(w)
				}
			}
		}
		for _, w := range res {
			z01.PrintRune(w)
		}
		z01.PrintRune('\n')
	}
}

// func main() {
// 	args := os.Args[1:]
// 	outp := ""

// 	if len(args) != 3 {
// 		return
// 	} else {
// 		for _, v := range args[0] {
// 			if string(v) == args[1] {
// 				outp = outp + args[2]
// 			} else {
// 				outp = outp + string(v)
// 			}
// 		}
// 		// for i := 0; i < len(args[0]); i++ {
// 		// 	for j := 0; j < len(args[1]); j++ {
// 		// 		if args[0][i] == args[1][j] {
// 		// 			outp = outp + string(args[2])
// 		// 		} else {
// 		// 			outp = outp + string(args[1])
// 		// 		}
// 		// 	}
// 		// }
// 		for _, w := range outp {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	outp := ""
// 	if len(args) != 3 {
// 		return
// 	} else {
// 		for _, v := range args[0] {
// 			if string(v) == args[1] {
// 				outp = outp + args[2]
// 			} else {
// 				outp = outp + string(v)
// 			}
// 		}

// 		for _, w := range outp {
// 			z01.PrintRune(w)
// 		}
// 		z01.PrintRune('\n')
// 	}
// }

// for _, v := range args[0] {
// 	if string(v) == args[1] {
// 		outp = append(outp, v)
// 	}
// }
// for _, z := range args[2] {
// 	if string(z) != args[1] {
// 		arr = append(arr, z)
// 	}
// }
// for _, x := range outp {
// 	z01.PrintRune(x)
// 	arr2 = append(arr2, outp)
// }
// for _, y := range arr {
// 	z01.PrintRune(y)
// 	arr2 = append(arr2, arr)

// }
// for _, u := range arr2 {
// 	fmt.Println(u)
// }
//}

// func main() {
// 	args := os.Args[1:]
// 	result := ""

// 	if len(args) != 3 {
// 		return
// 	} else {
// 		for _, w := range args[0] {
// 			if string(w) == args[1] {
// 				result = result + args[2]
// 			} else {
// 				result = result + string(w)
// 			}
// 		}
// 		for _, w := range result {
// 			z01.PrintRune(w)
// 		}
// 	}
// 	z01.PrintRune('\n')
// }
// package main

// import (
// "fmt"
// "os"

// "github.com/01-edu/z01"
// )

// func main() {
// args := os.Args[1:]
// result := ""
// res := ""
// arg := len(args)
// arr := []rune{}
// arr2 := []rune{}
// arr3 := [][]rune{}

// if arg != 3 {
// return
// } else {
// for i := 0; i < len(args[0]); i++ {
// for j := 0; j < len(args[1]); j++ {
// if args[0][i] == args[1][j] {
// for x := 0; x < len(args[2]); x++ {
// if args[2][x] != args[1][j] {
// result = args[0] + result
// res = args[2] + res
// }
// }
// }
// }
// }
// for _, v := range result {
// if result != res {
// for _, y := range res {
// if y != v {
// z01.PrintRune(y)
// arr2 = append(arr2, y)
// }
// }
// z01.PrintRune(v)
// arr = append(arr, v)
// }
// }
// for _, n := range arr {
// z01.PrintRune(n)
// arr3 = append(arr3, arr)
// }
// for _, t := range arr2 {
// z01.PrintRune(t)
// arr3 = append(arr3, arr2)
// }

// for _, k := range arr3 {
// fmt.Println(k)
// }
// }
// z01.PrintRune('\n')
// }
