package main

import "fmt"

type WeatherType = string

const (
	Sunny	WeatherType = "sun"
	RAIN	WeatherType = "rain"
)

var s int

func init() {
	s = 1
}

const const1 = 1

var (
	const11 = 1
	const22 = 2
)

func abc(weatherType WeatherType)  {

}

// 1. 正常声明变量
var i int

// 2. 正常声明 type
type A struct {

}

// 3. 基于 type 声明 type，但只继承属性，不能继承方法。而且不能通过 {} 扩展 B 的属性。  TODO 还没有找到使用场景
type B A

// 4. 复制一个类型，完全复制 A 的所有实现  可以用于别名场景，上游依赖 C，A 有修改时不影响上游的依赖查找等声明
type C = A
// 4.1 同时可以实现类似 java 枚举的作用。对 string 类型参数的传递加一层约束
type Colors string
const (
	RED		Colors = "Red"
	BLUE	Colors = "Blue"
	WHITE	Colors = "White"
)
// setColorBetter(RED)
func setColorBetter(colors Colors)  {
}
// setColorBadCase("redxxxx啊不小心传错了")
func setColorBadCase(colors string)  {
}
// 5. 组合而非继承
type Inn struct {
	_ A
	a int
}


// 类似 3，声明了一个 type，但不是来自 struct 而是引用类型
type dictor map[string]string

type DictionaryErr string
func (e DictionaryErr) Error() string {
	return string(e)
}

type Animal2 struct {
	age int
}

type Cat2 Animal2

// 给各种 type 声明函数，此时 type 作为 receiver。
// receiver不支持继承，所以对 type interface 接口实现的其它 type 不能直接使用其函数，需要单独声明
func (a *Animal2) getAge() int {
	return a.age
}

func (c *Cat2) getAge() int {
	return c.age
}

func (c *Cat2) getFeet() int {
	return 1
}

func main() {
	m1 := map[string]string{"1": "4"}
	m2 := dictor{"123": "456"}

	println(s)

	//c4 := Cat4{}
	fmt.Println(m1)
	fmt.Println(&m1)

	fmt.Println(m1)
	fmt.Println(m2)
	abc(Sunny)
	animal := Animal2{13}
	fmt.Println(animal)
	fmt.Println(animal.getAge())

	cat := Cat2{12}
	fmt.Println(cat)
	fmt.Println(cat.getAge())
	fmt.Println(cat.getFeet())
}