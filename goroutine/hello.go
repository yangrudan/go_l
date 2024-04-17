package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("Hello")
	wg.Done()  //通知wg将计数器-1
}

func main() { //开启一个主goroutine执行main函数
	wg.Add(1)  //计数牌+1
	go hello() //开启一个goroutine执行这个函数
	fmt.Println("end")
	wg.Wait() //等待派发出去的goroutine完成后收兵
}
