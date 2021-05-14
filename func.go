package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation:  %s", op)
	}

}

//取余除法
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//option是一个函数，有两个int类型的参数，返回一个int类型的结果，
//a,b两个参数是op需要的参数，最后返回int类型的结果
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args"+"(%d ,%d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))

}

///...int表示随便传多少个int值都可以
//可变参数列表
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	fmt.Println(eval(3, 4, "x"))
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(
		//直接定义匿名函数
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6))
}
