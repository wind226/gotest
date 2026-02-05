package main

import (
    "fmt"
    "os/exec"
)

func main() {
    // 执行简单命令
    cmd := exec.Command("ls", "-la")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(string(output))
}
