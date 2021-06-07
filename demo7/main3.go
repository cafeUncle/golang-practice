package main

import "fmt"

type Shape interface {
	getSize() int
}

type Rectangle struct {
	rSize int
}

type Circle struct {
	cSize int
}

func (r *Rectangle) getSize() int {
	return r.rSize
}

func (c *Circle) getSize() int  {
	return c.cSize
}

func main()  {
	areaTests := []struct {
		shape Shape
		want  int64
	}{
		{&Rectangle{12}, 12},
		{&Circle{10}, 314},
	}

	for index, item := range areaTests {
		fmt.Printf("%d %b\n", index, item)
	}


	rec := Rectangle{}

	var rec2 Rectangle

	fmt.Println(rec)
	fmt.Println(rec2)
}

