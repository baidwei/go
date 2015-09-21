package main

import "fmt"

func main() {
	//Switch语法
	//支持初始化表达式
	switch a := 5; a { //默认break，匹配成功后不会自动向下执行其他case,而是跳出整个switch
	case 0: //普通
		fmt.Println(0)
	case 1, 2: //多个分支，逗号分隔
		fmt.Println(1, 2)
	case 100: //什么都不做
	case 5:
		fmt.Println(5)
		fallthrough //进入后面的case 进行处理，而不是跳出block
	default:
		fmt.Println("default")
	}
	//注意，不需要break来明确退出一个case，一旦条件符合，自动终止,除非使用fallthough
}

/*
可以不带表达式
switch sExpr {
        case expr1: //sExpr和expr1必须类型一致,不限制为常量或者证书，可以用任何类型或表达式
            ...
}
*/
