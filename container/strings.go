package main

import (
	"fmt"
	"unicode/utf8"
)

/*
rune相当于go的char
使用range遍历pos，rune对
使用utf8.RuneCountInString获得字符数量
使用len获得字节长度
使用[]byte获得字节
*/
/*
其他字符串操作
Fields,
Split,
Join,
Contains,查找子串
Index,
ToLower,
ToUpper,
Trim,
TrimRight,
TrimLeft,
*/
func main() {
	s := "yes,我很爱青青"
	fmt.Println(len(s))
	fmt.Printf("%s\n", []byte(s))
	for _, b := range []byte(s) {
		fmt.Printf("%x ", b)
	}
	fmt.Println()
	for i, ch := range s { //ch is a rune
		fmt.Printf("(%d %x)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
}
