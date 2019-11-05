package fortools

import "fmt"

func For2() {
	for i := 0; ; i-- { // 允许 i++ i-- i+=2  不允许 --i ++i
		fmt.Println(i)
		if i < -5 {
			break
		}
	}

	for {
		fmt.Println("这是无限循环")
	}
}
