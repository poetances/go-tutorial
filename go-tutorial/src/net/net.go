package net

import (
	"fmt"
	"net/http"
)

func init() {
	println("net package init")
	mux := http.NewServeMux()

	// 注册路由
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/user", userDetailHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 启动服务器的goroutine
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "欢迎来到首页")
}

func userDetailHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "用户详情页面")
 }