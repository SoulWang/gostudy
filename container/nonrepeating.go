package main

import "fmt"

/*
寻找最长不含重复字符的子串
对于每一个字母s
laseOccurred[s]不存在，或者<start时，无需操作
laseOccurred[s]>=start时。更新start
更新laseOccurred[s]，更新maxLength
*/
func lengthIfNonRepeatingSubstr(s string) int {
	laseOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := laseOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		laseOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthIfNonRepeatingSubstr("abcabcbb"))
	fmt.Println(lengthIfNonRepeatingSubstr("bbbbb"))
	fmt.Println(lengthIfNonRepeatingSubstr("pwwkew"))
	fmt.Println(lengthIfNonRepeatingSubstr(""))
	fmt.Println(lengthIfNonRepeatingSubstr("b"))
	fmt.Println(lengthIfNonRepeatingSubstr("abcdef"))
	fmt.Println(lengthIfNonRepeatingSubstr("我很喜欢golang"))
	fmt.Println(lengthIfNonRepeatingSubstr("一二三三二一"))

	//strings.遇到字符串首先去这个包里面找是否有对应的方法
}
