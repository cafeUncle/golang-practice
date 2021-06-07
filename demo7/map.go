package main

import (
	"errors"
	"fmt"
)

func Search(dict map[string]string, key string) (result string, err error) {
	result = dict[key]
	// 指针可能 nil，值不会 nil? 看起来可能是每个 go 的引用，都至少有个 零值，即使 var Animal animal，不声明对象，也依然会有 Animal{} 的零值填充
	if result == "" {
		err = errors.New("key not exists")
		panic(err)
	}
	return
}

func main() {

	map1 := make(map[string]string)
	fmt.Println(map1)

	map2 := map[string]string{}
	fmt.Println(map2)

	map3 := map[string]string{"123": "456"}

	map3["345"] = "789"

	fmt.Println(map3)
	fmt.Println("345 's value:", map3["345"])
	fmt.Println("111 's value:", map3["111"])

	result, _ := Search(map3, "345")
	fmt.Println("search 345 's value:", result)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover err:", r)
		}
	}()

	result, _ = Search(map3, "111")
	fmt.Println("search 111 's value:", result)


	fmt.Println("there will not be shown forever because error")





}
