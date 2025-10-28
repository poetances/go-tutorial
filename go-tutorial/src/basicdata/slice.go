package basicdata

import "fmt"

/*
注意go语言中slice的声明方式
var identifier []type 

一个 slice 实际上是一个 结构体，包含：
  - 指针（Pointer）：指向底层数组的某个索引位置
  - 长度（Length）：slice 当前包含的元素个数
  - 容量（Capacity）：从 slice 开始位置到底层数组末尾的元素个数

type slice struct {
    array unsafe.Pointer // 指向底层数组的指针
    len   int           // 当前切片的长度（元素个数）
    cap   int           // 切片的容量（从底层数组的第一个元素到底层数组末尾的元素个数）
}

但实际上，slice表现出类引用类型，它的零值是 nil。

Go 没有传统“引用类型”，但切片、映射、通道等类型通过指针间接操作数据，行为上类似引用。
核心区别：Go 的“类似引用”类型是描述符或指针的复制，而非直接别名。

在 Go 中，make 是一个用于创建 slice（切片）、map（映射）和 channel（通道） 的内建函数。它用于初始化这些引用类型的数据结构，并分配底层存储。
*/
func SliceTutorial() {
	// init
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("slice1:", slice1)

	slice2 := make([]int, 5)
	fmt.Println("slice2:", slice2)

	var slice3 = []int{1, 2, 3, 4, 5}
	fmt.Println("slice3:", slice3)

	// append
	slice1 = append(slice1, 6)
	fmt.Println("slice1:", slice1)

	// copy
	slice4 := make([]int, 5)
	copy(slice4, slice2)
	fmt.Println("slice4:", slice4)

	// 截取
	subSlice1 := slice1[1:3]
	fmt.Println(subSlice1[:3]) // [1 2 3] (相当于 s[0:3])
	fmt.Println(subSlice1[2:]) // [3 4 5] (相当于 s[2:len(s)])
	fmt.Println(subSlice1[:])  // [1 2 3 4 5] (相当于 s[0:len(s)])
	fmt.Println("subSlice1:", subSlice1)

	// 遍历slice
	for i, v := range slice1 {
		fmt.Printf("index: %d, value: %d 地址：%p\n", i, v, &v)
	}

	var slice5 []int
	fmt.Println("slice5:", slice5)
	fmt.Println("slice5 is nil?", slice5 == nil) // true，nil 切片

	var slice6 = make([]int, 0)
	fmt.Println("slice6:", slice6)
	fmt.Println("slice6 is nil?", slice6 == nil) // false，空切片
}