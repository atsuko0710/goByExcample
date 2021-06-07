package main

import "fmt"

func main() {

	// 遍历数组
	nums := [...]int{1,2,3,4,5}
	for key, num := range(nums) {
		fmt.Println(key, num)
	}

	kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range(kvs) {
		fmt.Println(k, v)
	}

	// range 在字符串中迭代 unicode 编码。第一个返回值是rune 的起始字节位置，然后第二个是 rune 自己。
	for i, c := range "go" {
		fmt.Println(i,c)
	}
}
