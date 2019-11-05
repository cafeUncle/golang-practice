package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[0:len(os.Args)])
	fmt.Println(os.Args[1:len(os.Args)]) //os.Args[1:1] 只有1个元素，即只有下标0有值时，也不会报错，但会返回空集合
	//fmt.Println(os.Args[1]) //panic: runtime error: index out of range [1] with length 1
	//fmt.Println(os.Args[2:]) // panic: runtime error: slice bounds out of range [2:1]
	for i := 0; i < len(os.Args[1:]); i++ {
		fmt.Println(os.Args[i])
	}
}
