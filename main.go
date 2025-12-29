package main

import "fmt"

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	//	s1 := `第一行
	//  第二行 用了空格
	//   第三行
	//	第四行
	//结束行`
	//	fmt.Print(s1)
	//	var a = '中'
	//	fmt.Print(a)
	changeString()
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
	fmt.Print(s2)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
