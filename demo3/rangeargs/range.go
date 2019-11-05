package main

import (
	"fmt"
	"os"
)

func main() {
	var i = 0

	fmt.Println(i)
	//fmt.Println(j)

	for i == i {
		fmt.Println(1)
		break
	}

	//strings := os.Args[0:]
	//for strings { // non-bool 'strings' (type string[]) used as condition   不能直接放数组
	//	fmt.Println(os.Args)
	//}

	//fmt.Println(range os.Args)  // 表达式不可直接被输出

	for e, b := range os.Args { // 每次迭代，range产生两个值，1索引下标值 2索引元素值
		fmt.Println(e, "=", b)
	}

	for e := range os.Args { // 只用一个元素接收则只能接收到索引下标值
		fmt.Println(e)
	}

	// go不允许存在无用的临时变量，不允许通过编译。
	// 可以用空标志符 _ 下划线，用在语法需要变量名但程序逻辑不需要的地方，例如每次迭代产生的无用的索引
	for _, b := range os.Args {
		//fmt.Println(_) // can't use '_' as value
		fmt.Println(b)
	}

}
