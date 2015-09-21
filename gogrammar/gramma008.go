package main

import (
	"fmt"
)

type S struct {
	A    int
	B, c string
}

type (
	A struct {
		v int
	}
	B struct { // 定义结构体B，嵌入结构体A作为匿名字段
		A
	}
	C struct {
		*A
	}
)

func (a *A) setV(v int) {
	a.v = v
}

func (a A) getV() int {
	return a.v
}

func (b B) getV() string {
	return "B"
}

func (c *C) getV() bool {
	return true
}

//匿名结构体声明时省略了type关键字，并且没有名称
type Integer int

// 声明变量aa为空的匿名结构体类型
var aa struct{}

// 声明变量bb为包含一个字段的匿名结构体类型
var bb struct{ x int }

// 声明变量cc为包含两个字段的匿名结构体类型
var cc struct {
	u int
	v bool
}

func main() {
	fmt.Println("结构体 Struct")

	var s S = S{0, "1", "2"}
	fmt.Println(s)

	var ss S = S{B: "b", A: 1}
	fmt.Println(ss)
	fmt.Println(ss.A)

	a := A{}
	b := B{}     // 初始化结构体B，其内匿名字段A默认零值是A{}
	c := C{&A{}} // 初始化结构体C，其内匿名指针字段*A默认零值是nil，需要初始化赋值

	fmt.Println(a.v)
	// 结构体A嵌入B，A内字段自动提升到B
	fmt.Println(b.v)

	// 结构体指针*A嵌入C，*A对应结构体内字段自动提升到C
	fmt.Println(c.v)

	a.setV(3)
	b.setV(5)
	c.setV(7)
	fmt.Println(a.getV(), b.A.getV(), c.A.getV())
	fmt.Println(a.getV(), b.getV(), c.getV())

	// 声明d为包含3个字段的匿名结构体并初始化部分字段
	d := struct {
		x int
		y complex64
		z string
	}{
		z: "string",
		x: 111,
	}
	d.y = 22 + 33i
	fmt.Println(d)

	// 声明变量e为包含两个字段的匿名结构体类型
	// 包含1个匿名结构体类型的命名字段和1个命名类型的匿名字段
	e := struct {
		a struct{}
		// 结构体组合嵌入匿名字段只支持命名类型
		Integer
	}{}

	e.Integer = 4
	fmt.Println(e)
}

/*
结构体 Struct
结构体类型struct是一个字段的集合
结构体初始化通过结构体字段的值作为列表来新分配一个结构体
使用 Name: 语法可以仅列出部分字段(字段名的顺序无关)
结构体是值类型，传递时会复制值，其默认零值不是nil
结构体组合
将一个命名类型作为匿名字段嵌入一个结构体
嵌入匿名字段支持命名类型、命名类型的指针和接口类型

匿名结构体
匿名结构体声明时省略了type关键字，并且没有名称

*/
