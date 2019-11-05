package main

import "fmt"

type ICompare func(s *string, s2 *string) bool

func (ic *ICompare) do() {
	fmt.Println("123")
}

func ComparatorString(s *string, s2 *string) bool {
	return &s == &s2
}

func writeSthToWhere() {
	ic := ICompare(ComparatorString)
	// 可以用于重载某个类型的方法的场景? 还没有想到为啥
	ic.do()
}
