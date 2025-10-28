package basicdata

import (
	"fmt"
	"sync"
	"time"
)

/*
在 Go 语言中， channel 是一种用于在不同 Goroutine 之间进行通信和同步的核心机制。它基于 CSP（Communicating Sequential Processes）模型，强调通过通信共享内存，而不是通过共享内存来通信。
以下是 channel 的详细讲解，包括原理、使用方法以及常用技巧。
channel 的基本原理
(1) 底层数据结构
channel 的底层是一个环形队列（ hchan 结构体），包含以下关键字段：
buf ：指向缓冲区的指针（有缓冲 channel 时使用）。
sendx 和 recvx ：发送和接收的索引位置。
lock ：互斥锁，保证并发安全。
sendq 和 recvq ：等待发送和接收的 Goroutine 队列（阻塞时使用）。

type hchan struct {
    qcount   uint           // 当前队列中的元素数量
    dataqsiz uint           // 环形缓冲区的大小（有缓冲 channel）
    buf      unsafe.Pointer // 指向缓冲区的指针
    elemsize uint16         // 元素的大小
    closed   uint32         // 是否已关闭
    elemtype *_type         // 元素的类型信息
    sendx    uint           // 发送索引（环形缓冲区）
    recvx    uint           // 接收索引（环形缓冲区）
    recvq    waitq          // 等待接收的 Goroutine 队列
    sendq    waitq          // 等待发送的 Goroutine 队列
    lock     mutex          // 互斥锁，保证并发安全
}


(2) 核心特性
线程安全： channel 的操作是原子的，无需额外加锁。
阻塞与非阻塞：
无缓冲 channel：发送和接收必须同时准备好，否则会阻塞。
有缓冲 channel：缓冲区未满时发送不阻塞，未空时接收不阻塞。
方向性：可以声明为只读（ <-chan ）或只写（ chan<- ）。


var 变量 chan 类型
如：
var ch1 chan int   整型管道
var ch2 chan bool  布尔类型管道
var ch3 chan []int 切片类型的管道

// make创建管道
*/
func ChannelTutorial() {

	// 创建
	ch := make(chan int, 3) // 3，表示管道的容量
	// 存储数据到管道
	ch <- 10
	ch <- 30
	ch <- 50
	// 获取管道里面的数据
	<-ch // 取一次值
	a := <-ch
	fmt.Println(a) // 30. 管道里面的数据遵循FIFO

	fmt.Printf("值：%v 容量:%v 长度：%v\n", ch, cap(ch), len(ch))

	// 管道是引用类型
	ch1 := make(chan int, 4)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3

	ch2 := ch1
	ch2 <- 4

	// 管道阻塞
	ch3 := make(chan int, 1)
	ch3 <- 1
	// ch3 <-2 // error: all goroutines are asleep - deadlock! 如果管道容量是1，往里面放多余1的数，将会阻塞

	a1 := <-ch3
	// a2 := <-ch3 // error: all goroutines are asleep - deadlock! 同样，如果管道没有数据，还去取，还是会报错
	fmt.Println(a1)
}

// 循环遍历管道
func ChannelTutorial1() {
	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // 关闭管道。不然遍历会报错fatal error: all goroutines are asleep - deadlock!

	// 遍历。在没有携程的情况下，如果管道中没有数据，还继续读取就会报错
	for v := range ch {
		fmt.Println(v)
	}

	// 通过for 循环不需要close
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

// / goroutine + chan
var ws sync.WaitGroup
func read(ch chan int) {
	for v := range ch {
		fmt.Println("读取数据：", v)
		time.Sleep(time.Second)
	}
	ws.Done()
}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("写入数据：", i)
		time.Sleep(time.Second)
	}
	close(ch)
	ws.Done()
}

func ChannelTutorial2() {
	ch := make(chan int, 10)
	ws.Add(2)
	go read(ch)
	go write(ch)
	ws.Wait()
}

// 单向管道
func ChannelTutorial3() {

	// 单向管道的定义：chan<- int只读。 <-chan int只写。
	ch1 := make(chan<- int, 2)
	ch1 <- 1
	ch1 <- 2

	// <-ch1 Receiving from a send-only channel is not allowed and has been removed.

	ch2 := make(<-chan int, 2)
	// ch2 <- 1 Receiving from a send-only channel is not allowed and has been removed.
	<-ch2
}
