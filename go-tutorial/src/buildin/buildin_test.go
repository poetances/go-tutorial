package buildin

import (
	"fmt"
	"testing"
)

/*
你在 Go 语言的 builtin.go文件中观察到的这个细节非常关键，它直接指向了 type关键字两种用途的核心区别：​类型定义​ 和 ​类型别名。

简单来说，有等号（=）的是类型别名，没有等号的是类型定义。这两者在行为和用途上有本质的不同。

它们的区别：
type NewType OldType  创建了一个全新的类型​, 与原始类型不兼容，需显式转换
type Alias = OldType  仅仅是现有类型的一个别名​, 与原始类型完全等价，可直接互换使用
*/ 

type MyInt int
type MyEqualInt = int

// 这种机制非常有用，它允许你基于基础类型创建出具有特定语义和限制的类型，增强代码的类型安全和可读性。例如，你可以定义 type Celsius float64和 type Fahrenheit float64来区分摄氏度和华氏度，避免不小心将两者直接相加。
type Celsius float64
type Fahrenheit float64

// 当你使用等号时，例如 type MyAlias = int，你并没有创建新类型。MyAlias只是 int的一个别名或代号，它们代表的是内存中的同一种类型。

func TestBuildin(t *testing.T) {
	var a MyInt = 10
	var b MyEqualInt = 10

	fmt.Println(a, b)
	fmt.Printf("a is %T, b is %T\n", a, b) // a is buildin.MyInt, b is int

	var i = 10
	// var j MyInt = i 会报错，因为 MyInt 和 int 是两个不同的类型
	var j MyEqualInt = i
	fmt.Println(j)
}

// buildin中有很多方法，后面可以再实现
// The len built-in function returns the length of v, according to its type:
//
//   - Array: the number of elements in v.
//   - Pointer to array: the number of elements in *v (even if v is nil).
//   - Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
//   - String: the number of bytes in v.
//   - Channel: the number of elements queued (unread) in the channel buffer;
//     if v is nil, len(v) is zero.
//
// For some arguments, such as a string literal or a simple array expression, the
// result can be a constant. See the Go language specification's "Length and
// capacity" section for details.
//func len(v Type) int
// 注意 len 上面的解释，如果是字符串的是，返回的是字节数