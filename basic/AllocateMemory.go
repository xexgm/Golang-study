package basic

import (
	"fmt"
	"unsafe"
)

/**
结构体占用的内存，在编译器就已经决定了，绝不是动态分配
1.Go的结构体，只有字段本身的大小（加上可能的填充），不像java，有对象头的占用
2.内存对齐：大小 != 字段相加
3.所谓的动态大小，比如切片，它的指针只有8字节，动态的大小是分配在堆上的
*/

type Item struct {
	x int
}

type BadItem struct {
	a bool
	b int64
	c bool
}

type GoodItem struct {
	b int64
	a bool
	c bool
}

type SliceStruct struct {
	list []int
}

// TestAllocateMemory 结构体内存分配规律
func TestAllocateMemory() {
	// 1. 简单的 int 结构体
	var struct1 = Item{}
	fmt.Printf("Item size: %d bytes\n", unsafe.Sizeof(struct1))

	/**
	CPU喜欢按块读取，比如一次读取 8字节
	一个 int64 （8字节）如果跨越了两个 8字节的块，CPU就得读两次再拼接，效率低
	原本应该是 1 + 8 + 1，编译器为了让 CPU 读内存更快，会进行内存对齐
	所以现在是 1+7 + 8 + 1+7
	*/
	// 2. 内存对齐的浪费
	var struct2 = BadItem{}
	fmt.Printf("BadItem size: %d bytes\n", unsafe.Sizeof(struct2))

	/**
	8 + 1+1+6 只需要补一次
	*/
	// 3. 内存对齐的优化
	var struct3 = GoodItem{}
	fmt.Printf("GoodItem size: %d bytes\n", unsafe.Sizeof(struct3))

	/**
	只有 24 byte
	*/
	// 4. 切片结构体（即使切片里塞满数据）
	list := make([]int, 0, 10000)
	var struct4 = SliceStruct{list: list}
	fmt.Printf("SliceStruct size: %d bytes\n", unsafe.Sizeof(struct4))

	// 运行结果如下
	/**
	 	go run main.go
		Item size: 8 bytes
		BadItem size: 24 bytes
		GoodItem size: 16 bytes
		SliceStruct size: 24 bytes
	*/
}
