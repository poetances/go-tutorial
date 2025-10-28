package process

import (
	"fmt"
	"os"
)

/// if 语句是用于条件判断的基本控制结构
/// 格式：if condition { // 条件为真时执行的代码 }
func IfTutorial() { 
	a := 10
	if a > 5 {
		fmt.Println("a 大于 5")
	} else {
		fmt.Println("a 小于或等于 5")
	}
}

/// Go 的 if 语句支持在条件判断前执行一个初始化语句，通常用于简化代码。
/*
if initialization; condition {
    // 条件为 true 时执行的代码
}
*/
func IfSpecialTutorial() { 
	// 初始化语句
	if x := 10; x > 5 {
		fmt.Println("x 大于 5")
	}

	// 初始化语句中可以有多个返回值
	if ok, err := someFunc(); ok {
		fmt.Println("操作成功")
	} else {
		fmt.Println("操作失败:", err)
	}

	// 变量作用域控制
	if file, error := os.Open("file.txt"); error != nil {
		fmt.Println("无法打开文件:", error)
	} else {
		defer file.Close() // 确保文件在使用后关闭
		fmt.Println("文件已成功打开")
	}

	// 多条件判断。if 中的分号最多只有一个， 多个会编译失败
	if num := 15; num >= 0 && num <= 10 {
		fmt.Println("num 在 0 到 10 之间")
	} else if num > 10 && num <= 20 {
		fmt.Println("num 在 11 到 20 之间")
	} else {
		fmt.Println("num 大于 20 或小于 0")
	}

	// 类型断言
	var i interface{} = "hello"
	if str, ok := i.(string); ok {
		fmt.Println("i 是一个字符串:", str)
	} else {
		fmt.Println("i 不是一个字符串")
	}
}
func someFunc() (bool, error) {
	return true, nil
}