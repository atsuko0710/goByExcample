package main

import "fmt"

// for 循环
func main()  {

	i := 1
	for i < 5 {
		fmt.Println("i=", i)
		i = i + 1
	}

	for j := 4; j > 1; j-- {
		fmt.Println("j=",j)
	}

	for {
		fmt.Println("loop")
		break
	}
}