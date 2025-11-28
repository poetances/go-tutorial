package channel

import (
	"context"
	"fmt"
	"time"
)

// 1.Channel 的基本概念
// 1.1 什么是 channel， Channel 是 Go 语言中用于在不同 Goroutine之间进行通信和同步的核心机制。它基于 CSP 模型（Communicating Sequential Processes）实现。
// 强调通过通信共享内存，而不是通过共享内存来通信。
func ChannelBasic() {
	// Channel声明
	var ch chan int
	var ch1 chan string
	var ch2 chan []int
	fmt.Println(ch, ch1, ch2) // <nil> <nil> <nil>, channel 的零值是 nil，引用类型

	ch = make(chan int)
	ch1 = make(chan string)
	ch2 = make(chan []int)
	fmt.Println(ch, ch1, ch2) // 0x14000096310 0x14000096380 0x140000963f0 引用类型，打印的地址
}

// 1.2 Channel的底层机构
/*
Channel 的底层是一个环形队列（hchan 结构体），包含以下关键字段：
// 简化的 hchan 结构
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
*/

// 2. Channel 的基本操作
// 2.1 发送和接收
func DemonstrateBasicOperations() {
	ch := make(chan int, 3)
	// 发送数据
	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40 这里会阻塞，因为通道已满，这就是有缓冲 channel的特点

	// 接收数据
	value := <-ch
	fmt.Printf("接收到的值: %d\n", value) // 输出: 10

	// 接收并忽略值
	<-ch // 如果 channel 里面没值，这里会阻塞，直到有值

	// 检查 channel 的状态
	fmt.Printf("channel: %v, 容量：%d, 长度：%d\n", ch, cap(ch), len(ch))
	//输出：channel: 0x1400001e280, 容量：3, 长度：1。需要注意长度和容量之间的区别。
}

// 2.2 无缓存 Channel
func DemonstrateUnBufferChannel() {
	// 创建无缓存 channel
	ch := make(chan int)

	// 启动一个 Goroutine 来接收数据。
	go func() {
		value := <-ch
		fmt.Printf("接收到的值: %d\n", value)
	}()

	// 发送数据（会阻塞，知道有接收者）
	fmt.Println("准备发送数据。。。")
	ch <- 10
	// ch <- 20 这里会阻塞，因为接收的地方就一次，所以第二次发送会阻塞
	fmt.Println("数据发送成功。。。")

	// 上面代码有一些细节：
	// 1.一定要在数据发送前有接收数据的，否则会阻塞。
	// 2.一定要使用 go 来开启接收，因为直接进行接收，会阻塞当前 Goroutine。因为channel 里面没有数据，所以会阻塞。
	// 3.接收、发送一定是成对出现的，否则会阻塞。

	// 总结：无缓存 channel 的特点是发送和接收必须同时准备好，否则会阻塞。
	// 有缓冲 channel 的特点是发送和接收可以不同时准备好，发送时缓冲区满了，接收时缓冲区空了，都会阻塞。
	// 也就是：channel 在接收的时候会阻塞，直到有数据发送；在发送的时候会阻塞，直到有数据接收。
}

// 2.3 有缓存 Channel
func DemonstrateBufferChannel() {
	// 创建有缓存 channel
	ch := make(chan int, 3)

	// 发送数据（缓存区不满，就会不阻塞）
	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40 这里会阻塞，因为缓存区已满
	fmt.Printf("Channel 长度: %d, 容量: %d\n", len(ch), cap(ch))

	// 接收数据
	value := <-ch
	value1 := <-ch
	value2 := <-ch
	// value3 := <-ch 会阻塞，因为channel里面没有数据

	fmt.Printf("接收到的值: %d %d %d\n", value, value1, value2)
	fmt.Printf("接收后 Channel 长度: %d\n", len(ch))
}

// 3. channel 的方向性
// 3.1 双向 channel。默认情况下，channel 是双向的，可以用于发送和接收数据。
func DemonstrateBidirectionalChannel() {
	ch := make(chan int, 1)
	// 发送数据
	ch <- 10
	// 接收数据
	value := <-ch
	fmt.Printf("接收到的值: %d\n", value)
}

// 3.2 单向 channel。可以将 channel 声明为只发送或只接收的类型，以限制对 channel 的操作。
func DemonstrateUnidirectionalChannel() {
	// 创建双向 channel
	ch := make(chan int, 1)

	// 启动接收者
	go sender(ch)
	receiver(ch)
}

func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("发送数据：", i)
	}
	close(ch)
}
func receiver(ch <-chan int) {
	for value := range ch {
		fmt.Println("接收数据：", value)
	}
}

// 4. Channel的关闭
// 4.1 关闭 channel 的语法：close(ch), 关闭后，channel 不能再发送数据，但是可以继续接收数据。
func DemonstrateChannelClosing() {
	ch := make(chan int, 3)

	// 发送数据
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch) // 一定要 close 掉，否则会死锁，接受完成 永远不会执行
	// ch <- 40 // 这行会引发panic，因为channel已关闭
	for value := range ch {
		fmt.Println("接收到的值:", value)
	}
	fmt.Println("接收完成")

	value, ok := <-ch
	if !ok {
		fmt.Println("channel 已关闭")
	}
	fmt.Printf("接收到值: %d, 是否关闭: %t\n", value, !ok)
}

