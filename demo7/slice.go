package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main()  {

	s := []int{2,3,4,5,6,7,8}
	printSlice(s)

	// golang 切片的底层是个单链表

	// 切片的 cap，维护的不是具体的可用容量，而是根据切片结果的开头至原切片结尾，算的一个偏移量。 即原 cap 7, 从第 3 位截到第 5 位，新 cap 是 7 - 3 = 4
	// 基于切片结果进行切片，同理，如结果从 0 截到 2，则新 cap 是 4 - 0 = 4

	// 切片的 len 就是单纯的切片中元素个数

	s = s[0:2:7]
	printSlice(s) // 2 7

	s = s[0:6:7]
	printSlice(s) // 6 7

	s = s[3:6:7]
	printSlice(s) // 3 4

	s = s[2:]
	printSlice(s) // 1 1

	s = s[1:2:2]
	printSlice(s) // 1 6

}