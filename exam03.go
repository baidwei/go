package main

import "fmt"
import "math"

//Go支持字符常量，字符串常量，布尔常量和数值常量，使用关键字const定义常量

//const定义常量
const s string = "constant"

func main() {
	fmt.Println(s)
	//const 常量关键字，可以出现在任何var出现的位置
	//区别，常量必须有初始值
	const a = 500000000

	const b = 3e20 / a
	fmt.Println(b)

	//显式类型转换
	fmt.Println(int64(b))

	//数值型常量，可以在程序上下文中获取类型
	//比如变量赋值或者函数调用
	//例如：对math包中的sin函数，它需要一个float64类型的变量
	fmt.Println(math.Sin(a))
}
