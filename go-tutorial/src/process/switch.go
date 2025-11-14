package process
import "fmt"

/// author: zhuchaojun
/// date: 2024/6/12 15:28
/// description: SwitchTutorial

/*
Switch 基本语法和工作原理
switch 表达式 {
case 值1:
    // 执行语句1
case 值2, 值3:
    // 执行语句2
default:
    // 默认执行语句
}

go中的 switch 默认不会穿透，不需要增加 break
*/
func SwitchTutorial() {
	day := "monday"
	switch day {
	case "monday":
		fmt.Println("星期一")
	case "tuesday", "wednesday":
		fmt.Println("星期二或星期三")
	default:
		fmt.Println("其他日子")
	}
}

func InterfaceSwitch() {
	var value interface{} = "hello"
	switch v := value.(type) {
	case int:
		fmt.Println("整数:", v)
	case string:
		fmt.Println("字符串:", v)
	default:
		fmt.Println("未知类型")
	}
}

func NoexSwitch() {
	expression := 10
	switch {
	case expression > 0:
		fmt.Println("正数")
	case expression < 0:
		fmt.Println("负数")
	default:
		fmt.Println("零")
	}

	// 带初始化的switch
	switch a := 5; { 
	case a%2 == 0:
		fmt.Println("偶数")
	default:
		fmt.Println("奇数")
	}
}