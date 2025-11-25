package basicdata

import "testing"
import "fmt"

/*
author: zcj
date: 2022-03-10

go 中的基本数据类型，都是存放在 buildin.go 中。

buildin.go  位于 $GOROOT/src/builtin/builtin.go
buildin.go 这些内置函数由 Go 的编译器合运行时直接处理，以提高性能和简化语言设计。
*/
func TestBasicData(t *testing.T) {
	// BoolenTutorial()
	// NumerTutorial()
	// StringTutorial()
	// ArrayTutorial()
	// SliceTutorial()
	// MapTutorial()
	// ChannelTutorial() 
}

/*
3. 设计哲学
Go 语言中 nil 的设计目标是：

安全性： nil 切片和映射可以安全地用于读取和长度操作。
一致性： nil 切片和空切片的行为一致（ len 为 0）。
显式性：需要显式初始化（如 make ）才能写入映射，避免隐式错误。
4. 对比其他语言
Java/C#： null 引用调用方法会抛出 NullPointerException 。
Python： None 没有类型信息，操作会直接报错。
Go： nil 是“可操作的”，但部分操作（如写入 nil 映射）会 panic。


值类型和引用类型的零值对比：
值类型（如数组、结构体、基本数据类型）的零值是其字段的零值，永远不为 nil 。
引用类型（如切片、映射、通道、指针、函数、接口）的零值是 nil 。
*/
func TestEmpty(t *testing.T) {
	var a [4]int
	fmt.Println(a) // 数组不是引用类型，永远不为nil

	var b []int
	fmt.Println(b, b == nil, len(b))
	
	var c map[string]int 
	fmt.Println(c, c == nil)

	var d string 
	fmt.Println(d, d)

	var e chan int
	fmt.Println(e, e)

	var f *int
	fmt.Println(f, f == nil)

	var g func()
	fmt.Println(g == nil)
}