package main

import "fmt"

func greetingPrefix(prefix string) (finalPrefix string) {
	switch prefix {
	case "japan":
		finalPrefix = "japanese"
	case "china":
		finalPrefix = "chinese"
	case "france":
		finalPrefix = "france"
	default:
		finalPrefix = "un_country"
	}
	return
}

func Hello() string {
	return "Hello, world"
}

func split()  {
	arr := []int{1,2,3,4,5}
	fmt.Println(arr)
	fmt.Printf("%v\n", arr)
	fmt.Println(Sum(arr))
}

func Sum(arr []int) (sum2 int) {
	//for i:=0;i< len(arr); i++ {
	//	sum2 += arr[i]
	//}

	//for i := range arr {
	//	sum2 += arr[i]
	//}

	for index,val := range arr {
		fmt.Println(index)
		sum2 += val
	}
	return sum2
}

func main() {
	fmt.Println(Hello())
	fmt.Println(greetingPrefix("chinese"))
	split()
}