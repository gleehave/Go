package main

import "fmt"

type human interface {
	speak()
}

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

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		"miss",
		"moneypenny",
	}

	sal := secreteAgent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	saySomething(p1)
	saySomething(sal)
}
