// reversebits
// Instructions

// Write a function that takes a byte, reverses it bit by bit (as shown in the example) and returns the result.
// Expected function

// func ReverseBits(oct byte) byte {

// }

// Example:

// 1 byte

// 0010 0110 || / 0110 0100

// allowedFunctions
// 0	"os.Args"
// 1	"fmt.*"
// 2	"--allow-builtin"

package main

import "fmt"

func main() {
	fmt.Println(ReverseBits(38))
}

func ReverseBits(oct byte) byte {
	var res byte = 0

	for i := 8; i > 0; i-- {
		res = (res << 1) | (oct & 1)
		oct = (oct >> 1)
	}
	return res
}

// func ReverseBits(oct byte) byte {
// 	var res byte = 0

// 	for i := 1; i <= 8; i++ {
// 		// fmt.Println("{ res", res)
// 		res = (res << 1) | (oct & 1)
// 		oct = (oct >> 1)
// 		// fmt.Println("} res", res)
// 		// fmt.Println()
// 	}
// 	return res
// }

// func ReverseBits(oct byte) byte {
// 	var res byte = 0

// 	for i := 8; i > 0; i-- {
// 		res = (res << 1) | (oct & 1)
// 		oct = (oct >> 1)
// 	}
// 	return res
// }

// func ReverseBits(oct byte) byte {
// 	var result byte
// 	for i := 0; i < 8; i++ {
// 		result = result<<1 | oct&1
// 		oct = oct >> 1
// 	}
// 	return result
// }

// package main
// func ReverseBits(oct byte) byte {
// 	var result byte = 0
// 	for onybyte := 8; onybyte > 0; onybyte-- {
// 		result = (result << 1) | (oct & 1)
// 		oct >>= 1
// 	}
// 	return result
// }
