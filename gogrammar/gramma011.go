package main

import "fmt"

type Reader interface {
	Read(b []byte) (n int)
}
type Writer interface {
	Write(b []byte) (n int)
}

// 接口ReadWriter组合了Reader和Writer两个接口
type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	//...
}

func (f *File) Read(b []byte) (n int) {
	fmt.Println("read", len(b), "bytes data.")
	return len(b)
}
func (f *File) Write(b []byte) (n int) {
	fmt.Println("write", len(b), "bytes data.")
	return len(b)
}
func main() {
	fmt.Println("接口 Interface")
	var f *File = &File{}
	var r Reader = f
	var w Writer = f
	var rw ReadWriter = f
	bs := []byte("abcd")
	r.Read(bs)
	rw.Read(bs)
	w.Write(bs)
	rw.Write(bs)
}

//https://218.202.81.130/WebXJ/login.jsp
/*
接口类型是由一组方法定义的集合。
接口类型的值可以存放实现这些方法的任何值。

类型通过实现定义的方法来实现接口， 不需要显式声明实现某接口。

内置接口类型error是一个用于表示错误情况的常规接口，其零值nil表示没有错误
所有实现了Error方法的类型都能表示为一个错误


Go中支持自定义的类型可基于： 基本类型、数组类型、切片类型、字典类型、函数类型、结构体类型、通道类型、接口类型以及自定义类型的类型


*/
