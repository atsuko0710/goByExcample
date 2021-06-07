package main

import "fmt"

// 关联数组
func main()  {
	m := make(map[string]int)

	m["k1"] = 8
	m["k2"] = 9

	fmt.Println(m)

	_, prs := m["k3"]
	fmt.Println(prs)

	n := map[string]int{"foo":1, "bar":2}
	prs1, prs2 := n["bar"]
	fmt.Println(prs1,prs2)
	delete(n, "bar")
	fmt.Println(n)
}