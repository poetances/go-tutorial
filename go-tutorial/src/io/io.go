package io

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
	
}