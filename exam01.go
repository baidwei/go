package main //声明package

/*
Go语言主要特征：
1.自动垃圾回收
2.更丰富的内置类型
3.函数多返回值
4.错误处理
5.匿名函数和闭包
6.类型和接口
7.并发编程
8.反射
9.语言交互性
10.高性能/高效开发
*/
import (
	"fmt" //import包，不能包含没有用到的包，否则编译错误
)

func main() { //入口函数，无参数，无返回值
	fmt.Printf("hello golang world!")
	//go build exam01.go
	//go run exam01 或者exam01

}

/*
go build 编译
go clean 移除当前源码包里面的编译生成文件
go fmt   格式化代码
go get   动态获取远程代码包
go install 生成结果文件，并将编译好的结果移到$GOPATH/pkg或者$GPATH/bin
go test  运行测试用的可执行文件
go doc   godoc -http=:8080  查看文档
go fix   修复以前老版本代码到新版本
go version 查看当前版本
go env   查看当前go的环境变量
go list  列出当前所有安装package
go run   编译并运行go语言程序
*/
