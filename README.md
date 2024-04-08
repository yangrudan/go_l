# go_l
初学者学习指南

> npm install -g markdown-toc
>
> markdown-toc README.md --bullets='-' --maxdepth=2


- [go_l](#go_l)
  - [一. go mod](#%E4%B8%80-go-mod)
    - [1.1 历史发展](#11-%E5%8E%86%E5%8F%B2%E5%8F%91%E5%B1%95)
    - [1.2 使用mod](#12-%E4%BD%BF%E7%94%A8mod)
  - [二.变量](#%E4%BA%8C%E5%8F%98%E9%87%8F)
  - [三.字符串](#%E4%B8%89%E5%AD%97%E7%AC%A6%E4%B8%B2)
    - [3.1 strings](#31-strings)
    - [3.2 strconv](#32-strconv)
  - [四.日期和时间](#%E5%9B%9B%E6%97%A5%E6%9C%9F%E5%92%8C%E6%97%B6%E9%97%B4)
  - [五.指针](#%E4%BA%94%E6%8C%87%E9%92%88)
  - [六.数组与切片](#%E5%85%AD%E6%95%B0%E7%BB%84%E4%B8%8E%E5%88%87%E7%89%87)
    - [6.1 声明](#61-%E5%A3%B0%E6%98%8E)
    - [6.2 将数组传递给函数](#62-%E5%B0%86%E6%95%B0%E7%BB%84%E4%BC%A0%E9%80%92%E7%BB%99%E5%87%BD%E6%95%B0)
    - [6.3 切片](#63-%E5%88%87%E7%89%87)
    - [6.4 切片append](#64-%E5%88%87%E7%89%87append)
  - [七.bytes包](#%E4%B8%83bytes%E5%8C%85)
    - [7.1 bytes.Buffer的使用](#71-bytesbuffer%E7%9A%84%E4%BD%BF%E7%94%A8)
    - [7.2 bytes.Reader的应用](#72-bytesreader%E7%9A%84%E5%BA%94%E7%94%A8)
    - [7.3 应用](#73-%E5%BA%94%E7%94%A8)
  - [八.map](#%E5%85%ABmap)
  - [九. 双向链表](#%E4%B9%9D-%E5%8F%8C%E5%90%91%E9%93%BE%E8%A1%A8)


## 一. go mod

### 1.1 历史发展

GOPATH -> govendor -> go mod

> 在 Go 1.5 版本之前，所有的依赖包都是存放在 GOPATH 下，没有版本控制。

> Go 1.5 版本推出了 vendor 机制，所谓 vendor 机制，就是每个项目的根目录下可以有一个 vendor 目录，里面存放了该项目的依赖的 package.

> Go 1.11 版本推出 modules 机制，简称 mod，更加易于管理项目中所需要的模块。

### 1.2 使用mod

1. 设置 GO111MODULE

Go Modules 在 Go 1.11 及 Go 1.12 中有三个模式:

+ 默认模式（未设置该环境变量或 GO111MODULE=auto）

+ GOPATH 模式（GO111MODULE=off）

+ Go Modules 模式（ GO111MODULE=on）

2. 开启

```bash
# 临时开启 Go modules 功能
export GO111MODULE=on
# 永久开启 Go modules 功能
go env -w GO111MODULE=on

# 设置 Go 的国内代理，方便下载第三方包
go env -w GOPROXY=https://goproxy.cn,direct

# 查看
go env
```
3. 初始化项目

```bash
go mod init <project_name>

go get github.com/gin-gonic/gin

go mod tidy  # 去除不需要的模块， 清理.mod和.sum
```

go.mod 文件只存在于模块的根目录中。模块子目录的代码包的导入路径等于模块根目录的导入路径（就是前面说的 module path）加上子目录的相对路径。

如果创建了一个子目录叫 common，我们不需要（也不会想要）在子目录里面再运行一次 go mod init 了，这个代码包会被认为就是 apiserver 模块的一部分，而这个代码包的导入路径就是 apiserver/common

## 二.变量

声明变量的一般形式是使用 var 关键字：var identifier type

```go
var a, b *int  //将它们都声明为指针类型

var a int
var b bool
var str string


//因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
	str string
)

```

变量的命名规则遵循骆驼命名法

如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母也大写

**简短形式，使用 := 赋值操作符**

```go
//一般赋值
var identifier [type] = value
var a int = 15
var i = 5
var b bool = false
var str string = "Go says hello to the world!"

//函数体内声明局部变量
a := 1
```

## 三.字符串

因 go 语言行尾行尾自动补全 ; 分号， 字符串拼接如下

```go
str := "Beginning of the string " +
	"second part of the string"
```

Go 中使用 strings 包来完成对字符串的主要操作:

### 3.1 strings

```go
//前缀和后缀
strings.HasPrefix(s, prefix string) bool
strings.HasSuffix(s, suffix string) bool

//字符串包含关系
strings.Contains(s, substr string) bool

//返回字符串 str 在字符串 s 中的索引（str 的第一个字符的索引
strings.Index(s, str string) int
//最后出现位置的索引
strings.LastIndex(s, str string) int


//Replace 用于将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串
strings.Replace(str, old, new, n) string

//统计字符串出现次数
strings.Count(s, str string) int

//Repeat 用于重复 count 次字符串 s 并返回一个新的字符串, 类似python * 3
strings.Repeat(s, count int) string

//修改字符串大小写
strings.ToLower(s) string
strings.ToUpper(s) string

//修剪字符串
strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号
strings.Trim(s, "cut") 来将开头和结尾的 cut 去除掉

//分割字符串
strings.Fields(s) 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块
strings.Split(s, sep) 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

//拼接 slice 到字符串
strings.Join(sl []string, sep string) string

//从字符串中读取内容
strings.NewReader(str) 用于生成一个 Reader 并读取字符串中的内容，然后返回指向该 Reader 的指针
Read() 从 []byte 中读取内容。
ReadByte() 和 ReadRune() 从字符串中读取下一个 byte 或者 rune
```

### 3.2 strconv 
与字符串相关的类型转换都是通过 strconv 包实现的。

```go
strconv.Itoa(i int) string 返回数字 i 所表示的字符串类型的十进制数

strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string 将 64 位浮点型的数字转换为字符串，其中 fmt 表示格式（其值可以是 'b'、'e'、'f' 或 'g'），prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64

strconv.Atoi(s string) (i int, err error) 将字符串转换为 int 型。

strconv.ParseFloat(s string, bitSize int) (f float64, err error) 将字符串转换为 float64 型
```

## 四.日期和时间

```go
package main
import (
	"fmt"
	"time"
)

var week time.Duration
func main() {
	t := time.Now()
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011
	t = time.Now().UTC()
	fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011
	// calculating times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	// The time must be 2006-01-02 15:04:05
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}
```

## 五.指针

当一个指针被定义后没有分配到任何变量时，它的值为 nil。

一个指针变量通常缩写为 ptr

对于任何一个变量 var， 如下表达式都是正确的：var == *(&var)

## 六.数组与切片

### 6.1 声明

```go
var identifier [len]type
var arr1 [5]int

a := [...]string{"a", "b", "c", "d"}
for i := range a {
	fmt.Println("Array item", i, "is", a[i])
}

//数组常量
var arrAge = [5]int{18, 20, 15, 22, 16}
```
### 6.2 将数组传递给函数

+ 传递数组的指针
+ 使用数组的切片

**还可以这样返回！**

```go
package main
import "fmt"

func main() {
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
```

### 6.3 切片
切片（slice）是对数组一个连续片段的引用（该数组我们称之为相关数组，通常是匿名的），所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型）。这个片段可以是整个数组，或者是由起始和终止索引标识的一些项的子集。需要注意的是，终止索引标识的项不包括在切片内。切片提供了一个相关数组的动态窗口。

和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个**长度可变的数组**。

优点：*因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用*

声明切片的格式是： 
```go
var identifier []type（不需要说明长度）。
```

一个切片在未初始化之前默认为 nil，长度为 0。

切片的初始化格式是：
```go
var slice1 []type = arr1[start:end]
```

用make()创建切片, make 的使用方式是：func make([]T, len, cap)，其中 cap 是可选参数。

```go
var slice1 []type = make([]type, len)

//或者简写
slice1 := make([]type, len)  //make 接受 2 个参数：元素的类型以及切片的元素个数。


//下面两种方法可以生成相同的切片
make([]int, 50, 100)
new([100]int)[0:50]  //生成数组再切片
```

new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体;

make(T) 返回一个类型为 T 的初始值，它*只适用于3种内建的引用类型：切片、map 和 channel*。

### 6.4 切片append

append 方法总是返回成功，除非系统内存耗尽了

## 七.bytes包

Go语言的bytes包专为处理字节切片（[]byte）而设计，它提供了一系列强大的函数和类型，使得字节序列的操作既简单又高效。

> 此外，bytes包中的函数和类型都是并发安全的，这意味着在多线程环境下，您无需担心数据竞争或同步问题。这一点对于构建高并发的网络服务或应用来说尤为重要。

+ 比较字节切片 - 使用bytes.Compare和bytes.Equal进行字节切片的比较。

+ 字节切片的拼接 - 使用bytes.Join将多个字节切片拼接在一起。

+ 查找子串 - 使用bytes.Contains和bytes.Index查找字节切片中的子串。

+ 替换子串 - 使用bytes.Replace替换字节切片中的子串。


### 7.1 bytes.Buffer的使用

bytes.Buffer是一个用于字节切片的缓冲区，它提供了多种便捷的方法来处理和构建字节和字符串数据。

>Buffer可以被用来创建一个可变大小的字节缓冲区。通过bytes.NewBuffer或bytes.NewBufferString可以方便地创建一个新的Buffer实例。

```go
buf := bytes.NewBuffer([]byte("Go"))
fmt.Println(buf.String()) // "Go"
```

>读写数据

1. 使用Write、WriteByte和WriteString方法向Buffer中写入数据；
2. 使用Read、ReadByte和ReadBytes等方法从Buffer中读取数据。

```go
buf.WriteString("Lang")
fmt.Println(buf.String()) // "GoLang"

buf := bytes.NewBuffer([]byte("Go"))
b := make([]byte, 2)
n, _ := buf.Read(b)
fmt.Println(string(b[:n])) // "Go"
```

> 作为I/O Reader或Writer使用 - Buffer可以作为I/O接口的实现，用于文件操作和网络通信等场景。

```go
var b bytes.Buffer
b.Write([]byte("Hello "))
io.Copy(&b, strings.NewReader("World!"))
fmt.Println(b.String()) // "Hello World!"
```

### 7.2 bytes.Reader的应用

bytes.Reader是bytes包提供的一个类型，用于从字节切片中读取数据。它实现了io.Reader、io.Seeker和io.ReaderAt接口，使得从字节切片中读取数据变得既方便又高效。

> bytes.Reader可用于创建一个可读取指定字节切片的Reader。通过bytes.NewReader函数，您可以轻松地将字节切片转换为Reader。

```go
data := []byte("Go语言")
reader := bytes.NewReader(data)

buf := make([]byte, 4)
n, err := reader.Read(buf)
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(buf[:n])) // "Go语"

```

> 读取数据 - 使用Read方法从Reader中读取数据

### 7.3 应用

使用bytes包处理日志文件

假设您需要处理一个日志文件，该文件包含多行文本数据。您可以使用bytes包中的函数高效地读取和处理这些数据。

```go

func processLogData(logData []byte) {
    lines := bytes.Split(logData, []byte("\n"))
    for _, line := range lines {
        if bytes.Contains(line, []byte("ERROR")) {
            fmt.Println("Error found:", string(line))
        }
    }
}
```

构建动态字符串内容

在一个Web应用中，您可能需要动态构建HTML或JSON响应。bytes.Buffer在这方面非常有用，它可以帮助您高效地构建字符串。
```go
var buffer bytes.Buffer
for i := 0; i < 3; i++ {
    buffer.WriteString(fmt.Sprintf("<p>Paragraph %d</p>", i+1))
}
fmt.Println(buffer.String())
```

解析二进制数据

当您需要从二进制数据中读取特定格式的信息时，bytes.Reader非常有用。例如，读取一个二进制文件的特定部分。

```go

func readBinaryData(data []byte, offset int64, length int) ([]byte, error) {
    reader := bytes.NewReader(data)
    _, err := reader.Seek(offset, io.SeekStart)
    if err != nil {
        return nil, err
    }

    buf := make([]byte, length)
    _, err = reader.Read(buf)
    if err != nil {
        return nil, err
    }

    return buf, nil
}
```

## 八.map

（[keytype] 和 valuetype 之间允许有空格，但是 gofmt 移除了空格）

在声明的时候不需要知道 map 的长度，map 是可以动态增长的。

未初始化的 map 的值是 nil。

key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float。所以数组、切片和结构体不能作为 key (译者注：含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的），但是指针和接口类型可以。如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

```go
var map1 map[keytype]valuetype
var map1 map[string]int
```

值为任意类型

```go
package main
import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
}
```

## 九. 双向链表

Go语言list容器定义在"container/list"包中，实现了一个双向链表。本文第一部分总结源码包中的方法，第二部分展示使用list包的常见示例用法以及刷题时的用法。

|方法  |类型|  作用|
|-----|-----|-----|
|New()	|list包函数	|创建一个list
|Next() *Element|	Element	|获取当前结点的下一个结点
|Prev() *Element|	Element	|获取当前结点的上一个结点
|Len() int|	List|	获取链表长度
|Front() *Element|	List|	获取链表第一个结点
|Back() *Element|	List|	获取链表最后一个结点


## 十. 接口

```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```
（按照约定，只包含一个方法的）接口的名字由方法名加 [e]r 后缀组成，例如 Printer、Reader、Writer、Logger、Converter 等等。还有一些不常用的方式（当后缀 er 不合适时），比如 Recoverable，此时接口名以 able 结尾，或者以 I 开头（像 .NET 或 Java 中那样）。

实现接口的getValue() 就可以调用showValue

```go
package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

/* method to determine the value of a car */
func (c car) getValue() float32 {
	return c.price
}

/* contract that defines different things that have value */
type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)
}
```

### 10.1 接口的嵌套

一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

比如接口 File 包含了 ReadWrite 和 Lock 的所有方法，它还额外有一个 Close() 方法。


```go 
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}
type Lock interface {
    Lock()
    Unlock()
}
type File interface {
    ReadWrite
    Lock
    Close()
}

```

### 10.2 类型断言

程序中定义了一个新类型 Circle，它也实现了 Shaper 接口。 if t, ok := areaIntf.(*Square); ok 测试 areaIntf 里是否有一个包含 *Square 类型的变量，结果是确定的；然后我们测试它是否包含一个 *Circle 类型的变量，结果是否定的。

```go
package main
import (
	"fmt"
	"math"
)
type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}
func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5
	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}
```

### 10.3 类型type-switch

接口变量的类型也可以使用一种特殊形式的 switch 来检测：type-switch

```go
switch t := areaIntf.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}
```

## 十一. 协程和通道

一个并发程序可以在一个处理器或者内核上使用多个线程来执行任务，但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。

协程是轻量的，比线程更轻。它们痕迹非常不明显（使用少量的内存和资源）：使用 4K 的栈内存就可以在堆中创建它们。因为创建非常廉价，必要的时候可以轻松创建并运行大量的协程（在同一个地址空间中 100,000 个连续的协程）。并且它们对栈进行了分割，从而动态的增加（或缩减）内存的使用；栈的管理是自动的，但不是由垃圾回收器管理的，而是在协程退出后自动释放。

Go有一种特殊的类型，通道（channel），就像一个可以用于发送类型化数据的管道，由其负责协程之间的通信，从而避开所有由共享内存导致的陷阱；这种通过通道进行通信的方式保证了同步性。

> 通道只能传输一种类型的数据，比如 chan int 或者 chan string，所有的类型都可以用于通道，空接口 interface{} 也可以。甚至可以（有时> 非常有用）创建通道的通道。

```go
var ch1 chan string
ch1 = make(chan string)


//或者下面这样
ch1 := make(chan string)
```

### 11.1 通信操作符 <-

**流向通道（发送）**

ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

**从通道流出（接收），三种方式：**

int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；

假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch。

<- ch 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：if <- ch != 1000


### 11.2 通道阻塞





