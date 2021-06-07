package main

import "fmt"

type Animal interface {
	getHouse() string
}

// receiver 类似 java 里 this 的指向。
// receiver 不接受一个 interface. 所以无法实现抽象函数
//func (a *Animal) getHouse() string {
//	return ""
//}

type Cat struct {
	house string
}

type Dog struct {
	house string
}

func (a Cat) getHouse() string {
	return a.house
}

func (a Cat) setHouse(house string) {
	a.house = house
}

func (a *Dog) getHouse() string {
	return a.house
}

func (a *Dog) setHouse(house string) {
	a.house = house
}

func Eat() {
	cat := Cat{"123"}

	//cat.setHouse("456")
	cat.setHouse("456")

	fmt.Println(cat.getHouse())
	//fmt.Println(house.getHouse())

	//println(cat.getHouse())
	//
	//println(*(&cat.house))
	//
	//println(cat.getHouse())
	//
	//fmt.Println(cat)
}

func main() {
	Eat()
}
