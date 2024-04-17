package main

import (
	"fmt"
	"time"
)


func hello() {
	fmt.Println("Hello")
}

func main() {  //开启一个主goroutine执行main函数
	go hello()  //开启一个goroutine执行这个函数
	fmt.Println("end")
	time.Sleep(time.Second)
}