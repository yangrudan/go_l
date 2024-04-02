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

```bash
str := "Beginning of the string " +
	"second part of the string"
```

Go 中使用 strings 包来完成对字符串的主要操作:

### 3.1 前缀和后缀