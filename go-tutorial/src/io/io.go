package io

import (
	"fmt"
	"strings"
)

/*
Go的 io 包提供了I/O操作的基本接口，是Go语言中处理输入/输出的核心包之一。它定义了一系列接口，用于抽象不同类型的I/O操作，使代码更加灵活和可复用。

核心接口：
- Reader: 读取数据的接口
- Writer: 写入数据的接口
- Closer: 关闭资源的接口
*/

// 1.Reader 接口示例
/*
Reader 接口用于读取数据，定义如下：
type Reader interface {
    Read(p []byte) (n int, err error)
}
*/
func DemonstrateReader() {
    // 从字符串读取
    r := strings.NewReader("Hello, Go!")
	buf := make([]byte, 5)

    n, err := r.Read(buf)
    fmt.Printf("读取了 %d 字节: %s, 错误: %v\n", n, buf, err)

    // 继续读取
    n, err = r.Read(buf)
    fmt.Printf("又读取了 %d 字节: %s, 错误: %v\n", n, buf, err)

      // 再次读取，应该遇到EOF
    n, err = r.Read(buf)
    fmt.Printf("最后读取了 %d 字节: %s, 错误: %v\n", n, buf, err)
}

// 2.Writer 接口示例
/*
Writer 接口用于写入数据，定义如下：
type Writer interface {
    Write(p []byte) (n int, err error)
}
*/
func DemonstrateWriter() {
    var b strings.Builder

    n, err := b.Write([]byte("Hello, "))
    fmt.Printf("写入了 %d 字节, 错误: %v\n", n, err)

    n, err = b.Write([]byte("Go!"))
    fmt.Printf("又写入了 %d 字节, 错误: %v\n", n, err)

    fmt.Printf("最终内容: %s\n", b.String())
}
