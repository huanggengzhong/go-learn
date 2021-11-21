### 1.官网下载golang。
[下载链接](https://golang.google.cn/)

我是1.14.6版本，大家根据自己需要选择版本
### 2.安装插件和配置setting.json，右下角提示要安装其他一律允许。

![](https://files.mdnice.com/user/16210/4cc90bd2-a809-440c-b8bc-004d2d6bf7d6.png)

![](https://files.mdnice.com/user/16210/13d81bfc-db08-43f1-b445-573bcdf8e9ce.png)
setting.json
 ```js
    "go.goroot": "C:/Go",
    "go.gopath": "C:/Users/lish/go"
```

![](https://files.mdnice.com/user/16210/12306b1c-0b52-466b-a302-3aa8b9f21cbd.png)

### 3.运行测试
```go

package main 
import "fmt"
func main()  {
	fmt.Println("hello world")
}
```
![](https://files.mdnice.com/user/16210/922ee854-f5c3-47a8-91a6-3de99ac806c6.png)




