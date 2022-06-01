# 《The Go Programming Language》

## 2 程序结构

### 2.5 类型声明

像`==`和`<`这样的比较运算符可以比较一个命名类型的变量和一个相同类型的变量，或者有相同底层类型的字面值。不同命名的类型的变量不能比较，哪怕有相同的底层类型。

```go
// 摄氏温度
type Celsius float64

// 华氏温度
type Fahrenheit float64

func main() {
	// 相同类型的比较
	var c1 Celsius = 100
	var c2 Celsius = 100
	var c3 Celsius = 101
	fmt.Println(c1 == c2)	// true
	fmt.Println(c1 < c3)	// true
	
	// 与有相同底层类型的字面值比较 这里的100,101会被编译器认为是float64类型
	fmt.Println(c1 == 100)	// true
	fmt.Println(c1 < 101)	// true
	
	// 不同类型名称不能比较，哪怕底层类型一样，或者说就是底层类型也没用
	var f1 float64 = 100
	var f11 Fahrenheit = 100
	var f2 float64 = 101
	var f21 Fahrenheit = 101
	fmt.Println(c1 == f1)	// 编译错误 类型不匹配
	fmt.Println(c1 == f11)	// 编译错误 类型不匹配
	fmt.Println(c1 < f2)	// 编译错误 类型不匹配
	fmt.Println(c1 < f21)	// 编译错误 类型不匹配
    
    fmt.Println(c1 == Celsius(f11)) 
}
```

注意最后一个例子，它实际上只是类型转换，并没有更改f11的值，只是改了f11的类型。



一个命名的类型可以提供书写的方便，对于float64这类简单类型可能体现不出来，但是我们后面学习的结构体用这种命名类型的方式会简便很多。



命名的类型也可以为该类型的值定义新的行为，这些新的行为我们称之为类型的方法集，第六章我们会研究，这里稍微了解一下。



举个例子，下面的函数，Celsius类型的参数在函数名称前面，这就是与Celsius类型关联的方法集。

```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
```



