package main

import (
	"fmt"
	"math"
)

func main() {
	var a []int
	fmt.Println(a, len(a), cap(a))

	// 可用类似数组的方式初始化slice
	var b []int = []int{0, 1, 2}
	fmt.Println(b, len(b), cap(b))
	var c = []string{2: "123", 3: "345"}
	fmt.Println(c)

	var d = make([]int, 0)
	fmt.Println(d, len(d), cap(d))

	var e = make([]int, 3, 10)
	fmt.Println(e, len(e), cap(e))

	var f = new([]int)
	fmt.Println(f, len(*f), cap(*f))

	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//2个参数 slice[beginIndex:endIndex]
	sa := s[1:3]
	sb := s[:4]
	sc := s[1:]
	sd := s[1:1]
	se := s[:]
	fmt.Println(sa, sb, sc, sd, se)

	x := make([]int, 5, 10)
	//3个参数 slice[beginIndex:endIndex:capIndex]
	xa := x[9:10:10]
	xb := x[:3:5]
	fmt.Println(x, len(x), cap(x))
	fmt.Println(xa, len(xa), cap(xa))
	fmt.Println(xb, len(xb), cap(xb))

	ss := []string{}
	fmt.Println(ss)
	ss = append(ss, "a")
	fmt.Println(ss)
	ss = append(ss, "b", "c", "d")
	fmt.Println(ss)
	t := []string{"e", "f", "g"}
	ss = append(ss, t...)
	fmt.Println(ss)
	ss = append(ss, t[:2]...)
	fmt.Println(ss)

	ss[0] = "aa"
	fmt.Println(ss)

	ss[len(ss)-1] = "end"
	fmt.Println(ss)

	fmt.Printf("Now ,you have %g problems.", math.Nextafter(2, 3))

	deleteByAppend()
	deleteByCopy()
}
func deleteByAppend() {
	i := 3
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	s = append(s[:i], s[i+1:]...)
	fmt.Println(s)
}

func deleteByCopy() {
	i := 3
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	copy(s[i:], s[i+1:])
	s = s[:len(s)-1]
	fmt.Println(s)
}

/*
slice 切片是对一个数组上的连续一段的引用，并且同时包含了长度和容量信息
因为是引用类型，所以未初始化时的默认零值是nil，长度与容量都是0

使用内置函数make初始化slice，第一参数是slice类型，第二参数是长度，第三参数是容量(省略时与长度相同)

基于slice或数组重新切片，创建一个新的 slice 值指向相同的数组
重新切片支持两种格式：

2个参数 slice[beginIndex:endIndex]
需要满足条件：0 <= beginIndex <= endIndex <= cap(slice)
截取从开始索引到结束索引-1 之间的片段
新slice的长度：length=(endIndex - beginIndex)
新slice的容量：capacity=(cap(slice) - beginIndex)
beginIndex的值可省略，默认为0
endIndex 的值可省略，默认为len(slice)

3个参数 slice[beginIndex:endIndex:capIndex]
需要满足条件：0 <= beginIndex <= endIndex <= capIndex <= cap(slice)
新slice的长度：length=(endIndex - beginIndex)
新slice的容量：capacity=(capIndex - beginIndex)
beginIndex的值可省略，默认为0

向slice中增加/修改元素

删除slice中指定的元素
因为slice引用指向底层数组，数组的长度不变元素是不能删除的，所以删除的原理就是排除待删除元素后用其他元素重新构造一个数组

切片中有两个概念：一是len长度，二是cap容量，长度是指已经被赋过值的最大下标+1，可通过内置函数len()获得。容量是指切片目前可容纳的最多元素个数，可通过内置函数cap()获得。切片是引用类型，因此在当传递切片时将引用同一指针，修改值将会影响其他的对象。



var 创建变量

const 创建常量

iota 这个关键字用来声明enum的时候采用，它默认开始值是0，每调用一次加1

map 也就是Python中字典的概念，它的格式为map[keyType]valueType map的读取和设置也类似slice一样，通过key来操作，只是slice的index只能是｀int   ｀类型，而map多了很多类型，可以是int，可以是string

make 用于内建类型（map、slice和channel）的内存分配

new 用于各种类型的内存分配

goto 跳转到必须在当前函数内定义的标签

func 关键字func用来声明一个函数funcName

defer  延迟执行代码，类似于析构函数

panic 中断原有的控制流程

recover 恢复中断的函数

import 导入包文件

Go程序设计的一些规则

    大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量；小写字母开头的就是不可导出的，是私有变量。
    大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数；小写字母开头的就是有private关键词的私有函数。



Go里面有两个保留的函数：init函数（能够应用于所有的package）和main函数（只能应用于package main）。这两个函数在定义时不能有任何的参数和返回值。虽然一个package里面可以写任意多个init函数，但这无论是对于可读性还是以后的可维护性来说，都强烈建议用户在一个package中每个文件只写一个init函数。

Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数。每个package中的init函数都是可选的，但package main就必须包含一个main函数。
*/
