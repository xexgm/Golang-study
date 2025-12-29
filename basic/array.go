package basic

import "fmt"

// 全局初始化
var arr0 [5]int = [5]int{1, 2, 3, 4, 5}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var arr3 = [5]string{2: "apple", 3: "banana"}

var arr4 [5][3]int
var arr5 [2][3]int = [...][3]int{{1, 2, 3}, {5, 6, 7}}

// GetArr1 一维数组
func GetArr1() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age  int
	}{
		{"xexgm", 22},
		{"xedgm", 23}, // 最后一个 , 不可忽视
	}
	fmt.Println(arr0, arr1, arr2, arr3)
	fmt.Println(a, b, c, d)
}

func GetArr2() {
	a := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b := [...][2]int{{1, 1}, {2, 2}, {3, 3}}

	fmt.Println(arr4, arr5)
	fmt.Println(a, b)
}

// TestCopy 传进来的是值，而非地址, 会造成性能问题，建议使用 slice 或 数组指针
func TestCopy(x [2]int) {
	fmt.Printf("x: %p", &x)
	x[0] = 1000
	x[1] = 1000
}

// Traversal 多维数组遍历
func Traversal() {
	f := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// k1 代表行数
	// v1 代表第一行的 一维数组
	// k2 代表列数
	// v2 代表由 行列 锁定的 值
	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(k1:%d, k2:%d) = v2:%d ", k1, k2, v2)
		}
		fmt.Println()
	}
}

func PrintArr(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

// SumArr 求出数组所有元素之和
func SumArr(arr *[2]int) (sum int) {
	//fmt.Printf("地址: %p", arr)

	// 这里用 *arr 和 arr 都是一样的，但是建议用 *arr，是因为编译器有对数组指针做优化
	// 最好自己解引用！

	for _, v := range *arr {
		sum += v
	}
	return
}

// TwoIndex 找出数组中，和为定值的两个元素的下标
func TwoIndex(target int, arr *[5]int) {
	// 从左到右遍历
	// 根据 i，求 other
	// 往右边找，如果找得到，打印

	for i := 0; i < len(arr); i++ {
		other := target - arr[i]

		for j := i + 1; j < len(arr); j++ {
			if arr[j] == other {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}
