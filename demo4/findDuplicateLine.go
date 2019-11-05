package main

import (
	"bufio"
	"fmt"
	"os"
)

func find() {
	// 或 counts := map[string]int{}
	counts := make(map[string]int) // map存储一个键值对集合，此处 key 是 string, val 是 int

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}
	// 此处忽略 input.Err()中可能的错误

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
