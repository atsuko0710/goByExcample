package main

import "fmt"

func main() {
	i := 1
	fmt.Println(i)

	zeroval(i)
	fmt.Println(i)
	

	zeroptr(&i)
	fmt.Println(i)
	fmt.Println(&i)
}

// 传值
func zeroval(ival int) {
	ival = 0
}

// 传指针
func zeroptr(iptr *int) {
	*iptr = 0
}
