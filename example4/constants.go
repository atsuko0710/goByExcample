package main

import (
	"fmt"
	"math"
)

const s string = "string"

// 常量
func main()  {
	fmt.Println(s)

	const c = 50000
	// const c int = 50000 这样会报错
	// math.Sin 参数需要float类型，不定义类型const会自动转换
	fmt.Println(math.Sin(c))
}