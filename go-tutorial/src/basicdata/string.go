package basicdata

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
string虽然是值类型，但是内部是引用类型机制
基本特效：
1. 不可变性：字符串一旦创建，其内容不可更改。任何对字符串的修改操作实际上都会创建一个新的字符串。
2. 底层实现：string底层是一个只读的[]byte字节切片。存储的是 UTF-8 编码的字节序列。
3. 零值：string的零值是空字符串""，而不是nil。

常见的操作(strings包中)：
(1) 字符串拼接
	可以使用 + 运算符或 fmt.Sprintf() 函数进行字符串拼接。 如果有大量字符串的拼接，推荐使用 strings.Builder 来提高性能。
(2) 字符串长度
	len(s) ：返回字节长度（不是字符数）。
	utf8.RuneCountInString(s) ：返回字符数（支持 Unicode）。
(3) 字符串遍历
	按字节遍历：	
	s := "Hello"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i]) // 输出: H e l l o
	}
	按字符遍历（推荐）：
	s := "世界"
	for _, r := range s {
		fmt.Printf("%c ", r) // 输出: 世 界
	}
*/ 
func StringTutorial() {

	str1 := "this is str"
	fmt.Println(str1)

	str2 := `
	this is duo 
	this.is
	`
	fmt.Println(str2)

	// 1、长度
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
	lenstr2 := "你好，世界🌍"
	strlen := len(lenstr2)
	fmt.Println(" 字节长度lenstr2:", strlen) // 字节长度lenstr2: 19.因为汉字是 3 个字节，英文是 1 个字节，表情是 4 个字节

	countstr2 := utf8.RuneCountInString(lenstr2)
	fmt.Println(" 字符长度lenstr2:", countstr2) // 字符长度lenstr2: 6

	// 2、拼接
	str3 := str1 + str2
	fmt.Println("拼接str3：", str3)

	
	str3 = fmt.Sprintf("%s%s", str1, str2)
	fmt.Println("拼接str3：", str3)

	// 大量拼接，推荐使用 strings.Builder
	var builder strings.Builder
	builder.WriteString(str1)
	builder.WriteString(str2)
	str4 := builder.String()
	fmt.Println("拼接str4：", str4)

	// 3、遍历
	// 按字节遍历
	for i := 0; i < len(lenstr2); i++ {
		fmt.Printf("字节遍历:%c \n", lenstr2[i])		
	}
	// 按字符遍历
	for _, r := range lenstr2 {
		fmt.Printf("字符遍历:%c \n", r)
	}


	arr := strings.Split(str1, "-")
	fmt.Println("切割后的数组：", arr)

	joinstring := strings.Join(arr, "-")
	fmt.Println(joinstring)

	isContain := strings.Contains(str1, str2)
	fmt.Println(isContain)
	
}