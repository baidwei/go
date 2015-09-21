package main

import (
	"fmt"
)

func main() {
	fmt.Println("基础数据类型")
}

/*
基础数据类型 Basic data type
基本类型包含：数值类型，布尔类型，字符串
类型			取值范围			默认值
int				int32,int64			0
int8			-2^7~2^7-1			0
int16			-2^15~2^15-1		0
int32,rune		-2^31~2^31-1		0
int64			-2^63~2^63-1		0
float32			IEEE-754 32-bit		0.0
complex64		float32+floag32i	0+0i
bool			true,false			false
uintptr			int32,int64			0

uint			uint32,uint64		0
uint8,byte		0~2^8-1				0
uint16			0~2^16-1			0
uint32			0~2^32-1			0
unit64			0~2^64-1			0
float64			IEEE-754 64-bit		0.0
complex128		float64+floag64i	0+0i
string			""~"无穷"			"",''
error			-					nil
*/
/*
byte是uint8的别名
rune是int32的别名,代表一个UNICODE码点
int与int32或者int64是不同类型，只是根据架构对应32/64位值
uint与uint32或者uint64是不同类型，只是根据架构对应32/64位值
*/
