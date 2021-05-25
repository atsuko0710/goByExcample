package main

import "fmt"

// 切片
func main() {
	s := make([]int, 3)
	s = append(s, 8) // 动态数组
	s = append(s, 9)
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	c := make([]int, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println(l)

	l1 := s[2:]
	fmt.Println(l1)

	l2 := s[:4]
	fmt.Println(l2)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
