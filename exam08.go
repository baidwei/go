package main

import "fmt"

func main() {
	//array,slice,map
	//1.go语言中，数组是一个值类型(value type)
	/*所有的值类型变量,在赋值和作为参数传递时,都将产生一个复制动作
	  如果作为函数的参数类型，则在函数调用时参数发生数据复制，在函数体中无法修改传入数组的内容
	  数组相等用 = != 比较，不能用 < >
	*/
	var a [5]int
	var c [2][3]int //二维
	fmt.Println(a)
	fmt.Println(c)
	var b = [5]int{1, 2, 3, 4, 5} //初始化并赋值
	fmt.Println(b)

}

//http://www.cnblogs.com/hustcat/p/4002707.html
/*
Google Go语言
golang语法详解笔记:
http://yougg.github.io/static/gonote/gogrammar.html
*/
