package syntacticsugar

import (
	"fmt"
	"os"
	"time"
)

// 1.1 结构体字段访问语法糖
type PersonSugar struct {
	Name string
	Age  int
}
func PointerFieldAccess() {
	person := &PersonSugar{Name: "Alice", Age: 30}

	// 通过指针直接访问字段（语法糖）
	name := person.Name // 等价于 (*person).Name
	age := person.Age // 等价于 (*person).Age

	println("Name:", name)
	println("Age:", age)
}

// 1.2 指针调用方法
func (p *PersonSugar)SetName(name string) {
	p.Name = name
}
func (p PersonSugar)GetName() string {
	return p.Name
}
func (p PersonSugar)SetName2(name string) {
	p.Name = name // 会报警告， unused write to field Name
}
func PointMethodAccess() {
	person := &PersonSugar{Name: "Alice", Age: 30}
	person.SetName("Bob") 
	println("修改名字：%v", person.Name)

	person.SetName2("Charlie") // 这行代码不会改变person.Name的值

	println(person.GetName())// 编译器自动转换为 (*ptr).GetName()


	p := PersonSugar{Name: "李四"}
    // 语法糖：值类型调用指针接收者方法
    p.SetName("王五")  // 编译器自动转换为 (&p).SetName("王五")
	fmt.Println("值类型修改后的名字:", p.Name)
    
    // 语法糖：指针调用值接收者方法
    ptr := &p
    name := ptr.GetName()  // 编译器自动转换为 (*ptr).GetName()
    fmt.Println(name)
}

// 2.变量声明和初始化语法糖
// 2.1 简短变量声明语法糖
func ShortVariableDeclaration() {
	x := 10          // 等价于 var x int = 10
	name := "Alice"  // 等价于 var name string = "Alice"
	isActive := true // 等价于 var isActive bool = true
	
	fmt.Println(x, name, isActive)
}

// 2.2 结构体字面量语法糖
func StructLiteralSugar() {
	// 语法糖：字段名可以省略
	p1 := PersonSugar{Name: "Alice", Age: 30} // 标准写法
	p2 := PersonSugar{"Bob", 25} // 等于 PersonSugar{Name: "Bob", Age: 25}
	
	p3 := &PersonSugar{Name: "Charlie", Age: 28} // 指针类型的结构体字面量
	p4 := &PersonSugar{"Bob", 25} // 等于 PersonSugar{Name: "Bob", Age: 25}
	fmt.Println(p1, p2, p3, p4)
}

// 2.3 数组和切片字面量语法糖
func ArrayAndSliceLiteralSugar() {	
	arr := [3]int{1, 2, 3}       // 数组字面量

	slice := []int{4, 5, 6}     // 切片字面量

	arr2 := [...]int{7, 8, 9}    // 编译器自动推导数组长度
	fmt.Println(arr, slice, arr2)
}

// 3.函数相关语法糖
// 3.1 多返回值语法糖
func mutipleRetruns() (int, string, bool) {
	return 1, "hello", true
}
func ConsumeMultipleReturns() { 
	a, b, c := mutipleRetruns() // 语法糖：多返回值赋值
	fmt.Println(a, b, c)

	a1, _, c1 := mutipleRetruns() 
	fmt.Println(a1, c1)
}

// 4.控制结构语法糖
// 4.1 if 语句中的初始化
func IfWithInit() {
	if x := 20; x > 10 {
		fmt.Println("x 大于 10:", x)
	}
	/*
	等价于
	x := 20
	if x > 10 {
		fmt.Println("x 大于 10:", x)
	}
	*/

	// 常用于错误处理
	if file, err := os.Open("file.txt"); err != nil {
		fmt.Println("无法打开文件:", err)
	} else {
		defer file.Close()
		fmt.Println("文件打开成功")
	}
}
// 4.2 switch 语句中的初始化
func SwitchWithInit() {
	switch day := "Monday"; day {
	case "Monday":
		fmt.Println("今天是星期一")
	case "Tuesday":
		fmt.Println("今天是星期二")
	default:
		fmt.Println("不是星期一或星期二")
	}

	// type switch
	var i interface{} = "hello"
	switch v := i.(type) {
	case int:
		fmt.Println("i 是 int 类型:", v)
	case string:
		fmt.Println("i 是 string 类型:", v)
	default:
		fmt.Println("未知类型")
	}

	v1, ok := i.(string)
	if ok {
		fmt.Println("i 是 string 类型:", v1)
	}
}

