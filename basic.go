package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = "hello"
	cc = false
)

/*
定义变量：
变量名 变量类型
它都是有一个初始值
*/
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d  %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def" //:=定义一个变量
	b = 5
	fmt.Println(a, b, c, s)
}

//验证欧拉公式
func enler() {
	//c := 3 + 4i
	//fmt.Println(cmplx.Abs(c))
	fmt.Printf("%0.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1) //math.E等同于exp方法
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
}

//强制转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

//常量定义
func consts() {
	const fileName = "123.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(fileName, c)
}

//简单枚举
func enums() {
	//const (
	//	cpp    = 0
	//	java   = 1
	//	python = 2
	//	golang = 3
	//)
	const (
		cpp = iota //表示自增值
		_
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, gb, tb, pb)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)
	enler()
	triangle()
	consts()
	enums()
}
