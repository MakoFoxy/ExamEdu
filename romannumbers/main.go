// romannumbers
// Instructions

// Write a program called rn. The objective is to convert a number, given as an argument, into a roman number and print it with its roman number calculation.

// The program should have a limit of 4000. In case of an invalid number, for example "hello" or 0 the program should print ERROR: cannot convert to roman digit.

// Roman Numerals reminder:
// I 	1
// V 	5
// X 	10
// L 	50
// C 	100
// D 	500
// M 	1000

// For example, the number 1732 would be denoted MDCCXXXII in Roman numerals. However, Roman numerals are not a purely additive number system. In particular, instead of using four symbols to represent a 4, 40, 9, 90, etc. (i.e., IIII, XXXX, VIIII, LXXXX, etc.), such numbers are instead denoted by preceding the symbol for 5, 50, 10, 100, etc., with a symbol indicating subtraction. For example, 4 is denoted IV, 9 as IX, 40 as XL, etc.

// The following table gives the Roman numerals for the first few positive integers.
// 1 	I 	11 	XI 	21 	XXI
// 2 	II 	12 	XII 	22 	XXII
// 3 	III 	13 	XIII 	23 	XXIII
// 4 	IV 	14 	XIV 	24 	XXIV
// 5 	V 	15 	XV 	25 	XXV
// 6 	VI 	16 	XVI 	26 	XXVI
// 7 	VII 	17 	XVII 	27 	XXVII
// 8 	VIII 	18 	XVIII 	28 	XXVIII
// 9 	XIX 	19 	XIX 	29 	XXIX
// 10 	X 	20 	XX 	30 	XXX
// Usage

// $ go run . hello
// ERROR: cannot convert to roman digit
// $ go run . 123
// C+X+X+I+I+I
// CXXIII
// $ go run . 999
// (M-C)+(C-X)+(X-I)
// CMXCIX
// $ go run . 3999
// M+M+M+(M-C)+(C-X)+(X-I)
// MMMCMXCIX
// $ go run . 4000
// ERROR: cannot convert to roman digit
// $

// allowedFunctions
// 0	"os.Args"
// 1	"fmt.Printf"
// 2	"--allow-builtin"

package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}
	num := Atoi(args[0])
	// if !err {
	if num <= 0 || num >= 4000 {
		fmt.Printf("ERROR: cannot convert to roman digit")
		fmt.Printf("\n")
		return
	}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	arns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thns := []string{"", "M", "MM", "MMM"}

	res := thns[num/1000] + arns[num%1000/100] + tens[num%100/10] + ones[num%10]

	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
	arns1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
	thns1 := []string{"", "M+", "M+M+", "M+M+M+"}

	res1 := thns1[num/1000] + arns1[num%1000/100] + tens1[num%100/10] + ones1[num%10]

	fmt.Printf(res1[:len(res1)-1])
	fmt.Printf("\n")
	fmt.Printf(res)
	fmt.Printf("\n")

	// } else {
	// 	fmt.Printf("ERROR: cannot convert to roman digit")
	// 	fmt.Printf("\n")
	// 	return
	// }
}

// if len(os.Args) > 2 || len(os.Args) < 2 {
// 	return
// }
// num := Atoi(os.Args[1])

func Atoi(args string) int {
	s := 0
	temp := 1

	if args[0] == '-' {
		temp = -1
		args = args[1:]
	}
	for _, v := range args {
		if v < '0' || v > '9' {
			return 0
			// fmt.Printf("ERROR: cannot convert to roman digit\n")
			// os.Exit(0)
			//return -1, true
		}
		s = s * 10
		s = s + int(v-48)
	}
	return s * temp //, false
}

// func Atoi(args string) int {
// 	s := 0
// 	temp := 1

// 	if args[0] == '-' {
// 		temp = -1
// 		args = args[1:]

// 	}
// 	for i := 0; i < len(args); i++ {
// 		if !(args[i] >= 0 && args[0] <= 9) {
// 			return 0
// 		}
// 		s = s * 10
// 		s = s + int(args[i]-48)
// 	}
// 	return s * temp
// }

// func main() {
// 	if len(os.Args) > 2 || len(os.Args) < 2 {
// 		return
// 	}
// 	args := os.Args[1]

// 	num := Atoi(args)

// 	if num <= 0 || num >= 4000 {
// 		fmt.Printf("ERROR: cannot convert to roman digit")
// 		fmt.Printf("\n")
// 		return
// 	}

