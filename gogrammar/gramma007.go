package main

import (
	"fmt"
)

func main() {

	fmt.Println("字典/映射 Map")
	var m map[int]int     //定义
	m = make(map[int]int) //初始化
	m[0] = 0              //赋值
	fmt.Println(m)

	mm := map[int]int{
		0: 0,
		1: 1,
	}
	mn := map[string]int{
		"a": 1,
		"b": 2,
	}
	mm[2] = 2
	mm[3] = 3
	mm[4] = 4
	fmt.Println(mm, mn)
	mm[0] = 100
	fmt.Println(mm)
	delete(mm, 2)
	fmt.Println(mm)

}

/*

map是引用类型，使用内置函数 make进行初始化，未初始化的map零值为 nil长度为0，并且不能赋值元素

使用内置函数make初始化map

map的使用：读取、添加、修改、删除元素


*/
