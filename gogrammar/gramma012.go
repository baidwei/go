package main

import (
	"fmt"
)

func f0() int {
	return 333
}
func main() {
	fmt.Println("语句 Statement")
	fmt.Println("条件语句 if")
	i := 10
	if i < 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(i)
	}
	//可以在条件之前执行一个简单的语句，由这个语句定义的变量的作用域仅在 if / else if / else 范围之内
	if i := 0; i < 10 {
		fmt.Println(i)
	}

	//if语句作用域范围内定义的变量会覆盖外部同名变量，与方法函数内局部变量覆盖全局变量同理
	a, b := 1, 2
	fmt.Println(a, b)
	if a, b := 3, 4; a > 1 && b > 2 {
		fmt.Println(a, b)
	}
	fmt.Println(a, b)

	//if判断语句类型断言
	x := 9
	checkType(x)
	checkType(f0)
}

func checkType(x interface{}) {
	// 断言传入的x为int类型，并获取值
	if i, ok := x.(int); ok {
		fmt.Println("int :", i)
	}
	if f, ok := x.(func() int); ok {
		fmt.Println("func :", f())
	}
	// 如果传入x类型为int，则可以直接获取其值
	//a := x.(int)
	//fmt.Println(a)

	// 如果传入x类型不是byte，则会产生恐慌panic
	//b := x.(byte)
	//fmt.Println(b)

}

/*
分号/括号 ; {

Go是采用语法解析器自动在每行末尾增加分号，所以在写代码的时候可以省略分号。
Go编程中只有几个地方需要手工增加分号：
for循环使用分号把初始化、条件和遍历元素分开。
if/switch的条件判断带有初始化语句时使用分号分开初始化语句与判断语句。
在一行中有多条语句时，需要增加分号。

控制语句(if，for，switch，select)、函数、方法 的左大括号不能单独放在一行， 语法解析器会在大括号之前自动插入一个分号，导致编译错误


条件语句 if
*/
