package main

import "fmt"

func printarray(arr []int) {
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	//range表示数组的遍历
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("printarray(arr1)")
	printarray(arr1[:])
	fmt.Println("printarray(arr3)")
	printarray(arr3[:])
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}
