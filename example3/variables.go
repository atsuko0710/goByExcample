package main

import "fmt"

// 变量
func main()  {
	var a int
	a = 1
	fmt.Println(a)

	var b = 2
	fmt.Println(b)

	var c, d = 3,"test"
	fmt.Println(c,d)

	e := 5
	fmt.Println(e)

	var f,g string = "aa", "1"
	fmt.Println(f,g)
}