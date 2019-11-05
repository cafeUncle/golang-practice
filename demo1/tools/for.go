package fortools

import "fmt"

func PrintTo() {
	var i = 5
	fmt.Println(i)

	// i:=1 是短变量声明, 无需var, 和上下文的作用域都没有影响，输出 1 2 3 4 5 6 7 8 9  跳出for则销毁，下文使用i还需再次声明
	for i := 1; i < 10; i++ {
		//for ; i < 10; i++ { // 如不进行短变量声明，则会继承上文 i 的赋值  输出  5 6 7 8 9  跳出for后继续用新值10
		fmt.Println(i)
	}
	//i = 30 // 上文的for未声明短变量，继承上上文声明的值后，此处还可直接重新赋值， 无需var
	fmt.Println(i) // 如上文的for声明了短变量，未继承其上文的值5，则此处依然可以读到5的值

	// 以下四种声明方式完全等价
	s := "" // 声明简洁，通过在函数内部使用，不适合包级别的变量
	fmt.Println(s)
	var s1 string // 默认值即空字符串
	fmt.Println(s1)
	var s2 = "" // 很少用，除非声明多个变量，如下. (但也可以 var s1 s2 string 就行
	var s21, s22 = "", ""
	fmt.Println(s2)
	fmt.Println(s21, s22)
	var s3 string = "" // 显式变量类型，类型一致的情况下是冗余的信息，不一致的情况下是必须的
	fmt.Println(s3)

}

func main() { // 因为不是main包，所以实际无法运行，而且小写是私有，其他包的无法调用
	PrintTo()
}
