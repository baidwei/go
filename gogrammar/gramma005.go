package main

import (
	"fmt"
)

func main() {
	fmt.Println("数组 Array")
	var a [3]int = [3]int{0, 1, 2}
	var b [3]int = [3]int{}
	var c [3]int
	var d = [...]int{0, 1, 2, 3}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// 多维数组只能自动计算最外围数组长度
	x := [...][2]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println(x)
}

/*
数组 Array
数组声明带有长度信息且长度固定，数组是值类型默认零值不是nil，传递参数时会进行复制。
声明定义数组时中括号[ ]在类型名称之前，赋值引用元素时中括号[ ]在数组变量名之后
使用...自动计算数组的长度
初始化指定索引的数组元素，未指定初始化的元素保持默认零值

*/
