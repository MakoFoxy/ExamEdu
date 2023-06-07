// swapbits
// Instructions

// Write a function that takes a byte, swaps its halfs (like the example) and returns the result.
// Expected function

// func SwapBits(octet byte) byte {

// }

// Example:

// 1 byte

// 0100 | 0001
//     \ /
//     / \
// 0001 | 0100

// allowedFunctions
// 0	"--allow-builtin"
package main

func SwapBits(octet byte) byte {
	return octet<<4 + octet>>4
}

// func SwapBits(octet byte) byte {
// 	return octet<<4 + octet>>4
// }

// func SwapBits(octet byte) byte {
// 	return octet<<4 + octet>>4
// }
