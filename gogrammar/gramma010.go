package main

import (
	"fmt"
)

var ch chan int = make(chan int)

func sayChan(words int) {
	ch <- words
}
func main() {
	fmt.Println("通道 Channel")

	//var c0 = make(chan int)
	c0 := make(chan int)
	//var c0 chan int
	//c0 = make(chan int)
	go func() {
		fmt.Println("send data")
		c0 <- 3
		close(c0)
	}()

	fmt.Println("waiting receive data")
	fmt.Println("receive data:", <-c0)
	fmt.Println("再次从已关闭的channel中取值，返回channel传递类型的零值")
	fmt.Println(<-c0)

	for i := 0; i < 100; i++ {
		go sayChan(i)
	}

	fmt.Println("waiting receive")
	for j := 0; j < 100; j++ {
		s, ok := <-ch
		if ok {
			fmt.Println("receive:", s)
		} else {
			fmt.Println("end")
		}
		fmt.Println(ok)
	}

}

/*
channel用于两个goroutine之间传递指定类型的值来同步运行和通讯。
操作符<-用于指定channel的方向，发送或接收。
如果未指定方向，则为双向channel。


channel是引用类型，使用make函数来初始化。
未初始化的channel零值是nil，且不能用于发送和接收值


无缓冲的channe中有值时发送方会阻塞，直到接收方从channel中取出值。
带缓冲的channel在缓冲区已满时发送方会阻塞，直到接收方从channel中取出值。
接收方在channel中无值会一直阻塞。

通过channel发送一个值时，<-作为二元操作符使用

通过channel接收一个值时，<-作为一元操作符使用

关闭channel，只能用于双向或只发送类型的channel
只能由 发送方调用close函数来关闭channel
接收方取出已关闭的channel中发送的值后，后续再从channel中取值时会以非阻塞的方式立即返回channel传递类型的零值。

使用for range语句依次读取发送到channel的值，直到channel关闭。


*/
