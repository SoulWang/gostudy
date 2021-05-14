package main

import "fmt"

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

	fmt.Println("hello")
}
