package main

import (
	"bytes"
	"fmt"
)

func main2() {
	var strBytes []byte
	fmt.Scanln(&strBytes)

	for i := 0; i < len(strBytes); i++ {
		if strBytes[i] <= '9' && strBytes[i] >= '0' {
			inserElement := []byte{'n', 'u', 'm', 'b', 'e', 'r'}
			strBytes = append(strBytes[:i], append(inserElement, strBytes[i+1:]...)...)
			i = i + len(inserElement) - 1
		}
	}

	fmt.Printf(string(strBytes))
}

func main() {
	var strBytes []byte
	fmt.Scanln(&strBytes)

	numCount := 0
	for i := 0; i < len(strBytes); i++ {
		if strBytes[i] <= '9' && strBytes[i] >= '0' {
			numCount++
		}
	}

	oldLen := len(strBytes)
	insertBytes := []byte("number")
	insertLen := len(insertBytes)
	for i := 0; i < numCount; i++ {
		strBytes = append(strBytes, bytes.Repeat([]byte{' '}, insertLen-1)...)
	}

	l, r := oldLen-1, len(strBytes)-1
	for l < r {
		lByte := strBytes[l]
		rShift := 0
		if lByte >= '0' && lByte <= '9' {
			for i, b := range insertBytes {
				strBytes[r-insertLen+1+i] = b
			}
			rShift = insertLen
		} else {
			strBytes[r] = lByte
			rShift = 1
		}
		l--
		r -= rShift
	}

	fmt.Printf(string(strBytes))
}
