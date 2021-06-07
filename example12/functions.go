package main

import "fmt"

func main() {
	// fmt.Println(plus(1, 2))

	_, b, _ := vals()
	fmt.Println(b)

	arr := []int{1, 3, 5, 6, 8, 9, 10}
	fmt.Println(sum(arr...))

	intseq := closures()
	fmt.Println(intseq())
	fmt.Println(intseq())

	intseq2 := closures()
	fmt.Println(intseq2())
}

// 函数
func plus(a int, b int) int {
	return a + b
}

// 多返回函数
func vals() (int, int, int) {
	return 3, 5, 7
}

// 多参函数
func sum(nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

// 闭包
func closures() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
