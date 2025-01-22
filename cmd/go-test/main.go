package main

import "fmt"

func main() {
	fmt.Println("Hello, this is my first demo for learning go!")

	// 一、变量
	// 1.定义变量:
	//			方式一:var 变量名 类型 = 变量值
	var name string = "jmoon"
	var age int = 18
	fmt.Println(name)
	//			方式二:变量名:= 变量值
	name2 := "jmoon2"
	fmt.Println(name2)
	//			printf()可以根据自定义format格式打印出不同的东西
	//			%T:打印出变量的类型, %d:打印数字类型, %p:打印地址
	//			&:取地址符号
	fmt.Printf("%T", name2)
	fmt.Printf("name2:%d,内存地址:%p", age, &name2)
	// 			方式三:可以一次性定义多个变量
	var (
		name3 string = "jmoon3"
		age2  int    = 18
	)
	fmt.Println(name3, age2)
	// 2.变量交换
	a := 100
	b := 200
	fmt.Println("交换前：", a, b)
	a, b = b, a
	fmt.Println("交换后：", a, b)
	// 函数的多个返回值需要多个变量去接收，如果不需要某个返回值则使用“_”下划线代替(称为匿名变量)
	_, c := test(a, b)
	fmt.Println(c)

	printConst(URL)

	testIota()
}

// 3.函数定义
// func 函数名(入参变量名 入参变量类型...) (返回变量类型...) {}
func test(a int, b int) (int, int) {
	return a, b
}

// 4.常量定义: const 常量名(大写) 常量类型 = 常量值
// 写在func外的是全局变量，卸载func内的是局部变量，如果全局变量和局部变量同名时，优先局部变量
const URL string = "https://github.com/JmoonPark/go-test.git"

func printConst(constStr string) {
	fmt.Println(constStr)
}

// 5.常量计数器: iota
//
//	在const关键字出现时将被重置为0,const内部每新增一行常量声明将使iota计数一次
func testIota() {
	const (
		a = iota
		b
		c = "haha"
		d
		e = iota
		f = 111
		g
		h = iota
	)
	const (
		i = iota
		j
	)
	// 0 1 haha haha 4 111 111 7 0 1
	fmt.Println(a, b, c, d, e, f, g, h, i, j)
}
