package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) Growth() {
	p.age++
}

func (p *Person) ChangeName(newName string) {
	p.name = newName
}

func main() {
	p := Person{"wangzy", 30}
	p.Growth()
	fmt.Printf("%d\n", p.age)
	p.ChangeName("newName")
	fmt.Printf("%s\n", p.name)

	println()

	(&p).Growth()
	(&p).ChangeName("newNewName")
	fmt.Println(p)
}
