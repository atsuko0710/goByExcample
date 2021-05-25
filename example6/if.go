package main

import "fmt"

// if 分支
func main()  {
	if 8 % 2 == 0 {
		fmt.Println(true);
	}

	if 7 % 2 == 0 {
		fmt.Println(true);
	} else {
		fmt.Println(false);
	}

	// if中声明的变量可以用在分支
	if n := 8; n < 5 {
		fmt.Println("n 小于 5");
	} else if n < 10 {
		fmt.Println("n 小于 10");		
	} else {
		fmt.Println("n 大于 10");
	}
		
}