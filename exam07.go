package main

import "fmt"

func main() {
	//goto break continue
	a := 1
LABEL1:
	fmt.Println("inc a:=", a)
	a += 1
LABEL2:

	for a < 6 {
		fmt.Println(a)
		if a == 3 {
			a += 1
			continue LABEL2
		}
		if a == 5 {
			break
		}
		goto LABEL1
	}

}

/*
均可配合标签使用(标签区分大小写，若声明了没有使用会导致编译错误)
break/continue可配合标签用于多重循环跳出
goto是调整执行位置，与其他两个执行结果并不相同
1.goto
支持函数内部goto跳转
请善用,必须跳转到当前函数内定义的标签,标签名大小写敏感
continue 进入下一次循环
break    终止循环
*/
