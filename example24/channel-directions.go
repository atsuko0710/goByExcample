package main

import "fmt"

// 指定pings这个通道只能用来接收
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 指定pongs这个通道只能用来发送
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
