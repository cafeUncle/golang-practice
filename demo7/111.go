package main

import "fmt"

type Dict map[string]string

type CC struct {
	a int
}

func (d Dict) set(k, v string) {
	d[k] = v
}

func (c CC) set(k int)  {
	c.a = k
}

func main() {
	d := Dict{"1":"2"}
	d.set("3", "4")
	fmt.Println(d)

	c := CC{1}
	c.set(2)
	fmt.Println(c)
}
