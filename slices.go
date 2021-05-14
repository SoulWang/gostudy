package main

import "fmt"

func updateSlece(s []int) {
	s[0] = 100
}

//s1的下标为0.1.2.3
//因为s1里面是对arr底层数组的一个view，不是说到3就结束了
//这个下标到3之后还会打印一个4,5，这个我们直接用s1[4],s1[5]是看不到的
//所以s2取的是s1的3-5的数，相对就是0,1,2下标所对应的，也就说最后对应的就是图示
//slice可以向后扩展，不可以向前扩展
//s[i]不可以超越len(s),向后扩展不可以超越底层数组cap(s)
func sliceTest() {
	fmt.Println("Extending slice")
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("arr = %v, len(sarr)=%d,cap(arr)=%d\n",
		arr2, len(arr2), cap(arr2))
	s1 := arr2[2:6]
	fmt.Printf(
		"s1 = %v, len(s1)=%d,cap(s1)=%d\n",
		s1, len(s1), cap(s1))
	s2 := s1[3:5]
	fmt.Printf(
		"s2 = %v, len(s2)=%d,cap(s2)=%d\n",
		s2, len(s2), cap(s2))
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	//s4，s5 view different array,or no longer view arr
	//添加元素时如果超越cap，系统会重新分配更大的底层数组
	//由于值传递的关系，必须接受append的返回值
	//s = append(s,val)
	fmt.Println("arr = ", arr2)

}

//这些都是切片类型,相当于是一个视图
//切片Slice本身没有数据，是对底层array的一个view
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])

	s1 := arr[2:]
	fmt.Println("s1 = ", s1)

	s2 := arr[:]
	fmt.Println("s2 = ", s2)

	fmt.Println("After updateSlice(s1)")
	updateSlece(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlece(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	fmt.Println("Reslice before =", s2)
	s2 = s2[:5]
	fmt.Println("The first change = ", s2)
	s2 = s2[2:]
	fmt.Println("The last change =", s2)

	sliceTest()
}
