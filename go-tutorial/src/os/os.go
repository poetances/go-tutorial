package os

import (
	"fmt"
	"io"
	"os"
)

/*
我来为您详细讲解 Go 中的 os 包，这是 Go 标准库中非常重要的一个包，用于与操作系统交互。

Go OS 包详解
一、OS 包是什么？
os 包是 Go 提供的与操作系统接口的标准库，它提供了与操作系统交互的基本功能，包括文件操作、进程管理、环境变量访问等。

核心功能：

📁 文件和目录操作
🔧 进程管理
🌍 环境变量操作
📋 系统信息获取
🚀 命令行参数处理
*/

// 文件操作示例- 最常用的功能
func DemonstrateFileCreation() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current working directory:", cwd)

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// 写入内容
	content := "Hello, Go os package!"
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and content written successfully.")
}

// 文件读取示例
func DemonstrateFileReading() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(data))
}

// 文件信息获取
func DemonstrateFileInfo() { 
	// 获取文件信息
	fileInfo, err := os.Stat("example.txt")
	if err != nil {
		fmt.Println("Error getting file info:", err)
		return
	}

	fmt.Printf("File Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Last Modified: %s\n", fileInfo.ModTime())

	// 检查文件是否存在
    if _, err := os.Stat("nonexistent.txt"); os.IsNotExist(err) {
        fmt.Println("文件不存在")
    }
}

// 文件重命名和移动
func DemonstrateFileRenameAndMove() {
    // 重命名文件
    err := os.Rename("example.txt", "renamed.txt")
    if err != nil {
        fmt.Printf("重命名文件失败: %v\n", err)
        return
    }
    
    fmt.Println("文件重命名成功")
    
    // 移动文件（在不同目录间）
    // err = os.Rename("renamed.txt", "backup/renamed.txt")
}

// 文件删除
func DemonstrateFileDeletion() {
    err := os.Remove("renamed.txt")
    if err != nil {
        fmt.Printf("删除文件失败: %v\n", err)
        return
    }
    
    fmt.Println("文件删除成功")
}

// os环境变量
func DemonstrateEnvironmentVariables() {
    // 获取所有环境变量
    envVars := os.Environ()
    fmt.Println("所有环境变量:")
    for _, env := range envVars {
        fmt.Printf("  %s\n", env)
    }
    
    // 获取特定环境变量
    path := os.Getenv("PATH")
    fmt.Printf("\nPATH环境变量: %s\n", path)
    
    // 获取特定环境变量，带默认值
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost" // 默认值
        fmt.Printf("DB_HOST未设置，使用默认值: %s\n", dbHost)
    }
    
    // 设置环境变量
    os.Setenv("APP_ENV", "development")
    appEnv := os.Getenv("APP_ENV")
    fmt.Printf("APP_ENV: %s\n", appEnv)
    
    // 清除环境变量
    os.Unsetenv("APP_ENV")
    appEnv = os.Getenv("APP_ENV")
    fmt.Printf("清除后的APP_ENV: '%s'\n", appEnv)
    
    // 检查环境变量是否存在
    if _, exists := os.LookupEnv("HOME"); exists {
        home := os.Getenv("HOME")
        fmt.Printf("HOME目录: %s\n", home)
    }
}