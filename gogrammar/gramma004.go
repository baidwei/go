package main

import (
	"fmt"
)

//a:=10  non-declaration statement outside function body
func main() {
	fmt.Println("常量Constant")
}

/*
常量可以是字符、字符串、布尔或数值类型的值，数值常量是高精度的值
根据常量值自动推导类型
常量组内定义时复用表达式
常量组内定义的常量只有名称时，其值会根据上一次最后出现的常量表达式计算相同的类型与值
自动递增枚举常量 iota
iota的枚举值可以赋值给数值兼容类型
每个常量单独声明时，iota不会自动递增
常量组合声明时，iota每次引用会逐步自增，初始值为0，步进值为1
即使iota不是在常量组内第一个开始引用，也会按组内常量数量递增
枚举的常量都为同一类型时，可以使用简单序列格式(组内复用表达式).

枚举序列中的未指定类型的常量会跟随序列前面最后一次出现类型定义的类型
iota自增值只在一个常量定义组合中有效，跳出常量组合定义后iota初始值归0
定制iota序列初始值与步进值 (通过组合内复用表达式实现)

http://yougg.github.io/static/gonote/gogrammar.html


*/
