package main

import "fmt"

type rect struct {
	Width int
	Hight int
}

// 结构体方法
func (r rect) area() int {
	return r.Hight * r.Width
}

func main()  {
	r := rect{Width: 2,Hight: 10}
	fmt.Println(r.area())
}