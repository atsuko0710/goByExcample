package main

import "fmt"

func f(form string) {
	for i := 0; i < 3; i++ {
		fmt.Println(form, ":", i)
	}
}

func main() {
	f("direct")

	// 两个 Go 协程没有交替输出
	go f("goroutine")

	go func(msg string) {
        fmt.Println(msg)
    }("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
