package main

import "fmt"

type mapPerson struct {
	name string
	age  int
}

func nilMap() {
	m := make(map[int]int, 2)
	fmt.Println(len(m))
	// fmt.Println(cap(m))
}

func rangeMap() {
	m := map[int]*mapPerson{
		1: &mapPerson{
			name: "Tom",
			age:  18,
		},
		2: &mapPerson{
			name: "Jack",
			age:  20,
		},
		3: &mapPerson{
			name: "Lucy",
			age:  22,
		},
	}
	for _, v := range m {
		v.age += 1

	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
