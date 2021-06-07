package main

import (
	"errors"
	"fmt"
)

type inner struct {
}

type outer struct {
	in inner
}

func (o *outer) print() {
	println(&o)
	println(&(o.in))
}

func main() {
	o := outer{}
	in := inner{}

	fmt.Println("o 's address:", &o)
	fmt.Println("in 's address", &in)

	//o.in = inner{}
	//
	//println(&o)
	//println(&(o.in))
	//
	//o.printSlice()

	//defer recover() := func {
	//
	//}

	err := errors.New("ude")
	panic(err)
}
