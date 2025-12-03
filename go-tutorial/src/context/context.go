package contextpackage

import (
	"context"
	"fmt"
	"time"
)

/*
一、Context 是什么？
Context（上下文） 是 Go 语言中用于管理请求生命周期的标准库，主要解决并发编程中的三个核心问题：
取消信号传播 - 在 goroutine 树中传播取消信号
超时控制 - 设置操作的截止时间，防止资源无限占用
值传递 - 在调用链中安全地传递请求范围的数据

二、Context 接口详解
type Context interface {
    // 返回截止时间，如果未设置返回 false
    Deadline() (deadline time.Time, ok bool)

    // 返回一个通道，当 Context 被取消时关闭
    Done() <-chan struct{}

    // 返回取消的原因
    Err() error

    // 根据键返回值
    Value(key interface{}) interface{}
}


2.1 根 Context
// 创建空 Context（永不取消，无截止时间）
ctx := context.Background()  // 通常作为根 Context
ctx := context.TODO()        // 不确定用途时的占位符

2.2 派生 Context 的 5 种方式
// 1. 带取消功能的 Context
ctx, cancel := context.WithCancel(parentCtx)
defer cancel()  // 确保调用取消函数释放资源

// 2. 带截止时间的 Context
ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(2*time.Hour))
defer cancel()

// 3. 带超时的 Context
ctx, cancel := context.WithTimeout(parentCtx, 30*time.Second)
defer cancel()

// 4. 带值的 Context
ctx := context.WithValue(parentCtx, key, value)
*/
// 2.2 接口方法详解
// 1. Deadline()
func ExampleDeadline() {
	// 无截止时间
	ctx := context.Background()
	deadline, ok := ctx.Deadline()
	fmt.Printf("截止时间：%v, 是否设置：%v\n", deadline, ok)

	// 有节制时间
	ctxWithDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(1*time.Second))
	deadline1, ok1 := ctxWithDeadline.Deadline()
	fmt.Printf("截止时间：%v, 是否设置：%v\n", deadline1, ok1)
	fmt.Printf("剩余时间：%v\n", time.Until(deadline))

	cancel()
}

// 2. Done()
func ExampleDone() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("goroutine 中的 ctx:", ctx)
        // select 是阻塞的，如果有 default 分支，则会执行 default 分支
		select {
		case <-ctx.Done():
			fmt.Println("Context 已取消")
		case <-time.After(3 * time.Second):
			fmt.Println("操作完成")
		}
        fmt.Println("goroutine 结束")
	}()

	// 1s 后取消 Context
	time.Sleep(1 * time.Second)
	cancel()
	fmt.Println("主程序结束")

    // 这里必须要执行，否则 goroutine 可能不会有机会执行。
    // 需要理解 goroutine 的调度机制。
	time.Sleep(500 * time.Millisecond) // 等待 goroutine 处理完成
}

// 3. Err()
func ExampleErr() {
    ctx, cancel := context.WithCancel(context.Background())
    cancel()

    fmt.Println("取消后的错误:", ctx.Err()) // 输出: 取消后的错误: context canceled

    // 超时错误示例
    ctx1, _ := context.WithTimeout(context.Background(), time.Second)
    time.Sleep(2 * time.Second)
    fmt.Println("超时后的错误:", ctx1.Err()) // 输出: 超时后的错误: context deadline exceeded
}

// 4. Value()
func ExampleValue() {
    type contextKey string 
    const userIDKey contextKey = "userID"

    ctx := context.WithValue(context.Background(), userIDKey, 42)

    // 获取值
    if userId, ok := ctx.Value(userIDKey).(int); ok {
        fmt.Println("User ID:", userId) // 输出: User ID: 42
    } else {
        fmt.Println("User ID not found")
    }

    // 尝试获取不存在的键
    if role, ok := ctx.Value("role").(string); ok {
        fmt.Println("Role:", role)
    } else {
        fmt.Println("Role not found") // 输出: Role not found
    }
}
