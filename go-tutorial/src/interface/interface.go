package interfacepackage

import (
	"fmt"
	"reflect"
)

/*
在 Go 语言中，接口是一种抽象类型，它定义了一组方法签名。任何类型只要实现了接口中定义的所有方法，就被视为实现了该接口。
*/
// 1. 接口定义
type Writer interface {
	Write(data []byte) (int, error)
}
type Reader interface {
	Read(data []byte) (int, error)
}
type ReadWriter interface {
	Reader
	Writer
}

// 1.2 接口的隐式实现. Go语言中，接口是隐式实现的，不需要显示声明某个接口
type File struct {
	name string
}
// File实现了 Writer 接口
func (f File)Write(data []byte) (int, error) {
	fmt.Printf("写入文件 %s: %s\n", f.name, string(data))
	return len(data), nil
}
// File 实现了 Reader接口
func (f File)Read(data []byte) (int, error) {
	fmt.Printf("从文件 %s 读取数据\n", f.name)
	return len(data), nil
}
func DemonstrateImplicitImplementation() {
	file := File{name: "test.txt"}

	var rw ReadWriter = file
	_, err := rw.Read([]byte("hello"))
	_, err1 := rw.Write([]byte("world"))
	if err != nil || err1 != nil {
		fmt.Println("操作失败")
	} else {
		fmt.Println("操作成功")
	}
}

// 2. 接口的类型和值
// 2.1 接口的内部表示。Go 中的接口在内部由两个字段组成：类型和值
func DemonstrateInterfaceInternal() {
	var w Writer
	fmt.Printf("空接口：类型=%v, 值=%v, 是否为nil=%t\n", reflect.TypeOf(w), reflect.ValueOf(w), w==nil)

	file := File{name: "test.txt"}
	w = file
	fmt.Printf("非空接口：类型=%v, 值=%v, 不等于 nil\n", reflect.TypeOf(w), reflect.ValueOf(w))

	// 将接口设置 nil
	w = nil
	fmt.Printf("nil接口：类型=%v, 值=%v, 是否为nil=%t\n", reflect.TypeOf(w), reflect.ValueOf(w), w==nil)
}

// 2.2 空接口. interface{}没有任何方法，所有类型都实现了空接口
func DemonstrateEmptyInterface() {
	  // 空接口可以存储任何类型的值
    var data interface{}
    
    data = 42
    fmt.Printf("整数: %v (类型: %T)\n", data, data)
    
    data = "hello"
    fmt.Printf("字符串: %v (类型: %T)\n", data, data)
    
    data = []int{1, 2, 3}
    fmt.Printf("切片: %v (类型: %T)\n", data, data)
    
    // 空接口的实际应用
    // 1. 作为函数参数，接受任意类型
    printAnything(42)
    printAnything("hello")
    printAnything([]string{"a", "b", "c"})
    
    // 2. 作为容器元素，存储不同类型
    items := []interface{}{1, "two", true, 3.14}
    fmt.Printf("混合类型切片: %v\n", items)
    
    // 3. 作为 map 的值，存储不同类型
    config := map[string]interface{}{
        "host":    "localhost",
        "port":    8080,
        "debug":   true,
        "timeout": 30.5,
    }
    fmt.Printf("配置: %+v\n", config)
}		
func printAnything(value interface{}) {
    fmt.Printf("值: %v, 类型: %T\n", value, value)
}

// 3. 类型断言
// 3.1 基本类型断言。类型断言用于检查接口变量是否实现了某个具体的类型，如果实现了，则可以安全地转换为该类型。	
func DemonstrateTypeAssertion() {
	var value interface{} = "hello"

	if str, ok := value.(string); ok {
		fmt.Println("字符串:", str)
	}

	 // 尝试断言为不匹配的类型
    if num, ok := value.(int); ok {
        fmt.Printf("整数值: %d\n", num)
    } else {
        fmt.Println("不是整数类型")
    }
}
// 3.2 类型 switch
func DemonstrateTypeSwitch() {
	var value interface{} = "hello"	
	switch v := value.(type) {
	case string:	
		fmt.Println("字符串:", v)
	case int:
		fmt.Printf("整数: %d\n", v)
	default:
		fmt.Println("未知类型")
	}
}