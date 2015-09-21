package main

import (
	"fmt"
)

func main() {
	fmt.Println("内建函数 Builtin")
	fmt.Println("append 内建函数append将元素追加到切片的末尾")
	fmt.Println("func cap(v Type) int 内建函数cap返回 v 的容量")
	fmt.Println("func close(c chan<- Type) 内建函数close关闭信道，该通道必须为双向的或只发送的。它应当只由发送者执行，而不应由接收者执行，其效果是在最后发送的值被接收后停止该通道")
	fmt.Println("func complex(r, i FloatType) ComplexType 使用实部r和虚部i生成一个复数。")
	fmt.Println("func copy(dst, src []Type) int 内建函数copy将元素从来源切片复制到目标切片中，也能将字节从字符串复制到字节切片中。copy返回被复制的元素数量，它会是 len(src) 和 len(dst) 中较小的那个。来源和目标的底层内存可以重叠。")
	fmt.Println("func delete(m map[Type]Type1, key Type) 内建函数delete按照指定的键将元素从映射中删除。若m为nil或无此元素，delete不进行操作")
	fmt.Println("func imag(c ComplexType) FloatType 返回复数c的虚部。")
	fmt.Println("func len(v Type) int 内建函数len返回 v 的长度")
	fmt.Println("func make(Type, size IntegerType) Type 内建函数make分配并初始化一个类型为切片、映射、或通道的对象")
	/*其第一个实参为类型，而非值。make的返回类型与其参数相同，而非指向它的指针。其具体结果取决于具体的类型：
	切片：size指定了其长度。该切片的容量等于其长度。切片支持第二个整数实参可用来指定不同的容量；它必须不小于其长度，因此 make([]int, 0, 10) 会分配一个长度为0，容量为10的切片。
	映射：初始分配的创建取决于size，但产生的映射长度为0。size可以省略，这种情况下就会分配一个小的起始大小。
	通道：通道的缓存根据指定的缓存容量初始化。若 size为零或被省略，该信道即为无缓存的。
	*/
	fmt.Println("func new(Type) *Type 内建函数new分配内存。其第一个实参为类型，而非值。其返回值为指向该类型的新分配的零值的指针")
	fmt.Println("func panic(v interface{}) 内建函数panic停止当前Go程的正常执行")
	/*
		内建函数panic停止当前Go程的正常执行。当函数F调用panic时，F的正常执行就会立刻停止。F中defer的所有函数先入后出执行后，F返回给其调用者G。G如同F一样行动，层层返回，直到该Go程中所有函数都按相反的顺序停止执行。之后，程序被终止，而错误情况会被报告，包括引发该恐慌的实参值，此终止序列称为恐慌过程
	*/
	fmt.Println("func print(args ...Type) 内建函数print以特有的方法格式化参数并将结果写入标准错误，用于自举和调试。")
	fmt.Println("func println(args ...Type) println类似print，但会在参数输出之间添加空格，输出结束后换行")
	fmt.Println("func real(c ComplexType) FloatType 返回复数c的实部")
	fmt.Println("func recover() interface{} 内建函数recover允许程序管理恐慌过程中的Go程")
	/*
		内建函数recover允许程序管理恐慌过程中的Go程。在defer的函数中，执行recover调用会取回传至panic调用的错误值，恢复正常执行，停止恐慌过程。若recover在defer的函数之外被调用，它将不会停止恐慌过程序列。在此情况下，或当该Go程不在恐慌过程中时，或提供给panic的实参为nil时，recover就会返回nil。
	*/

	/*
			初始化函数 init
			init函数是用于程序执行前做包的初始化工作的函数
			init函数的声明没有参数和返回值
			一个package或go源文件可以包含零个或多个init函数
			init函数被自动调用，在main函数之前执行，不能在其他函数中调用，显式调用会报错该函数未定义

		   	所有init函数都会被自动调用，调用顺序如下：
			   同一个go文件的init函数调用顺序是 从上到下的
			   同一个package中按go源文件名字符串比较 从小到大顺序调用各文件中的init函数
			   不同的package，如果不相互依赖的，按照main包中 先import的后调用的顺序调用其包中的init函数
			   如果package存在依赖，则先调用最早被依赖的package中的init函数
	*/
}
