package main

import (
	"fmt"
)

type S struct {
	A int
	B string
	C string
}

func main() {
	fmt.Println("指针 Pointer")
	var i int = 1
	pi := &i // 指向数值的指针
	fmt.Println(pi)

	a := []int{0, 1, 2}
	pa := &a // 指向引用对象的指针
	fmt.Println(pa)

	var s *S = &S{0, "1", "2"}
	fmt.Println(s)

	//内置函数new(T)分配了一个零初始化的 T 值，并返回指向它的指针
	var ii = new(int)
	fmt.Println(ii)
	var ss *S = new(S)
	fmt.Println(ss)

	//使用*读取/修改指针指向的值
	k := new(int)
	*k = 3
	fmt.Println(k, *k)
	//指针使用点号来访问结构体字段
	//结构体字段/方法可以通过结构体指针来访问，通过指针间接的访问是透明的
	fmt.Println(s.A)
	fmt.Println((*s).A)

}

/*
指针 Pointer
通过取地址操作符&获取指向值/引用对象的指针
内置函数new(T)分配了一个零初始化的 T 值，并返回指向它的指针
使用*读取/修改指针指向的值
指针使用点号来访问结构体字段
结构体字段/方法可以通过结构体指针来访问，通过指针间接的访问是透明的。
指针的指针
跨层指针元素的使用
在指针引用多层对象时，指针是针对引用表达式的最后一位元素。

Go的指针没有指针运算
*/
