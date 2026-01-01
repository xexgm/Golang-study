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
	案例7 找出数组中和为定值的两个元素的下标
	*/
	//b := [5]int{1, 3, 5, 8, 7}
	//basic.TwoIndex(8, &b)

	/**
	案例8 切片练习
	*/
	//arr := [5]int{1, 2, 3, 4, 5}
	//var s = arr[:]
	//fmt.Printf("原切片: 指针的地址%p", s)
	//fmt.Print(s)
	//fmt.Println()
	//
	//basic.ModifySli(s)
	//
	//fmt.Print("后来的原切片: ")
	//fmt.Print(s)

	/**
	案例9 切片的 len 和 cap
	*/
	//basic.InitSlice()

	/**
	案例10 通过 make动态创建slice，避免了数组必须用常量做长度的麻烦，还可直接用指针访问底层数组
	退化成普通数组操作
	以下是数组操作的麻烦之处
	*/
	//s := [...]int{0, 1, 2, 3, 4}
	//fmt.Println(s)
	//p := &s[1]
	//p1 := &s[2]
	//*p = 666
	//*p1 = 777
	//fmt.Println(s)

	/**
	案例11 二维数组，[][]T,的含义是，第一个[] 代表了数组，它的每一个元素，都是 []T
	*/
	//data := [][]int{
	//	[]int{1, 2, 3},
	//	[]int{4, 5, 6},
	//	[]int{7, 8, 9},
	//}
	//fmt.Println(data)

	/**
	案例12 结构体数组，第一个 {} 代表了 结构体的结构，第二个{},是用来初始化的
	*/
	//data := [5]struct {
	//	x int
	//}{
	//	{1},
	//	{2},
	//	{3},
	//}
	//fmt.Println(data)
	//data[0].x = 100
	//data[1].x = 200
	//data[2].x = 300
	//fmt.Println(data)

	/**
	案例13 使用 append 操作切片
	*/
	//basic.AppendSlice()

	/**
	案例14 slice中 cap 的分配规律
	*/
	//basic.SliceCap()

	/**
	案例15 slice 调整大小
	*/
	//basic.SliceResize()

	/**
	案例16 字符串和切片
	*/
	//basic.StringAndSlice()

	/**
	案例17 指针
	*/
	//basic.TestPtr()

	/**
	案例18 指针取值
	*/
	//basic.GetValByPtr()

	/**
	案例19 空指针
	*/
	//basic.NilPtr()

	/**
	案例20 new 练习
	*/
	//basic.TestNew()

	/**
	案例21 make 练习
	*/
	//basic.TestMake()

	/**
	案例22 内存分配 练习
	*/
	basic.TestAllocateMemory()
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
