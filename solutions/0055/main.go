package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scanln(&k)
	var strBytes []byte
	fmt.Scanln(&strBytes)

	if k < 0 || len(strBytes) < k {
		return
	}

	reverseString(strBytes)
	reverseString(strBytes[0:k])
	reverseString(strBytes[k:len(strBytes)])

	fmt.Printf(string(strBytes))
}

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
