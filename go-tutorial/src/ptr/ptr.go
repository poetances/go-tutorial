package ptr

import "fmt"

/*
指针是一个存储另一个变量内存地址的变量。在 Go 中，指针类型表示为 *T ，其中 T 是指针指向的变量类型。

*/
func BasicPointerOperation() {
	num := 42
	// 1. 取地址操作符 &
	var ptr *int = &num // 获取 num 的地址
	fmt.Printf("变量值：%d\n", num)
	fmt.Printf("变量地址：%p\n", &num)
	fmt.Printf("指针值存储的地址：%p\n", ptr)
	 // 2. 解引用操作符 *
	fmt.Printf("指针指向的值：%d\n", *ptr)
    // 3. 通过指针修改值
	*ptr = 100 
	fmt.Printf("修改后的值：%d\n", num)

}

func PointerZeroValue() {
	var ptr *int // 声明但为初始化，指针的零值是 nil
	if ptr == nil {
		fmt.Println("指针的零值是 nil")
	}

	fmt.Println(ptr)
	// 尝试解引用 nil 指针会导致 panic
	// fmt.Println(*ptr) // 这会导致 panic: runtime error: invalid memory address
}

func CreatePointerWithNew() {
	ptr := new(int) // 使用 new 函数创建指针，指向一个 int 类型的零值

	fmt.Printf("new() 创建的指针：%p\n", ptr)
	fmt.Printf("指针指向的值：%d\n", *ptr)

	*ptr = 50  // 设置值
    fmt.Printf("设置后的值: %d\n", *ptr) 
}

// 3. 指针的使用场景
// 3.1 修改函数外部的变量
func ModifyByPointer() {
	x := 10
	fmt.Printf("修改前 x 的值: %d\n", x)
	modifyValue(&x) // 传递 x 的地址
	fmt.Printf("修改后 x 的值: %d\n", x)
}
func modifyValue(num *int) {
	*num = 200
}
// 3.2 避免大对象的拷贝
type LargeStruct struct {
	data [100000]int
	name string
}
func processLargeStruct(ls LargeStruct) {
	fmt.Println(ls.data)
	// 这里会拷贝整个 LargeStruct，消耗内存和时间
}
func processLargePointer(ls *LargeStruct) {
	ls.name = "Processed"
	 // 只传递 8 字节（64位系统）的指针，高效
}
func DemonstrateEfficiency() {
	ls := LargeStruct{name: "大数据"}
	// 值传递
	processLargeStruct(ls)
	// 指针传递
	processLargePointer(&ls)
	fmt.Printf("处理后的结构体名称: %s\n", ls.name)
}

//指针的高级用法
//指针与切片
func PointerWithSlice() {
	slice := []int{1, 2, 3}
	fmt.Printf("修改前切片: %v\n", slice)
	modifySlice(slice)
	fmt.Printf("修改后切片: %v\n", slice)

	modifySlicePointer(&slice)
	fmt.Printf("修改后切片: %v\n", slice)
}
func modifySlice(s []int) {
	s[0] = 100
}
func modifySlicePointer(s *[]int) {
	*s = append(*s, 4, 5)
}

// 4.3 指针与函数
func PointerWithFunction() {
	num := 50
	fmt.Printf("修改前 num 的值: %d\n", num)
	modifyUsingPointer(&num)
	fmt.Printf("修改后 num 的值: %d\n", num)
}
func modifyUsingPointer(num *int) {
	*num = 100
}

//5.2 避免返回局部变量的指针
func ReturnPointerTutorial() {
	fmt.Println("避免返回局部变量的指针")
	
	// 错误示例
	ptr := wrongReturnPointer()
	fmt.Println(*ptr)  // 可能导致未定义行为
	
	// 正确示例
	ptr1 := correctReturnPointer()
	fmt.Println(*ptr1)  // 输出: 42
}
// 错误：返回局部变量的指针
func wrongReturnPointer() *int {
    num := 42
    return &num  // 危险：num 是局部变量，函数返回后会被销毁
}

// 正确：使用 new 创建的变量
func correctReturnPointer() *int {
    num := new(int)  // 在堆上分配内存
    *num = 42
    return num       // 安全：num 在堆上，不会被销毁
}

// 5.3指针与垃圾回收
func PointerAndGC() {
    // 循环引用示例
    type Node struct {
        name string
        next *Node
    }
    
    node1 := &Node{name: "节点1"}
    node2 := &Node{name: "节点2"}
    
    node1.next = node2
    node2.next = node1  // 形成循环引用
    
    // Go 的垃圾回收器可以处理循环引用
    // 但仍要注意避免不必要的指针引用
}

// 何时使用指针
func WhenToUsePointer() {
	// 1. 修改函数外部的变量
	modifyExternal := func (num *int) {
		*num = 100
	}
	num := 50
	modifyExternal(&num)
	fmt.Println(num)

	// 2. 避免大对象的拷贝
	type BigStruct struct {
		data [100000]int
	}
	processBigStruct := func (bs *BigStruct) {
		bs.data[0] = 1 // 这里其实是 go 的语法糖，等价于 (*bs).data[0] = 1
	}
	processBigStruct1 := func (bs BigStruct) {
		bs.data[0] = 1
	}
	bs := BigStruct{}
	processBigStruct(&bs)
	fmt.Println(bs.data[0])

	bs1 := BigStruct{}
	processBigStruct1(bs1)
	fmt.Println(bs1.data[0])
}