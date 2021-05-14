package main

import "fmt"

/**
1.创建：make(map[string]int)
2.获取元素：m[key]
3.key不存在时，获得Value类型的初始值zero value
4.用value,ok:=m[key]来判断是否存在key
5.用delete删除一个key
6.使用range遍历key，或者遍历key-value对
7.不保证遍历顺序，如需顺序，需手动对key排序
8.使用len获得元素个数
*/
func main() {
	/*map的第一种建立方法*/
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	/*map的第二种建立方法*/
	m2 := make(map[string]int) //m2==empty map
	fmt.Println(m2)

	/*map的第三种建立方法*/
	var m3 map[string]int //m3== nil
	fmt.Println(m3)

	/*遍历map，会发现key是无序*/
	fmt.Println("Travering map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	/*当我们去取map里面不存在的key的时候，他还是能够拿到值的，zero value*/
	if caurseName, ok := m["caurse"]; ok {
		fmt.Println(caurseName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)
	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}

/**
map使用哈希表，必须可以比较相等
除了slice，map，function的内建类型都可以作为key
Struct类型不包含上述字段，也可作为key
*/
