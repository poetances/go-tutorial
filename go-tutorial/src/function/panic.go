package function

import (
	"errors"
	"fmt"

)

/*
Go 语言不像 Java、Python 那样提供完整的 try-catch-finally 异常处理机制，而是采用 返回错误值（error）+ panic/recover 方式来处理错误。

*/

// 1.1 使用 errors.New() 生成错误
func divid2(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a/b, nil
}

// 1.2 使用 fmt.Errorf() 创建带格式的错误
func openFile(fileName string) error {
	return fmt.Errorf("无法打卡文件:", fileName)
}

// 1.3 使用 errors.Is() 和 errors.As() 处理特定类型的错误
var errorNoFound = errors.New("资源未找到")
func findResource(id int) error {
	if id == 0 {
		return errorNoFound
	}
	return nil
}

/*
2. panic 和 recover 处理严重错误
Go 语言提供了 panic 和 recover，用于处理不可恢复的错误（如数组越界、空指针引用等）。

2.1 panic 触发严重错误
panic 用于中断程序执行，并输出错误信息。通常用于不可恢复的错误，例如：
	•	除零错误
	•	空指针引用
	•	超出数组范围
*/
func testPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到错误:", r)
		}
	}()

	fmt.Println("开始执行")
	panic("发生错误")
	fmt.Println("这里不会执行")
}

/*
如果不希望程序崩溃，可以使用 recover() 捕获 panic 并处理错误。

3. defer、panic 和 recover 结合使用
	•	defer 关键字可以在 panic 发生时执行清理操作，例如关闭文件、释放资源。
	•	recover 需要在 defer 中调用，才能正确捕获 panic。
*/

/*
语言	Go	Swift
可恢复错误	返回 error	do-catch
不可恢复错误	panic + recover	fatalError()
延迟执行	defer	defer
资源清理	defer	defer
*/
func PanicTutorial() {

	result, error := divid2(10, 0)
	if error != nil {
		fmt.Println("发生错误", error)
	} else {
		fmt.Println("计算结果", result)
	}

	err := findResource(0)
	if errors.Is(err, errorNoFound) {
		fmt.Println("错误: 资源未找")
	}
}
