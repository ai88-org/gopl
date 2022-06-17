# 《The Go Programming Language》

## 接口

接口是*type*和*interface*关键字定义的一组方法集合。

### 入门案例

```go
#gopl/interface/eg1
package main

import "fmt"

// 创建一个接口类型，有两个方法
type Usb interface {
	Start()
	Stop()
}

// 手机结构体
type Phone struct {
}

// 手机实现了Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机启动了...")
}

func (p Phone) Stop() {
	fmt.Println("手机拔出来了...")
}

// 相机结构体
type Camera struct {
}

// 相机实现了start方法
func (c Camera) Start() {
	fmt.Println("相机开机了...")
}

// 相机实现了stop方法
func (c Camera) Stop() {
	fmt.Println("相机关机了...")
}

// 电脑结构体
type Computer struct {
}

// 电脑的工作方法
// 参数是一个USB接口
func (c Computer) Working(u Usb) {
	u.Start()
	u.Stop()
}

func main() {
	c := Computer{}
	p := Phone{}
	cm := Camera{}

	c.Working(p)
	c.Working(cm)
}

```

- 接口本身不能创建实例，但是可以指向实现了接口的自定义类型变量。实现接口的意思是实现了接口的所有方法。

```go
package main

func main() {
	var a AInterface
	a.SayHello()
}

// 直接报错 panic: runtime error: invalid memory address or nil pointer dereference
type AInterface interface {
	SayHello()
}
```

这样写是正确的：

```go
package main

import "fmt"

func main() {
	var a AInterface
	var m myType
	a = m
	a.SayHello()
}

type AInterface interface {
	SayHello()
}

type myType struct {
}

func (m myType) SayHello() {
	fmt.Println("Hello...")
}
```

- 只要是自定义类型都可以实现接口，不仅仅是结构体。

```go
package main

import "fmt"

func main() {
	var a AInterface
	var m myType
	a = m
	a.SayHello()

	var i integer
	a = i
	a.SayHello()
}

type AInterface interface {
	SayHello()
}

type myType struct {
}

// 自定义的类型也能实现接口
type integer int

func (i integer) SayHello() {
	fmt.Println("this is integer say Hello...")
}

func (m myType) SayHello() {
	fmt.Println("Hello...")
}

```

- 自定义类型可以实现多个接口，实现每个接口的方法就可以了。

```go
// 自定义类型可以实现多个接口
package main

import "fmt"

func main() {
	var a A
	var b B
	var j JustX

	a = j
	b = j
	fmt.Println(a.Say)
	fmt.Println(b.Hello)
}

type A interface {
	Say()
}

type B interface {
	Hello()
}

type JustX struct {
}

func (s JustX) Say() {
	fmt.Println("s say ...")
}

func (s JustX) Hello() {
	fmt.Println("s hello ...")
}

```

- 接口可以继承其他接口，方法需要被实现该接口的类型全部实现。

```go
package main

import "fmt"

func main() {
	var a A
	var x JustX
	a = x
	a.yeah()
	a.Say()
	a.Hello()
}

type B interface {
	Say()
}

type C interface {
	Hello()
}

// A接口继承了B、C接口
type A interface {
	B
	C
	yeah()
}

type JustX struct {
}

func (j JustX) Say() {
	fmt.Println("say...")
}

func (j JustX) Hello() {
	fmt.Println("hello...")
}

func (j JustX) yeah() {
	fmt.Println("yeah...")
}
```

- 空接口可以接受任意的自定义类型，因为其中没有方法，也就不需要实现了。

- 下面这个编译报错，相同名称的函数，返回值类型不一样，则无法被另外接口继承（go在1.14版本之后，go接口类型允许嵌入的不同接口类型的方法存在交集，前提是交集中方法不仅方法名称要一样，方法签名也要一样，也就是参数和返回值也要一样）

```go
// go version go1.18 darwin/amd64
package main

import "fmt"

func main() {
}

type B interface {
	// 返回值类型应该一样
	Test01() int
	Test02()
}

type C interface {
  // 返回值类型应该一样
	Test01() string
	Test03()
}

type A interface {
	B
	C
}

type JustX struct {
}

func (x JustX) Test01() {
	fmt.Println("this is test01")
}

func (x JustX) Test02() {
	fmt.Println("this is test02")
}

func (x JustX) Test03() {
	fmt.Println("this is test03")
}

```

- 必须是严格的自定义类型本身实现了这个方法，哪怕是这个类型的指针类型实现了方法都不行

```go
package main

import "fmt"

func main() {
	var a A
	var j Just
	a = j
	a.Say()
}

type A interface {
	Say()
}

type Just struct {
}

// 是Just的指针类型实现了这个方法
// 并不是Just类型本身实现了这个方法，所以也会报错
func (j *Just) Say() {
	fmt.Println("this is say")
}

// 报错信息
cannot use j (variable of type Just) as type A in assignment:
	Just does not implement A (Say method has pointer receiver)
```

- 如果一个变量的类型是空接口类型，也就意味着方法集合为空，也就是说任何类型都实现了接口，所以可以被任何类型赋值。

```go
package main

import "fmt"

func main() {
	var i interface{} = 10
	i = "helloworld"
	i = true

	type T struct {
	}
	i = T{}
	fmt.Println(i)
}

```

### 类型断言

格式如下:

```go
v,ok := i.(T)
```

i为接口类型变量，T为非接口类型，并且T是想要还原的类型。

下面例子虽然`Dog`类型也实现了接口，但是接口类型变量i并未接受Dog类型变量，所以断言的`ok`就是false

```go
package main

import "fmt"

func main() {
	var c = Cat{}
	var d Dog
	fmt.Println(d)
	var interface2 i = c
	if v, ok := interface2.(Cat); ok {
		fmt.Printf("%v\t%T\n", v, v)
	}

	if v, ok := interface2.(Dog); !ok {
		fmt.Printf("%v\t%T\n", v, v)
	}
}

type i interface {
	Say()
}

type Cat struct {
}

type Dog int

func (s Cat) Say() {
	fmt.Println("cat hello....")
}

func (s Dog) Say() {
	fmt.Println("dog hello....")
}

```

