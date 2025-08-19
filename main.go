package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // 默认端口，方便本地调试
	}

	http.HandleFunc("/", HomeHandler)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // 设置响应类型为 HTML

	// 输出 HTML 内容
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>欢迎访问我的网站</title>
			<style>
				body { font-family: Arial, sans-serif; background: #f4f4f4; text-align: center; padding-top: 100px; }
				h1 { color: #333; }
			</style>
		</head>
		<body>
			<h1>Hello from Koyeb 👋</h1>
			<p>这是一个简单的网页，由 Go 编写并部署。</p>
		</body>
		</html>
	`)
}