// 	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
// 	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
// 	arns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
// 	thsn := []string{"", "M", "MM", "MMM"}

// 	res := thsn[num/1000] + arns[num%1000/100] + tens[num%100/10] + ones[num%10]

// 	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
// 	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
// 	arns1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
// 	thsn1 := []string{"", "M+", "M+M+", "M+M+M+"}

// 	res1 := thsn1[num/1000] + arns1[num%1000/100] + tens1[num%100/10] + ones1[num%10]

// 	fmt.Printf(res1[:len(res1)-1])

// 	fmt.Printf("\n")
// 	fmt.Printf(res)

// 	fmt.Printf("\n")
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
// 	for i := 0; i < len(args); i++ {
// 		s = s * 10
// 		s = s + int(args[i]-48)
// 	}
// 	return s * temp
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

// func main() {
// 	if len(os.Args) < 2 || len(os.Args) > 2 {
// 		return
// 	}
// 	args := os.Args[1]
// 	num := Atoi(args)

// 	if num <= 0 || num >= 4000 {
// 		fmt.Printf("ERROR: cannot convert to roman digit")
// 		fmt.Printf("\n")
// 		return
// 	}

// 	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
// 	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
// 	arns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
// 	thsn := []string{"", "M", "MM", "MMM"}
// 	res := thsn[num/1000] + arns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]

// 	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
// 	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
// 	arns1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
// 	thsn1 := []string{"", "M+", "M+M+", "M+M+M+"}

// 	res1 := thsn1[num/1000] + arns1[(num%1000)/100] + tens1[(num%100)/10] + ones1[num%10]

// 	fmt.Printf(res1[:len(res1)-1])
// 	fmt.Printf("\n")
// 	fmt.Printf(res)
// 	fmt.Printf("\n")
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

// func Atoi(args string) int {
// 	s := 0
// 	for i := 0; i < len(args); i++ {
// 		s = s * 10
// 		s = s + int(args[i]-'0')
// 	}
// 	return s
// }

// func Atoi(s string) int {
// 	n := 0

// 	for i := 0; i < len(s); i++ {
// 		n = n * 10
// 		n = n + int(s[i]-'0')
// 	}
// 	return n
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	s := os.Args[1] // 123
// 	// fmt.Println(s[0])
// 	// if os.Args[1] <= "0" {
// 	// 	fmt.Printf("ERROR: cannot convert to roman digit")
// 	// 	fmt.Printf("\n")
// 	// 	return
// 	// }
// 	num := (Atoi(s))

// 	if num <= 0 || num >= 4000 {
// 		fmt.Printf("ERROR: cannot convert to roman digit")
// 		fmt.Printf("\n")

// 		return
// 	}

// 	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
// 	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
// 	rns := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
// 	thns := []string{"", "M", "MM", "MMM"}
// 	res := thns[num/1000] + rns[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]

// 	ones1 := []string{"", "I+", "I+I+", "I+I+I+", "(V-I)+", "V+", "V+I+", "V+I+I+", "V+I+I+I+", "(X-I)+"}
// 	tens1 := []string{"", "X+", "X+X+", "X+X+X+", "(L-X)+", "L+", "L+X+", "L+X+X+", "L+X+X+X+", "(C-X)+"}
// 	rns1 := []string{"", "C+", "C+C+", "C+C+C+", "(D-C)+", "D+", "D+C+", "D+C+C+", "D+C+C+C+", "(M-C)+"}
// 	thns1 := []string{"", "M+", "M+M+", "M+M+M+"}

// 	res1 := thns1[num/1000] + rns1[(num%1000)/100] + tens1[(num%100)/10] + ones1[num%10]

// 	fmt.Printf(res1[:len(res1)-1])
// 	fmt.Printf("\n")
// 	fmt.Printf(res)
// 	fmt.Printf("\n")
// }

// func Atoi(s string) int {
// 	n := 0

