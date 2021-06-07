package main

import "fmt"

func main()  {
	// 创建一个新的通道
	message := make(chan string)

	// 在一个协程发送一个新的值到通道中
	go func() {
		message <- "ping"
	}()

	// 将上个协程的值传到另一个协程中
	msg := <- message
	fmt.Println(msg)
}