// 5. 类型语法糖
// 5.1 类型别名和类型定义
func TypeAliasesAndDefinitions() {
	type MyInt int // 类型定义
	type MyString = string // 类型别名

	var i MyInt = 10
	var s MyString = "hello"
	fmt.Println(i, s)
}

// 5.2 匿名结构体和接口
func AnonymousType() { 
	person := struct {
		string
		int
	} {"John", 30} // 理论上是不建议这样，不带标签

	fmt.Println(person.int, person.string)

	var writer interface {
		Write([]byte) (int, error)
	} = os.Stdout

	_, err := writer.Write([]byte("Hello, World!\n"))
	if err != nil {
		fmt.Println("写入错误:", err)
	}
}

// 6. 切片和映射相关语法糖
// 6.1 切片扩展语法糖
func SliceExtensionSugar() {
	slice := []int{1, 2, 3}
	sub1 := slice[1:] 
	sub2 := slice[:2]
	sub3 := slice[1:2]
	sub4 := slice[:]
	fmt.Println(sub1, sub2, sub3, sub4)

	slice = append(slice, 4, 5, 6) // 语法糖：append 函数
	fmt.Println("扩展后的切片:", slice)

	more := []int{7, 8, 9}
	slice = append(slice, more...)
	fmt.Println("合并切片后的结果:", slice)	
}

// 6.2 映射字面量语法糖
func MapLiteralSugar() {
	m := map[string]int {
		"apple":  5,
		"banana": 10,
		"orange": 15,
	}

	fmt.Println(m)

	if val, ok := m["banana"]; ok {
		fmt.Println("香蕉的数量:", val)
	}

	// 删除键
	delete(m, "banana")
	fmt.Println("删除香蕉后的映射:", m)
}

// 7. 并发的语法糖
// 7.1 select 语句语法糖
func SelectStatement() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        ch1 <- "来自 ch1"
    }()
    
    go func() {
        ch2 <- "来自 ch2"
    }()
    
    // 语法糖：select 等待多个通道
    select {
    case msg1 := <-ch1:
        fmt.Println("收到:", msg1)
    case msg2 := <-ch2:
        fmt.Println("收到:", msg2)
    case <-time.After(1 * time.Second):
        fmt.Println("超时")
    }
}

// 7.2 range 遍历通道
func RangeChannel() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
    
    // 语法糖：range 遍历通道
    for value := range ch {
        fmt.Println(value)
    }
}

// 8. 字符串和字符的语法糖
// 8.1 字符串字面量
func StringLiterals() {
    // 语法糖：原始字符串字面量
    raw := `这是一个
多行
字符串`
    
    // 语法糖：字符串格式化
    name := "张三"
    age := 30
    message := fmt.Sprintf("姓名: %s, 年龄: %d", name, age)
    
    fmt.Println(raw)
    fmt.Println(message)
}

// 8.2 for range 遍历字符串
func RangeString() {
    str := "hello, 世界"
    
    // 语法糖：for range 遍历字符串（处理 Unicode）
    for i, r := range str {
        fmt.Printf("位置 %d: %c\n", i, r)
    }
}

// 9. 接口的语法糖
// 9.1 空接口
func EmptyInterface() {
    // 语法糖：空接口可以接受任何类型
    var data interface{} = 42
    data = "hello"
    data = []int{1, 2, 3}
	
    fmt.Println(data)
}

// 9.2 接口断言
func InterfaceAssertion() {
    var data interface{} = "hello"
    
    // 语法糖：类型断言
    if str, ok := data.(string); ok {
        fmt.Println("字符串:", str)
    }
    
    // 语法糖：类型 switch
    switch v := data.(type) {
    case string:
        fmt.Println("字符串:", v)
    case int:
        fmt.Println("整数:", v)
    }
}

// 10. 错误处理的语法糖
func ErrorWrapping() {
    // Go 1.13+ 的错误包装语法糖
    _, err := os.Open("nonexistent.txt")
    if err != nil {
        // 语法糖：错误包装
        wrappedErr := fmt.Errorf("打开文件失败: %w", err)
        
        // 语法糖：错误解包
        if os.IsNotExist(wrappedErr) {
            fmt.Println("文件不存在")
        }
    }
}