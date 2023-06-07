// piglatin
// Instructions

// Write a program that transforms a string passed as argument in its Pig Latin version.

// The rules used by Pig Latin are as follows:

//     If a word begins with a vowel, just add "ay" to the end.
//     If it begins with a consonant, then we take all consonants before the first vowel and we put them on the end of the word and add "ay" at the end.
//     Only the latin vowels will be considered as vowel(aeiou).

// If the word has no vowels, the program should print "No vowels".

// If the number of arguments is different from one, the program prints nothing.
// Usage

// $ go run .
// $ go run . pig | cat -e
// igpay$
// $ go run . Is | cat -e
// Isay$
// $ go run . crunch | cat -e
// unchcray$
// $ go run . crnch | cat -e
// No vowels$
// $ go run . something else | cat -e
// $
// allowedFunctions
// 0	"os.Args"
// 1	"github.com/01-edu/z01.PrintRune"
// 2	"--allow-builtin"
package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	s := ""
// 	str := args[0]
// 	for i, v := range str {
// 		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
// 			s = str[i:] + str[:i] + "ay"
// 			break
// 		}
// 	}
// 	if s == "" {
// 		s = "No vowels"
// 	}
// 	for _, v := range s {
// 		z01.PrintRune(v)
// 	}
// 	z01.PrintRune('\n')
// }

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s := ""
	str := args[0]
	for _, w := range str {
		if w != 'U' || w != 'I' || w != 'A' || w != 'O' || w != 'E' || w != 'u' || w != 'i' || w != 'a' || w != 'o' || w != 'e' {
			s = "No vowels"
		}
	}
	for v := range str {
		if str[v] == 'U' || str[v] == 'I' || str[v] == 'A' || str[v] == 'O' || str[v] == 'E' || str[v] == 'u' || str[v] == 'i' || str[v] == 'a' || str[v] == 'o' || str[v] == 'e' {
			s = str[v:] + str[:v] + "ay"
			break
		}
	}

	for _, x := range s {
		z01.PrintRune(x)
	}
	z01.PrintRune('\n')
}

// func main() {
// 	if len(os.Args) == 2 {
// 		args := os.Args[1:]

// 		s := ""

// 		str := args[0]

// 		for _, w := range str {
// 			if w != 'U' || w != 'U' || w != 'I' || w != 'A' || w != 'O' || w != 'E' || w != 'u' || w != 'i' || w != 'o' || w != 'e' {
// 				s = "No vowels"
// 			}
// 		}

// 		for v := range str {
// 			if str[v] == 'U' || str[v] == 'U' || str[v] == 'I' || str[v] == 'A' || str[v] == 'O' || str[v] == 'E' || str[v] == 'u' || str[v] == 'i' || str[v] == 'o' || str[v] == 'e' {
// 				s = str[v:] + str[:v] + "ay"
// 			}
// 		}

// 		for _, x := range s {
// 			z01.PrintRune(x)
// 		}
// 		z01.PrintRune('\n')
// 	} else if len(os.Args) > 2 {
// 		z01.PrintRune('\n')
// 	}
// }

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	}
// 	s := ""
// 	str := args[0]
// 	for _, w := range str {
// 		if w != 'U' || w != 'I' || w != 'A' || w != 'O' || w != 'E' || w != 'u' || w != 'i' || w != 'a' || w != 'o' || w != 'e' {
// 			s = "No vowels"
// 		}
// 	}
// 	for i := range str {
// 		if str[i] == 'U' || str[i] == 'I' || str[i] == 'A' || str[i] == 'O' || str[i] == 'E' || str[i] == 'u' || str[i] == 'i' || str[i] == 'a' || str[i] == 'o' || str[i] == 'e' {
// 			s = str[i:] + str[:i] + "ay"
// 		}
// 	}
// 	fmt.Println(s)
// }
