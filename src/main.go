package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secreteAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `Says, "Good morning, James"`)
}

func (sa secreteAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `Says, "Shaken, not stirred."`)
}

func main() {
	p1 := person{
		"miss",
		"moneypenny",
	}
	p1.speak()

	sal := secreteAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}
	sal.speak()
}
