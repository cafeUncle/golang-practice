package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	simpleJoin()
	stringsJoin()
	find()
}

func simpleJoin() {
	start := time.Now()
	fmt.Println("simpleJoin start:", start)
	var result, sep string
	for i := 0; i < len(os.Args); i++ {
		result += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(result)
	end := time.Now()
	fmt.Println("simple耗时：", end.Nanosecond()-start.Nanosecond())
}

func stringsJoin() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, ""))
	end := time.Now()
	fmt.Println("strings耗时：", end.Nanosecond()-start.Nanosecond())
}
