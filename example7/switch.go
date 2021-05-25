package main

import (
	"fmt"
	"time"
)

func main()  {
	
	// 普通用法
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("three")
	}

	t := time.Now().Weekday()
	fmt.Println(t)
	switch t {
	case time.Saturday, time.Sunday : 
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}
}