// 4.2 多返回值接收
// 使用多返回值语法可以同时获取channel的值和关闭状态
func DemonstrateMultiValueReceive() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 20
	ch <- 30
	close(ch)

	// 使用多返回值接收。这里必须有值才能避免死循环（这也是 channel 的底层逻辑）。
	// 或者用 goroutine 来接收（其实相当于阻塞了 Goroutine）
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("channel 已关闭")
			break
		}
		fmt.Printf("接收：%d\n", value)
	}
}

// 5. Select 语句
// 5.1 基本 select 语句
func DemonstrateBasicSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// 启动 Goroutine 发送数据
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 ch1 的消息"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自 ch2 的消息"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("接收到 ch1 的消息:", msg1)
		case msg2 := <-ch2:
			fmt.Println("接收到 ch2 的消息:", msg2)
		}
	}
}

// 5.2 Select 与默认分支
func DemonstrateSelectWithDefault() {
	ch := make(chan int, 1)
	ch <- 5

	// 非阻塞发送
	select {
	case ch <- 10:
		fmt.Println("发送成功")
	default:
		// 这个时候会进这个分支，因为channel里面有值
		fmt.Println("发送失败， channel已满")
	}

	ch1 := make(chan int, 1)
	select {
	case value := <-ch1:
		fmt.Printf("接收到值: %d\n", value)
	default:
		// 这个时候会进这个分支，因为channel里面没有值
		fmt.Println("接收失败，channel 为空")
	}
}

// 6.Channel使用场景
// 6.1 数据传递
// Channel 最基本的用法就是在 Goroutine 之间传递数据。
func DemonstrateDataTransfer() {
	dataCh := make(chan []int)

	// 启动生产者
	go func() {
		data := []int{1, 2, 3, 4, 5}
		fmt.Printf("生产者发送数据: %v\n", data)
		dataCh <- data
	}()

	// 启动消费者
	go func() {
		time.Sleep(1 * time.Second)
		data := <-dataCh
		fmt.Printf("消费者收到数据: %v\n", data)

		// 处理数据
		sum := 0
		for _, v := range data {
			sum += v
		}
		fmt.Printf("数据求和结果: %d\n", sum)
	}()

	// 等待完成
	time.Sleep(1 * time.Second)
}

// 6.2 信号通知
// Channel 可以用于通知Goroutine 停止或继续执行
func DemonstrateSignalNotification() {
	stopCh := make(chan struct{})

	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("接收到停止信号")
				return
			default:
				fmt.Println("工作中...")
				time.Sleep(1 * time.Second)
			}
			fmt.Println("for 循环")
		}
	}()

	// 运行一段时间后发送停止信息
	time.Sleep(5 * time.Second)
	close(stopCh)

	// 等待工作 Goroutine 退出
	time.Sleep(1 * time.Second)
	fmt.Println("主程序退出")
}

// 6.3 多任务分发
// 使用 Channel 实现任务分发模式，将任务分发给多个 Goroutine 执行
func DemonstrateTaskDistribution() {
	// 创建多任务
	taskCh := make(chan int, 10)
	resultCh := make(chan int, 10)

	// 启动多个工作 Goroutinue
	for i := 0; i < 3; i++ {
        go func(workerID int) {
			fmt.Printf("Worker %d 启动\n", workerID)
            for task := range taskCh {
                result := task * task  // 计算平方
                fmt.Printf("Worker %d 处理任务 %d，结果 %d\n", workerID, task, result)
                resultCh <- result
            }
			fmt.Printf("Worker %d 关闭\n", workerID)
        }(i)
    }

	// 发送任务
	time.Sleep(1 * time.Second)
	go func() {
		for i := 0; i <= 10; i++ {
			taskCh <- i
		}
		close(taskCh)
	}()
	
	// 收集结果
	go func() {
		for i := 0; i < 10; i++ {
			result := <-resultCh
			fmt.Printf("收集到结果：%d\n", result)
		}
	}()

	// 等待完成
	time.Sleep(2 * time.Second)
	fmt.Println("已完成")
}

// 8. Channel最佳实践
// 8.1 避免死锁
func deadlockExample() {
	ch := make(chan int)
    
    // 以下代码会死锁，因为没有接收者
    // ch <- 1
    
    // 以下代码也会死锁，因为没有发送者
    // <-ch
	fmt.Println(<-ch)
}
// 正确示例：确保有对应的发送者和接收者
func correctExample() {
    ch := make(chan int)
    
    // 启动 goroutine 作为接收者
    go func() {
        value := <-ch
        fmt.Printf("接收到值: %d\n", value)
    }()
    
    // 发送数据
    ch <- 42
}
func DemonstrateDeadlock() {
	deadlockExample()
	correctExample()
}

// 8.2 使用 Select 处理超时
func DemonstrateTimeoutWithSelect() {
	ch := make(chan int)
	
	// 启动发送
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case value := <-ch:
	fmt.Println("收到值:", value)
	case <-time.After(1 * time.Second):
		// 超过 1s 就会出现超时
		fmt.Println("超时")
	}
}

// 8.3 使用 Context控制 Goroutine
func DemonstrateContextControl() {
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 Goroutine
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("接收到取消信号")
				return
			default:
				fmt.Println("工作中...")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	// 运行一段时间
	time.Sleep(5 * time.Second)
	cancel()

	// 等待 goroutine 完成
	time.Sleep(1 * time.Second)
}

