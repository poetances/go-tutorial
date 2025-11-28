package contextpackage

import "context"

/*
一、Context 是什么？
定义：Context 是 Go 语言中用于管理请求生命周期的标准库，主要用于在不同 goroutine 之间传递上下文信息。

✅ 取消信号传播
✅ 超时控制
✅ 值传递
✅ 截止时间管理

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
*/


