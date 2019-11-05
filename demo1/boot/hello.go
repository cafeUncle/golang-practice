package main

import (
	"fmt"
	"go.practice.cafeUncle.com/demo1/tools"
)

// 整体来说，go的包分为main包即启动包，引入包即用于引用的包。
// 前者必须是main包且有且只有一个main方法
// 后者包名不允许为main，且不能有main方法

// 每个试图 go build或install的文件夹下，必须有一个go文件，由文件import去关联包  【结果不一致】   go build有时无输出，有时输出到bin，有时输出到pkg
// 每个试图 go build或install的文件夹下，必须有且只有一个main包及其main方法。    【结果不一致】   go install，有时输出到bin，有时输出到pkg
// 每个文件夹下的所有go文件，不允许存在多个包名。
// 所以同一个文件夹下的文件，都属于一个包，所以不同文件内的方法名也不允许相同
// 想运行的包名必须是main，否则只能被引入。
// 引用的包中允许有其他main包或main方法，小写开头，所以不允许引用
// 所以每个文件夹下最多只有一个go文件可以包含并运行main方法
func main() {
	fmt.Print("Hello, World!\n")

	fortools.PrintTo()

	//saypac.SayBack()
	//
	//fortools.For2()
}
