package variable

import (
	"fmt"
)

/*
常量（字面量）
常量是一个简单值的标识符，在程序运行时，不会被修改的量。
常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。
常量可以是字符、字符串、布尔值或数值。

var identifier type
var identifier1, identifier2 type


Author: zcj
*/

func VarTutorial() {
	var i int
	var f float64 = 12.0
	var b bool = true
	var c complex64 = 5 + 5i
	var bb byte = 255

	var arr [3]int = [3]int{1, 2}
	fmt.Println(i, f, b, c, bb, arr)
}

func VarTutorial2() {
	ii := 13
	var f float64 = 12

	var c = 12
	const d = 12
	const e = 12.0
	const n = "name"
	const numExpression = 1 + 2

	// 批量声明
	var (
		age  = 12
		name = "name"
	)

	const (
		age2  = 12
		name2 = "name"
	)

	// iota 常量重置，一般用于枚举
	const iota = 0
	const (
		A = iota
		B
	)

	// 与表达式结合。这种写法，var声明不支持，var声明必须有初始值
	const (
		KB = 1 << (10 * iota) // 1 << (10 * 1)
		MB                    // 1 << (10 * 2)
		GB                    // 1 << (10 * 3)
	)

	// 掩码
	const (
		Read = 1 << iota
		Write
	)
	fmt.Println(ii, f, c, d, e, n, numExpression, age, name, age2, name2, A, B, KB, MB, GB, Read, Write)
}

func VarTutorial3() {
	// 变量
	var intNum int
	var str string
	var char byte

	intNum = 12
	str = "name"
	char = 'a'

	// 批量声明
	var name string
	var age int
	name, age = "David", 12

	// 也可以。去掉var和后置类型。短变量声明
	name1, age1 := "David", 12

	// 交换变量，不需要指针，非常直观.
	// 短变量声明，只能用于函数内部
	num1, num2 := 25, 36
	num1, num2 = num2, num1

	// 忽略类型。如果有些变量不使用，可以忽略类型。
	a, b, _ := 1, 2, 3

	fmt.Println(intNum, str, char, name, age, name1, age1, num1, num2, a, b)
}

func VarTutorial4() {
	var a, b string
	a = "Hello"
	b = "Wrold"
	fmt.Println(a, b)
	fmt.Printf("%v", a)

	var (
		c = "Hello"
		d = "World"
	)
	fmt.Println(c, d)
}