package main

import (
	"Golang-studyDemo/basic"
	"fmt"
	"math"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	/**
	案例1 变量的定义
	*/
	//	s1 := `第一行
	//  第二行 用了空格
	//   第三行
	//	第四行
	//结束行`
	//	fmt.Print(s1)
	//	var a = '中'
	//	fmt.Print(a)

	/**
	案例2 一维数组 和 二维数组的定义
	*/
	//changeString()
	//sqrtDemo()
	//fmt.Println()
	//basic.GetArr1()
	//basic.GetArr2()

	/**
	案例3 函数传递的是 值，而非地址
	*/
	//a := [2]int{}
	//fmt.Printf("a: %p", &a)
	//basic.TestCopy(a)
	//// 打印 a 的值
	//fmt.Println(a)
	//// len() 和 cap() 都是内置函数，返回数组长度
	//fmt.Println(len(a), cap(a))

	/**
	案例4 多维数组遍历
	*/
	//basic.Traversal()

	/**
	案例5 数组拷贝和传参
	*/
	//a := [5]int{}
	//basic.PrintArr(&a)
	//fmt.Println(a)
	//
	//b := [5]int{1, 2, 3, 4, 5}
	//basic.PrintArr(&b)
	//fmt.Println(b)

	/**
	案例6 求出数组元素之和
	*/
	//rand.Seed(time.Now().Unix())
	//var b [2]int
	//for i, _ := range b {
	//	// 产生一个 0-1000 的随机数
	//	b[i] = rand.Intn(10)
	//}
	//fmt.Println(b)
	//sum := basic.SumArr(&b)
	//fmt.Println(sum)

	/**
	找出数组中和为定值的两个元素的下标
	*/
	b := [5]int{1, 3, 5, 8, 7}
	basic.TwoIndex(8, &b)
}

func changeString() {
	// 强制类型转换
	s1 := "hello world"
	bytes1 := []byte(s1)
	bytes1[0] = 'd'
	fmt.Print(string(bytes1))

	fmt.Println()

	// 用 rune 转换
	s2 := "博客"
	bytes2 := []rune(s2)
	bytes2[0] = '光'
	fmt.Print(string(bytes2))
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt() 接收的参数是 float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
