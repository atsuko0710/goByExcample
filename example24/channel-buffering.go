package main

import "fmt"

func main() {
	// 设置通道最多缓存3个值，超过了会报错
	message := make(chan string, 3)

	message <- "string1"
	message <- "string2"
	message <- "string3"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
