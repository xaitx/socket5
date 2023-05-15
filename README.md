# socket5

socket5是一个基于Go语言开发的SOCKS5代理服务器库，可以处理未加密的TCP流量。该库支持多种认证方式，包括无认证、账号密码认证等，同时支持自定义认证接口。该库还支持日志记录和流量监控，用户可以选择使用默认的实现或自定义实现。

(部分功能未实现)

## 安装

使用go get命令获取socket5库：

arduinoCopy code

```
go get github.com/username/repo/socket5
```

## 使用

下面是一个简单的例子：

goCopy code

```
package main

import "socket5"

func main() {
    c := socket5.NewConfig()
    s := socket5.NewServer(*c, nil, nil, nil)
    s.ListenAndServe()
}

```

在上面的代码中，我们使用默认的配置创建一个socket5服务器并开始监听。

如果您想自定义服务器的配置，可以创建一个Config实例并将其传递给NewServer方法：

goCopy code

```
package main

import (
    "socket5"
)

func main() {
    c := socket5.Config{
        Address: "127.0.0.1",
        Port:    8080,
        Auth:    &MyAuth{},
        Logger:  &MyLogger{},
        Monitor: &MyMonitor{},
    }

    s := socket5.NewServer(c, nil, nil, nil)
    s.ListenAndServe()
}

```

在上面的代码中，我们自定义了服务器的监听地址、端口号以及开启了认证、使用自定义的认证器、日志记录器和流量监控器。