package main

import (
	"fmt"
)

func add() {
	fmt.Println("my use add(){} function,and no return values.")
}

func add1() int {
	fmt.Println("my use add1() int {} function,and return value is", 0)
	return 0
}

func add2(x int, y int) int {
	fmt.Println("my use add2(x int,y int)int {} function,and return x+y=", x+y)
	return x + y
}

func add3(x int, y int) (int, int) {
	fmt.Println("my use add3(x int,y int)(int,int){} function,and return x+y=", x+y, "and x*y=", x*y)
	return x + y, x * y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	add()
	add1()
	add2(1, 2)
	add3(1, 2)
	fmt.Println(split(17))
}
