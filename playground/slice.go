package main

import "fmt"

type slicePerson struct {
	name string
	age  int
}

func nilSlice() {
	s := new([]int)
	*s = append(*s, 1, 2, 3)
	println(*s)
}

func sliceMap() {
	// 切片不能作为map的key
	// s:= []int{1, 2, 3}
	// m:= make(map[[]int]int)

	// 数组可以作为map的key
	s := [3]int{1, 2, 3}
	m := make(map[[3]int]int)
	m[s] = 1
	println(m[s])
}

func splitSlice() {

	s := []int{1, 2, 3, 4}
	s1 := s[1:2:3] // 容量计算3-1
	s2 := s[1:2]   // 容量计算4-1

	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
}

func rangeSlice() {
	s := []int{1, 2, 3}
	for _, v := range s {
		v += 1
	}
	fmt.Println(s)

	s2 := []*slicePerson{
		{name: "Tom", age: 18},
		{name: "Jack", age: 20},
		{name: "Lucy", age: 22},
	}
	for _, v := range s2 {
		v.age += 1
	}
	for _, person := range s2 {
		fmt.Println(person)
	}
}

func pointerSlice() {
	list := new([]int)
	// list = append(list, 1)
	*list = append(*list, 1)
	fmt.Println(list)
}