// 	for i := 0; i < len(s); i++ {
// 		n = n * 10
// 		n = n + int(s[i]-'0')
// 	}
// 	return n
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// type roman struct {
// 	num int
// 	rm  string
// 	wr  string
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr := Atoi(os.Args[1])
// 		arr := []roman{
// 			{1000, "M", "M+"},
// 			{900, "CM", "(M-C)+"},
// 			{500, "D", "D+"},
// 			{400, "CD", "(D-C)+"},
// 			{100, "C", "C+"},
// 			{90, "XC", "(C-X)+"},
// 			{50, "L", "L+"},
// 			{40, "XL", "(L-X)+"},
// 			{10, "X", "X+"},
// 			{9, "IX", "X-I+"},
// 			{5, "V", "V+"},
// 			{4, "IV", "V-I+"},
// 			{1, "I", "I+"},
// 		}
// 		if nbr == 0 || nbr >= 4000 {
// 			fmt.Printf("ERROR: can not convert to roman digit")
// 		} else {
// 			ans := ""
// 			res := ""
// 			for i := 0; i <= len(arr); i++ {
// 				if nbr/arr[i].num != 0 {
// 					m := nbr / arr[i].num
// 					for j := 0; j < m; j++ {
// 						ans += arr[i].wr
// 						res += arr[i].wr
// 					}
// 					nbr = nbr % arr[i].num
// 				}
// 			}
// 			fmt.Printf(ans[:len(ans)-1])
// 			fmt.Printf("\n")
// 			fmt.Printf(res)
// 		}
// 		fmt.Printf("\n")

// 	}
// }

// func Atoi(s string) int {
// 	num := 0
// 	for _, k := range s {
// 		if k >= '0' && k <= '9' {
// 			num = num*10 + int(k-48)
// 		} else {
// 			return 0
// 		}
// 	}
// 	return num
// }

// type roman struct {
// 	num int
// 	rm  string
// 	wr  string
// }

// func main() {
// 	if len(os.Args) == 2 {
// 		nbr := Atoi(os.Args[1])
// 		array := []roman{
// 			{1000, "M", "M+"},
// 			{900, "CM", "(M-C)+"},
// 			{500, "D", "D+"},
// 			{400, "CD", "(D-C)+"},
// 			{100, "C", "C+"},
// 			{90, "XC", "(C-X)+"},
// 			{50, "L", "L+"},
// 			{40, "XL", "(L-X)+"},
// 			{10, "X", "X+"},
// 			{9, "IX", "(X-I)+"},
// 			{5, "V", "V+"},
// 			{4, "IV", "(V-I)+"},
// 			{1, "I", "I+"},
// 		}
// 		if nbr == 0 || nbr >= 4000 {
// 			fmt.Printf("ERROR: can not convert to roman digit")
// 		} else {
// 			ans := ""
// 			res := ""
// 			for i := 0; i < len(array); i++ {
// 				if nbr/array[i].num != 0 {
// 					m := nbr / array[i].num
// 					for j := 0; j < m; j++ {
// 						ans += array[i].wr
// 						res += array[i].rm
// 					}
// 					nbr %= array[i].num

// 				}
// 			}
// 			fmt.Printf(ans[:len(ans)-1])
// 			fmt.Printf("\n")
// 			fmt.Printf(res)
// 		}
// 		fmt.Printf("\n")
// 	}
// }

// func Atoi(s string) int {
// 	num := 0
// 	for _, k := range s {
// 		if k >= '0' && k <= '9' {
// 			num = num*10 + int(k-48)
// 		} else {
// 			return 0
// 		}
// 	}
// 	return num
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	args := os.Args[1:]
// 	if len(args) != 1 {
// 		return
// 	} else if Atoi(args[0]) >= 4000 || !isNumber(args[0]) {
// 		err := "ERROR: cannot convert to roman digit"
// 		fmt.Printf("%v\n", err)
// 	} else {
// 		result := intToRoman(Atoi(args[0]))
// 		fmt.Printf("%v\n", result)
// 	}
// }

// func Atoi(s string) int {
// 	number := 0
// 	for _, w := range s {
// 		number = number*10 + int(w-48)
// 	}
// 	return number
// }

// func isNumber(s string) bool {
// 	for _, w := range s {
// 		if w < '0' || w > '9' {
// 			return false
// 		}
// 	}
// 	return true
// }

// func intToRoman(num int) string {
// 	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
// 	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
// 	result := ""
// 	formula := ""
// 	i := 0
// 	for num > 0 {
// 		for values[i] > num {
// 			i++
// 		}
// 		num -= values[i]
// 		result += romans[i]
// 		if len(romans[i]) == 2 {
// 			formula += "(" + string(romans[i][1]) + "-" + string(romans[i][0]) + ")" + "+"
// 		} else {
// 			formula += string(romans[i]) + "+"
// 		}
// 	}
// 	formula = formula[:len(formula)-1]
// 	fmt.Printf("%v\n", formula)
// 	return result
// }
