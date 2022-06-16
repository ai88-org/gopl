# 《The Go Programming Language》

## 8 Goroutines and Channels

### 8.1 Goroutine

当一个程序开始的时候，只有一个goroutine，就是main。



使用go关键字调用方法或者函数，该方法或者函数将会在新创建的goroutine中运行，调用语句本身会很快的完成。（这句话对于理解go xx() 的运行非常有帮助！！！）

```go
f() // 调用函数，等其完成
go f() // 开启新的goroutine调用f()，不会等待
```



下面这个例子里面，计算第45个斐波那契数字需要消耗一些时间，就会让小动画存在一点时间，当main goroutine执行完毕之后，程序就结束了。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
```



> clock1

```go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
    // 生成listener对象，这个对象监听8000端口，一直阻塞，直到有新的连接访问这个端口
    // 之后返回这个连接
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

    // 不断的监听这个端口
	for {
        // 阻塞，等待新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
        // 处理这个连接
		handleConn(conn)
	}

}

// 向conn写入时间字符串，一秒钟一次
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
```



> netcat

```go
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
    // 把连接输出到标准输出
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
```



netcat1先连接，netcat2后连接，由于是顺序执行，只要netcat1不断开，则netcat2就一直没有机会去连接。

我们只需要再clock1的handleConn方法加一个`go`关键字即可。

`go handleConn`创建了一个goroutine之后，语句立即结束，不会阻塞在这里。

然后就会连接上一个新的conn，也就是netcat2也会被拿到。

