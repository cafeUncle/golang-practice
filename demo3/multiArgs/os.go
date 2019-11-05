package main

import (
	"fmt"
	"os"
)

func main() {
	// 关于变量声明，通常是 var name string 或 var name = "LiLei"
	// 可以用 := 来省略写法，如 name := "LiLei"
	// 对已声明的变量再次赋值时，就可以直接用 = 了

	// 如果是函数外的声明，不能省略，只能用 var
	// 函数外的每个声明，都必须有关键字前缀，如 var 或 func，所以 := 不能在外边用

	// 常量通常声明在函数外，同样也不能用 := 声明，需要用 const 关键字来指明它是一个常量

	var i string = "A"
	fmt.Println(i)

	// go的变量声明后可以不赋值，数值类型默认为0，字符串类型默认为空字符串 ""
	// 但要么赋值，要么声明类型，不能只 var i，影响下文编译正确性
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
