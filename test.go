package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    // 发送 GET 请求
    resp, err := http.Get("https://api.github.com")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    
    // 读取响应
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    
    fmt.Println("Status:", resp.Status)
    fmt.Println("Body:", string(body))
}
