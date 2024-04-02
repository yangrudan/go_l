# go_l
初学者学习指南


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

make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 和 channel。

### 6.4 切片append

append 方法总是返回成功，除非系统内存耗尽了

## 七.bytes包

Go语言的bytes包专为处理字节切片（[]byte）而设计，它提供了一系列强大的函数和类型，使得字节序列的操作既简单又高效。

> 此外，bytes包中的函数和类型都是并发安全的，这意味着在多线程环境下，您无需担心数据竞争或同步问题。这一点对于构建高并发的网络服务或应用来说尤为重要。

+ 比较字节切片 - 使用bytes.Compare和bytes.Equal进行字节切片的比较

+ 字节切片的拼接 - 使用bytes.Join将多个字节切片拼接在一起。

+ 查找子串 - 使用bytes.Contains和bytes.Index查找字节切片中的子串
