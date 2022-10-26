package main

import "fmt"

type person struct {
	fname string
	lname string
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `Says, "Good morning, James"`)
}

func main() {
	xi := []int{2, 4, 6, 8}
	fmt.Println(xi)

	m := map[string]int{
		"Todd": 45,
		"Job":  42,
	}
	fmt.Println(m)

	p1 := person{
		"miss",
		"moneypenny",
	}
	p1.speak()
}
