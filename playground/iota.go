package main

import "fmt"

const (
	x1 = iota
	_
	y1
	z1 = "zz"
	k1
	p1 = iota
)

func printIota() {
	fmt.Println(x1, y1, z1, k1, p1) // 0 2 zz zz 5
}

func GetValue() int {
	return 1
}

func getValueType() {
	i := GetValue()
	v := interface{}(i) // 将int转换为interface{}后使用类型断言
	switch v.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

	// 不能这样写，i是int类型，不能断言
	// 只有接口类型才可以使用类型选择
	// switch i.(type) {
	// case int:
	// 	println("int")
	// case string:
	// 	println("string")
	// case interface{}:
	// 	println("interface")
	// default:
	// 	println("unknown")
	// }
}
