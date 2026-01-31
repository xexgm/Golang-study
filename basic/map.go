package basic

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func UseMap() {
	var scoreMap map[string]int

	scoreMap = make(map[string]int, 10)

	scoreMap["光芒"] = 100
	scoreMap["gm"] = 99

	fmt.Println(scoreMap)
	fmt.Printf("scoreMap len: %d\n", len(scoreMap))

	var userInfo map[string]string = map[string]string{
		"username": "xexgm",
		"password": "123456",
	}
	fmt.Println(userInfo)
	fmt.Printf("userInfo len: %d\n", len(userInfo))

	// map 也支持在声明的时候填充元素
	userInfo2 := map[string]string{
		"username": "xexgm",
		"password": "123456",
	}
	fmt.Println(userInfo2)
}

func IsExist() {
	var scoreMap = make(map[string]int, 10)
	scoreMap["gm"] = 100
	scoreMap["xgm"] = 99

	// 如果 key 存在，ok为true，不存在 ok 为 false，v为值类型的零值
	v, ok := scoreMap["gm"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	// golang 中 使用 for range 遍历 map
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	for k := range scoreMap {
		fmt.Println(k)
	}
}

func TestDelete() {
	scoreMap := make(map[string]int, 10)
	scoreMap["gm"] = 100
	scoreMap["xgm"] = 99

	delete(scoreMap, "xgm")

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
}

func RangeBySort() {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 随机构造 map
	scoreMap := make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		val := rand.Intn(101)
		scoreMap[key] = val
	}

	//  将map的key 存入 切片
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 切片进行排序
	sort.Strings(keys)

	// 根据排序后的切片，取map的元素
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// MapSlice 类型为map的切片
func MapSlice() {
	var mapSlice = make([]map[string]string, 3)

	for idx, val := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", idx, val)
	}

	// init
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["gm"] = "100"
	mapSlice[0]["xgm"] = "100"

	for idx, val := range mapSlice {
		fmt.Printf("index: %d, value: %v\n", idx, val)
	}
}

// SliceMap 元素为切片的map
func SliceMap() {
	// 三个元素的map
	SliceMap := make(map[string][]string, 3)

	key := "中国"
	val, ok := SliceMap[key]
	if !ok {
		val = make([]string, 5)
	}
	val = append(val, "北京", "上海")
	SliceMap[key] = val
	fmt.Println(SliceMap)
}
