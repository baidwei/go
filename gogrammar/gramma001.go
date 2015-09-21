package main

import "fmt"

func main() {
	fmt.Println("使用package关键字声明当前源文件所在的包")
	fmt.Println("包声明语句是所有源文件的第一行非注释语句")
	fmt.Println("包名称中不能包含空白字符")
	fmt.Println("包名必须与源文件所在的目录名称保持一致")
	fmt.Println("每个目录中只能定义一个package")
	fmt.Println("导入路径是对应包在$GOROOT/pkg/$GOOS_$GOARCH/、$GOPATH/pkg/$GOOS_$GOARCH/或当前路径中的相对路径")
}

/*
// 导入$GOROOT/$GOOS_$GOARCH/中的相对路径包(官方标准库)
import "fmt"
import "math/rand"

// 导入$GOPATH/$GOOS_$GOARCH/中的相对路径包
import "github.com/user/project/pkg"
import "code.google.com/p/project/pkg"

用圆括号组合导入包
import ("fmt"; "math")

import (
    "fmt"
    "math"
)

导入包可以定义别名，防止同名称的包冲突
import (
    "a/b/c"

    c1 "x/y/c"     // 将导入的包c定义别名为 c1

    格式化 "fmt"    // 将导入的包fmt定义别名为 格式化

    m "math"       // 将导入的包math定义别名为 m
)
引用包名与导入路径的最后一个包目录名称一致
// 引用普通名称的导入包
c.hello()

// 引用定义别名的包
格式化.Println(m.Pi)


静态导入，在导入的包名之前增加一个小数点.
// 类似C中的include 或Java中的import static
import . "fmt"

// 然后像使用本包元素一样使用fmt包中可见的元素，不需要通过包名引用
Println("no need package name")

导入包但不直接使用该包，在导入的包名之前增加一个下划线_
// 如果当前go源文件中未引用过log包，将会导致编译错误
import "log"    // 错误
import . "log"  // 静态导入未使用同样报错

// 在包名前面增加下划线表示导入包但是不直接使用它，被导入的包中的init函数会在导入的时候执行
import _ "github.com/go-sql-driver/mysql"




规范导入路径
包声明语句后面添加标记注释，用于标识这个包的规范导入路径。
package pdf // import "rsc.io/pdf"


如果使用此包的代码的导入的路径不是规范路径，go命令会拒绝编译。
例如有 rsc.io/pdf 的一个fork路径 github.com/rsc/pdf
如下程序代码导入路径时使用了非规范的路径则会被go拒绝编译

import "github.com/rsc/pdf"




*/
