package basic

import "fmt"

func TestPtr() {
	a := 10
	ptr := &a
	fmt.Printf("a的值: %d, a的地址: %p\n", a, &a)
	// %p 表示，把值当作 十六进制的内存地址来打印
	fmt.Printf("ptr的值:%p, ptr的地址: %p", ptr, &ptr)
}

func GetValByPtr() {
	a := 10
	p := &a
	c := *p

	fmt.Printf("a的值: %d, a的地址: %p\n", a, &a)
	fmt.Printf("p的值: %p, p的地址: %p\n", p, &p)
	fmt.Printf("c的值: %d, c的地址: %p\n", c, &c)
}

func NilPtr() {
	var p *string
	fmt.Println(p)
	fmt.Printf("p的值是: %p\n", p)

	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空")
	}
}

func TestNew() {
	a := new(int)
	b := new(bool)
	fmt.Printf("a的类型: %T\n", a)
	fmt.Printf("b的类型: %T\n", b)
	fmt.Printf("a的值: %d\n", *a)
	fmt.Printf("b的值: %t\n", *b)

	var c *int
	fmt.Printf("c的值: %p, c的地址: %p\n", c, &c)
	c = new(int)
	fmt.Printf("c指向的值: %d, c的值: %p, c的地址: %p\n", *c, c, &c)
}

func TestMake() {
	var b map[string]int
	b = make(map[string]int, 10)
	b["光芒"] = 100
	fmt.Println(b)
}
