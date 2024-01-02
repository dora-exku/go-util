### Go-Util golang 工具包
### 1. 介绍
Go-Util 是一个golang工具包，包含了一些常用的工具类，方便开发者使用。
### 2. 安装
```shell
go get github.com/chenyingdi/go-util
```
### 使用
#### password
```go
package main

import (
    "fmt"
    "github.com/chenyingdi/go-util/password"
)

func main() {
    // 生成密码 
    password1 := "secret"
    hashedPassword := password.MakePassword(password1)
    fmt.Println(hashedPassword)
    // 验证密码
    fmt.Println(password.CheckPassword(hashedPassword, password1))
}
```