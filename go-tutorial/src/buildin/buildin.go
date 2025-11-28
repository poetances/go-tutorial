package buildin

import "fmt"

/*
func append(slice []Type, elems ...Type) []Type
func cap(v Type) int
func close(c chan<- Type)
func copy(dst, src []Type) int
func delete(m map[Type]Type1, key Type)
func len(v Type) int
func make(Type, size IntegerType) Type
func new(Type) *Type
func complex(real, imag FloatType) ComplexType
func imag(c ComplexType) FloatType
func real(c ComplexType) FloatType
func panic(v interface{})
func recover() interface{}

builtin.go 只是一个声明文件，不包含实现。
	•	真正的实现大部分在 runtime 包，如：
	•	runtime.growslice() 处理 append()
	•	runtime.makemap() 处理 make(map)
	•	runtime.closechan() 处理 close()
	•	runtime.panic() 处理 panic()

总结
	1.	builtin.go 声明了一些特殊函数，但它们的实现由编译器或 runtime 负责。
	2.	内置函数分为几类：
	•	创建对象（make()、new()）
	•	查询属性（len()、cap()）
	•	操作集合（append()、copy()、delete()）
	•	异常处理（panic()、recover()）
	•	通道控制（close()）
	•	复杂数计算（complex()、real()、imag()）
	3.	这些函数是 Go 语言的核心特性，不能被重写，也不能直接调用 builtin 包。

关于内置函数
	•	builtin.go 只是声明文件，没有具体实现，所有内置函数由 编译器或 runtime 直接优化，提高性能。
	•	Go 选择 builtin.go 这种设计，而不是 Swift 直接给 slice/map 增加方法，主要是 性能、简洁性、低级控制 的考虑。
	•	make() 只能用于 slice/map/channel，因为它们的底层数据结构必须由 runtime 进行特殊初始化。
	•	Swift 的方式更灵活，Go 的方式更高效，各有优缺点。
*/

func init() {
	fmt.Println("build int init")
}

func BuildInTutorial() {
	// ptr使用
	ptr := new(int)
	*ptr = 10
	println(*ptr)
}
