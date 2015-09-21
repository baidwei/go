package main

import "fmt"

//for 循环是Go语言唯一的循环结构。
func main() {
	//1.基本的单一循环，类似其他语言的while语句
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	//2.经典的for循环，初始化，条件判断、循环后变化
	for j := 4; j <= 6; j++ {
		fmt.Println(j)
	}
	//无条件的FOR循环是死循环，除非使用break或return退出循环
	for {
		fmt.Println("loop")
		break
	}
}

//for循环和保留字range一起使用，完成迭代器iterator操作
//string, array, slice, map, channel都可以用range进行迭代器操作
//range返回序号和值，若是不想要，可以用_
