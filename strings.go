package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ok,开始学习Go！"
	fmt.Println(len(s))
	// 英文一个字节，中文三个字节
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d %x) ", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:",utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ",ch)
	}
	fmt.Println()

	// rune 4个字节
	for i, ch := range []rune(s) { // ch is a rune
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()

}
