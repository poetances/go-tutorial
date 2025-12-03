package log

import (
	"fmt"
	"log"
	"os"
)

func DemonstrateLog() {

	// 普通日志
	log.Println("应用程序启动")

	// 信息日志
	log.Printf("服务器启动在端口 %d", 8080)

	// 警告日志（Go标准库没有专门的警告级别，使用普通日志表示）
	log.Println("警告: 配置文件未找到，使用默认配置")

	// 错误日志
	log.Printf("错误: 数据库连接失败 - %v", fmt.Errorf("连接超时"))

	// 调试日志
	log.Printf("调试: 处理请求 %s", "/api/users")
}

// 日志级别与格式
func DemonstrateLogFormats() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("应用程序启动")

	// 自定义日志器
	customLogger := log.New(os.Stdout, "CUSTOM: ", log.LstdFlags)
	customLogger.Println("自定义前缀的日志")

	// 不同格式标志
	formats := []int{
		log.Ldate,         // 日期: 2023/12/25
		log.Ltime,         // 时间: 10:20:30
		log.Lmicroseconds, // 微秒: 10:20:30.123456
		log.Llongfile,     // 完整文件路径: /path/to/file.go:123
		log.Lshortfile,    // 文件名和行号: file.go:123
		log.LUTC,          // UTC时间
		log.Lmsgprefix,    // 消息前缀（从Go 1.14开始）
		log.LstdFlags,     // 标准标志 (Ldate | Ltime)
	}

	for i, flag := range formats {
		logger := log.New(os.Stdout, fmt.Sprintf("[%d] ", i), flag)
		logger.Println("测试不同格式标志")
	}
}
