package main

import "fmt"

type person struct {
	fname string
	lname string
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
	fmt.Printf("%T", p1)
}
