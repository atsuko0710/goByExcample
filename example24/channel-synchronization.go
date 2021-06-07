package main

import (
	"fmt"
	"time"
)

func worker(done chan bool)  {
	fmt.Print("working ...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 用来通知该方法结束
	done <- true
}

func main() {
	done := make(chan bool, 1)
	
	go worker(done)

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	<-done
}