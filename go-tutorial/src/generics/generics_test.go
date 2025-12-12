package generics

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerics(t *testing.T) {
	handler := NewHandler[string]()

	printstring := func(s string) {
		fmt.Printf("处理字符串: %s\n", s)
		time.Sleep(100 * time.Millisecond)
	}
	handler.Start(10, printstring)

	 for i := 0; i < 15; i++ {
        success := handler.Push(fmt.Sprintf("项目-%d", i))
        if !success {
            fmt.Printf("队列已满，无法推送项目-%d\n", i)
        }
        time.Sleep(50 * time.Millisecond)
    }

	// 运行一段时间后停止处理器
    time.Sleep(2 * time.Second)
	
	handler.Stop()
}