package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 结构体
func main() {
	fmt.Println(Person{"dadan", 25})
	fmt.Println(Person{Name: "dadan", Age: 25})

	fmt.Println(&Person{Name: "dadan", Age: 25})

	p := Person{Name: "dadan", Age: 25}
	person := &p
	fmt.Println(person)
	person.Age = 30
	fmt.Println(person)	
	fmt.Println(person.Age)	
}
