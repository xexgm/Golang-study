package basic

import "fmt"

// ModifySli 接收一个切片，这个切片是原切片的 副本
func ModifySli(sli []int) {
	sli[0] = 666 // 切片的 ptr 指向的第一个元素，进行修改，底层数组也会修改
	fmt.Printf("append前 切片结构体中的指针:%p", sli)

	/**
	append 如果容量不够，会发生扩容，切片的指针，会指向新的数组，
	*/
	sli = append(sli, 777) // 追加一个元素
	fmt.Printf("append后 切片结构体中的指针:%p", sli)
	// 打印
	fmt.Print("切片的副本: ")
	fmt.Print(sli)
	fmt.Println()
}

// MakeSli 创建切片的各种方式
func MakeSli() {
	// 1.i声明切片
	var s1 []int
	if s1 == nil {
		fmt.Printf("空")
	} else {
		fmt.Printf("非空")
	}
	// 2. :=
	s2 := []int{}
	fmt.Println(s2)
	// 3.make()          类型 长度 容量
	var s3 []int = make([]int, 0, 3)
	fmt.Println(s3)
	// 4.初始化赋值
	var s4 = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr[1:4]
	fmt.Println(s6)
}

// InitSlice 切片初始化
func InitSlice() {
	var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 = arr[:]
	var slice1 = arr[2:8]
	fmt.Printf("slice0的长度为:%d, 容量为:%d", len(slice0), cap(slice0))
	fmt.Println()
	fmt.Printf("slice1的长度为:%d, 容量为:%d", len(slice1), cap(slice1))
	fmt.Println()
	var slice2 = make([]int, 2, 5)
	fmt.Println(slice2)
	fmt.Println()

	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 通过make来创建切片，指定len和cap
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 通过make来创建切片，省略cap，那么 cap == len
	fmt.Println(s3, len(s3), cap(s3))
}

func AppendSlice() {
	var a = []int{0, 1, 2}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{3, 4, 5}
	fmt.Printf("slice b : %v, ptr:%p, slice:%p\n", b, b, &b)
	var c = append(a, b...)
	fmt.Printf("slice c : %v, ptr:%p\n", c, c)
	var d = append(c, 6)
	fmt.Printf("slice d : %v, ptr:%p\n", d, d)
	var e = append(d, 7, 8, 9)
	fmt.Printf("slice e : %v, ptr:%p\n", e, e)
	b = append(b, 6)
	fmt.Printf("slice b : %v, ptr:%p, slice:%p\n", b, b, &b)
	/**
	可以观察到，append后，底层数组扩容了，b切片，底层的数组又创建了一个新的数组，所以指针指向不一样了
	*/
}

// SliceCap slice中cap的分配规律,结论就是，cap每次都 * 2
func SliceCap() {
	s := make([]int, 0, 1) // len为0，cap为1
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		// 如果 切片的容量比之前大了，那就打印
		if n := cap(s); n > c {
			fmt.Printf("slice cap:%d -> %d\n", c, n)
			c = n
		}
	}
}

// SliceResize 切片调整大小
func SliceResize() {
	array := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("slice array : %v\n", array)
	b := array[1:2]
	fmt.Printf("slice b : %v", b)
	c := b[0:3]
	fmt.Printf("slice c : %v", c)
}

// StringAndSlice string底层是一个byte数组，所以可以进行切片操作
func StringAndSlice() {
	str := "hello world"
	s1 := str[:5]
	fmt.Println(s1)
	s2 := str[6:]
	fmt.Println(s2)

	// string 本身是不可变的，如果要改变string中的字符,需要如下操作

	// 创建一个新的byte切片
	s := []byte(str)
	s = s[:6]
	s = append(s, 'x', 'e', 'x', 'g', 'm')
	str = string(s)
	fmt.Println(str)
}
