package main

import "fmt"

func main() {
	//1. var 关键字定义一个或多个变量
	var a string = "initial"
	fmt.Println(a)

	//2. 也可以一次定义多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	//3.GO也会推断具有初始化值的变量类型
	var d = true
	fmt.Println(d)

	//4. 定义变量时，没有给出初始值的变量，默认是零值
	//例如：整型的零值是0
	var e int
	fmt.Println(e)

	//5. 通过":="语法，可以定义并初始值
	f := "shot"
	fmt.Println(f)
}

/*
基本类型
类型			长度	说明
bool			1		true/false,默认false,不能把非0当成true(不用数字代表true/false)
byte			1		uint8别名
rune			4		int32别名，代表uncode code point
int/unit				依赖所运行的平台，32bit/64bit
int8/unit8		1		-128~127;0~255
int16/unit16	2		-32768~32767;0~65535
int32/unit32	4
int64/unit64	8
float32			4		精确到7位小数，相当于C的float类型
float64			8		精确到15位小数，相当于C的double类型

complex64		8
complex128		16

uintptr					足够保存指针的32位、64位整数，指针(可以存指针的整型类型)
array					值类型，数组
struct					值类型，结构体
string 					值类型，字符串类型，常用
slice					引用类型，切片
map						引用类型，字典
channel					引用类型，通道
interface				接口类型，接口
function 				函数类型，函数

类型转换
不支持隐形类型转换，必须进行显示类型转换


参考:http://www.jb51.net/list/list_246_1.htm
*/
