package function

import (
	"fmt"
	"time"
)

// FunctionTutorial Go 语言函数详细教程
func FunctionTutorial() {
	fmt.Println("=== Go 语言函数详细教程 ===")
	
	// 1. 基本函数定义和调用
	basicFunctionExamples()
	
	// 2. 多返回值函数
	multipleReturnValues()
	
	// 3. 命名返回值
	namedReturnValues()
	
	// 4. 可变参数函数
	variadicFunctions()
	
	// 5. 函数作为一等公民
	firstClassFunctions()
	
	// 6. 闭包和匿名函数
	closuresAndAnonymousFunctions()
	
	// 7. 延迟执行 (defer)
	deferExamples()
	
	// 8. 错误处理模式
	errorHandlingPatterns()
	
	// 9. 性能优化技巧
	performanceOptimizations()
	
	// 10. 实际开发最佳实践
	bestPractices()
}

// ==================== 1. 基本函数定义和调用 ====================

// basicFunctionExamples 基本函数示例
func basicFunctionExamples() {
	fmt.Println("\n1. 基本函数定义和调用:")
	
	// 无参数无返回值
	sayHello()
	
	// 有参数有返回值
	result := add(10, 20)
	fmt.Printf("add(10, 20) = %d\n", result)
	
	// 参数类型相同的简写
	result2 := addThree(1, 2, 3)
	fmt.Printf("addThree(1, 2, 3) = %d\n", result2)
}

// sayHello 无参数无返回值函数
func sayHello() {
	fmt.Println("Hello, Go!")
}

// add 两个整数相加
func add(a int, b int) int {
	return a + b
}

// addThree 三个整数相加（参数类型相同可简写）
func addThree(a, b, c int) int {
	return a + b + c
}

// ==================== 2. 多返回值函数 ====================

// multipleReturnValues 多返回值示例
func multipleReturnValues() {
	fmt.Println("\n2. 多返回值函数:")
	
	// 基本多返回值
	quotient, remainder := divide(10, 3)
	fmt.Printf("divide(10, 3) = 商:%d, 余数:%d\n", quotient, remainder)
	
	// 忽略部分返回值
	q, _ := divide(15, 4) // 忽略余数
	fmt.Printf("divide(15, 4) = 商:%d (忽略余数)\n", q)
	
	// 交换两个值
	x, y := swap(100, 200)
	fmt.Printf("swap(100, 200) = %d, %d\n", x, y)
}

// divide 除法运算，返回商和余数
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// swap 交换两个字符串
func swap(x, y int) (int, int) {
	return y, x
}

// ==================== 3. 命名返回值 ====================

// namedReturnValues 命名返回值示例
func namedReturnValues() {
	fmt.Println("\n3. 命名返回值:")
	
	// 命名返回值
	sum, diff := calc(20, 10)
	fmt.Printf("calc(20, 10) = 和:%d, 差:%d\n", sum, diff)
	
	// 复杂计算示例
	area, perimeter := rectangle(5.0, 3.0)
	fmt.Printf("rectangle(5.0, 3.0) = 面积:%.1f, 周长:%.1f\n", area, perimeter)
}

// calc 计算和与差（命名返回值）
func calc(x, y int) (sum int, difference int) {
	sum = x + y
	difference = x - y
	return // 裸返回，自动返回命名返回值
}

// rectangle 计算矩形面积和周长
func rectangle(length, width float64) (area float64, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 裸返回
}

// ==================== 4. 可变参数函数 ====================

// variadicFunctions 可变参数函数示例
func variadicFunctions() {
	fmt.Println("\n4. 可变参数函数:")
	
	// 基本可变参数
	sum1 := sum(1, 2, 3)
	sum2 := sum(1, 2, 3, 4, 5)
	fmt.Printf("sum(1, 2, 3) = %d\n", sum1)
	fmt.Printf("sum(1, 2, 3, 4, 5) = %d\n", sum2)
	
	// 切片展开为可变参数
	numbers := []int{10, 20, 30, 40}
	sum3 := sum(numbers...) // 切片展开
	fmt.Printf("sum([]int{10, 20, 30, 40}...) = %d\n", sum3)
	
	// 混合参数
	result := joinWithSeparator(", ", "Go", "Python", "Java")
	fmt.Printf("joinWithSeparator(\", \", \"Go\", \"Python\", \"Java\") = %s\n", result)
}

