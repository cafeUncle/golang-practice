package saypac

import "fmt"

func SayBack() {
	fmt.Println("hi~")
}

// 只能执行 package main 包下的 main 方法
func main() {
	print(123)
}
