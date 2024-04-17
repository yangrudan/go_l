package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello", i)
	wg.Done()  //通知wg将计数器-1
}

func main() { //开启一个主goroutine执行main函数
	wg.Add(10000)  //计数牌+1
	for i := 0; i < 10000; i++{
		go hello(i) //开启一个goroutine执行这个函数
	}
	fmt.Println("end")
	wg.Wait() //等待派发出去的goroutine完成后收兵
}