// sum 可变参数求和
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// joinWithSeparator 使用分隔符连接字符串
func joinWithSeparator(separator string, strings ...string) string {
	result := ""
	for i, s := range strings {
		if i > 0 {
			result += separator
		}
		result += s
	}
	return result
}

// ==================== 5. 函数作为一等公民 ====================

// firstClassFunctions 函数作为一等公民示例
func firstClassFunctions() {
	fmt.Println("\n5. 函数作为一等公民:")
	
	// 函数类型定义
	type MathFunc func(int, int) int
	
	// 函数赋值给变量
	var addFunc MathFunc = func(a, b int) int {
		return a + b
	}
	
	var multiplyFunc MathFunc = func(a, b int) int {
		return a * b
	}
	
	// 函数作为参数
	result1 := calculate(10, 5, addFunc)
	result2 := calculate(10, 5, multiplyFunc)
	fmt.Printf("calculate(10, 5, add) = %d\n", result1)
	fmt.Printf("calculate(10, 5, multiply) = %d\n", result2)
	
	// 函数作为返回值
	incrementer := createIncrementer(5)
	fmt.Printf("incrementer(10) = %d\n", incrementer(10))
}

// calculate 函数作为参数
func calculate(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// createIncrementer 函数作为返回值
func createIncrementer(increment int) func(int) int {
	return func(x int) int {
		return x + increment
	}
}

// ==================== 6. 闭包和匿名函数 ====================

// closuresAndAnonymousFunctions 闭包和匿名函数示例
func closuresAndAnonymousFunctions() {
	fmt.Println("\n6. 闭包和匿名函数:")
	
	// 立即执行函数
	func()  {
		fmt.Println("立即执行的匿名函数")
	}()
	
	// 闭包：捕获外部变量
	counter := createCounter()
	fmt.Printf("counter() = %d\n", counter()) // 1
	fmt.Printf("counter() = %d\n", counter()) // 2
	fmt.Printf("counter() = %d\n", counter()) // 3
	
	// 闭包：状态保持
	sequencer := createSequencer()
	fmt.Printf("sequencer() = %d\n", sequencer()) // 1
	fmt.Printf("sequencer() = %d\n", sequencer()) // 2
	fmt.Printf("sequencer() = %d\n", sequencer()) // 3
}

// createCounter 创建计数器闭包
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// createSequencer 创建序列生成器
func createSequencer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// ==================== 7. 延迟执行 (defer) ====================

// deferExamples defer 示例
func deferExamples() {
	fmt.Println("\n7. 延迟执行 (defer):")
	
	// 基本 defer
	deferBasic()
	
	// defer 执行顺序（LIFO）
	deferOrder()
	
	// defer 与返回值
	result := deferWithReturn()
	fmt.Printf("deferWithReturn() = %d\n", result)
	
	// defer 在循环中的使用
	deferInLoop()
}

// deferBasic 基本 defer 使用
func deferBasic() {
	fmt.Println("开始执行 deferBasic")
	defer fmt.Println("defer 语句 1")
	defer fmt.Println("defer 语句 2")
	fmt.Println("结束执行 deferBasic")
	// 输出顺序: 结束执行 deferBasic, defer 语句 2, defer 语句 1
}

// deferOrder defer 执行顺序
func deferOrder() {
	fmt.Println("defer 执行顺序:")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer %d\n", i)
	}
	fmt.Println("循环结束")

	// 输出顺序: 循环结束, defer 2, defer 1, defer 0
}

// deferWithReturn defer 与返回值
func deferWithReturn() (result int) {
	defer func() {
		result++ // 修改命名返回值
		fmt.Printf("defer 中修改返回值: %d\n", result)
	}()
	
	result = 10
	fmt.Printf("函数中设置返回值: %d\n", result)
	return result
}

// deferInLoop defer 在循环中的使用
func deferInLoop() {
	fmt.Println("循环中的 defer:")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in loop: %d\n", i)
	}
}

// ==================== 8. 错误处理模式 ====================

