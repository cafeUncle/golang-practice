package main

import (
	"fmt"
)

const CONST1 = 1

var (
	CONST11 = 1
	CONST22 = 2
)

// 正常声明变量
var i int

// 这是啥语法？ dictor 是变量还是 struct
type dictor map[string]string

// 同上
type DictionaryErr string

// 字段的话，为啥还能声明函数
func (e DictionaryErr) Error() string {
	return string(e)
}

type Animal struct {
	age int
}

// struct 的话，这种算继承么。看起来字段可以继承下来，但方法不能
type Cat Animal

func (a *Animal) getAge() int {
	return a.age
}

func (c *Cat) getAge() int {
	return c.age
}

func (c *Cat) getFeet() int {
	return 1
}

func main() {
	m1 := map[string]string{"1": "4"}
	m2 := dictor{"123": "456"}

	fmt.Println(m1)
	fmt.Println(&m1)

	fmt.Println(m1)
	fmt.Println(m2)

	animal := Animal{13}
	fmt.Println(animal)
	fmt.Println(animal.getAge())

	cat := Cat{12}
	fmt.Println(cat)
	fmt.Println(cat.getAge())
	fmt.Println(cat.getFeet())
}

//Map 有一个有趣的特性，不使用指针传递你就可以修改它们。这是因为 map 是引用类型。这意味着它拥有对底层数据结构的引用，就像指针一样。它底层的数据结构是 hash table 或 hash map，你可以在这里阅读有关 hash tables 的更多信息。
//Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。
//引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。
//由于 nil 指针异常，你永远不应该初始化一个空的 map 变量：
//var m map[string]string
//相反，你可以像我们上面那样初始化空 map，或使用 make 关键字创建 map：
//dictionary = map[string]string{}
// OR
//dictionary = make(map[string]string)