// errorHandlingPatterns 错误处理模式示例
func errorHandlingPatterns() {
	fmt.Println("\n8. 错误处理模式:")
	
	// 基本错误处理
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Printf("错误: %v\n", err)
	} else {
		fmt.Printf("safeDivide(10, 2) = %d\n", result)
	}
	
	// 除零错误
	_, err2 := safeDivide(10, 0)
	if err2 != nil {
		fmt.Printf("错误: %v\n", err2)
	}
	
	// 多返回值错误处理
	fileContent, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("文件读取错误: %v\n", err)
	} else {
		fmt.Printf("文件内容: %s\n", fileContent)
	}
}

// safeDivide 安全的除法运算
func safeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// readFile 模拟文件读取
func readFile(filename string) (string, error) {
	if filename == "" {
		return "", fmt.Errorf("文件名不能为空")
	}
	// 模拟文件读取
	return "文件内容", nil
}

// ==================== 9. 性能优化技巧 ====================

// performanceOptimizations 性能优化技巧
func performanceOptimizations() {
	fmt.Println("\n9. 性能优化技巧:")
	
	// 1. 避免不必要的函数调用
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		// 内联小函数可以提高性能
		_ = i * i
	}
	fmt.Printf("内联计算耗时: %v\n", time.Since(start))
	
	// 2. 使用指针避免大结构体复制
	type LargeStruct struct {
		Data [1000]int
	}
	
	large := LargeStruct{}
	
	// 值传递（复制整个结构体）
	start1 := time.Now()
	processByValue(large)
	fmt.Printf("值传递耗时: %v\n", time.Since(start1))
	
	// 指针传递（只复制指针）
	start2 := time.Now()
	// processByPointer(&large)
	fmt.Printf("指针传递耗时: %v\n", time.Since(start2))
}

func processByValue(s struct{ Data [1000]int }) {
	// 模拟处理
	_ = s.Data[0]
}


// ==================== 10. 实际开发最佳实践 ====================

// bestPractices 最佳实践
func bestPractices() {
	fmt.Println("\n10. 实际开发最佳实践:")
	
	fmt.Println("✅ 1. 函数命名:")
	fmt.Println("   - 使用驼峰命名法")
	fmt.Println("   - 函数名应该描述其行为")
	fmt.Println("   - 避免过于通用的名称")
	
	fmt.Println("\n✅ 2. 参数设计:")
	fmt.Println("   - 参数数量控制在 3-4 个以内")
	fmt.Println("   - 相关参数使用结构体封装")
	fmt.Println("   - 使用选项模式处理可选参数")
	
	fmt.Println("\n✅ 3. 错误处理:")
	fmt.Println("   - 总是检查并处理错误")
	fmt.Println("   - 提供有意义的错误信息")
	fmt.Println("   - 使用自定义错误类型")
	
	fmt.Println("\n✅ 4. 文档注释:")
	fmt.Println("   - 为公共函数编写文档注释")
	fmt.Println("   - 说明参数、返回值和副作用")
	fmt.Println("   - 提供使用示例")
	
	fmt.Println("\n✅ 5. 单一职责:")
	fmt.Println("   - 每个函数只做一件事")
	fmt.Println("   - 保持函数简洁（通常不超过 50 行）")
	fmt.Println("   - 复杂的逻辑拆分成多个小函数")
	
	fmt.Println("\n✅ 6. 测试友好:")
	fmt.Println("   - 函数应该是可测试的")
	fmt.Println("   - 避免隐藏的依赖")
	fmt.Println("   - 使用接口而不是具体实现")
}

// ==================== 原有函数保留 ====================

// Maximum 比较两个数的大小
func Maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Swap 交换两个字符串
func Swap(x, y string) (string, string) {
	return y, x
}

// Calc 计算和与差（命名返回值）
func Calc(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

// Calc2 计算和与差（命名返回值简写）
func Calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return sum, sub
}

/*
func example() {
    // 这个 recover 不会捕获 panic，因为它不在 defer 中
    if r := recover(); r != nil {
        fmt.Println("这不会捕获 panic")
    }
    
    panic("发生错误")
    
    // 这行不会执行，因为 panic 已经发生
    if r := recover(); r != nil {
        fmt.Println("这也不会捕获 panic")
    }
}
这是一个错误的示例，recover需要在defer函数中使用，才能捕获 panic。
*/
func PanicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()
	panic("This is error")
}
