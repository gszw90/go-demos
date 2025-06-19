package main

import "fmt"

func panicDefer() {
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	defer func() {
		fmt.Println("defer 2")
	}()
	defer func() {
		fmt.Println("defer 3")
	}()
	panic("panic")
}

// 并发读写map时会panic
func mapPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	m := make(map[int]int)

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {
		m[1] = 1
	}
}

// 无限递归的栈溢出
func recursion() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	var f func()
	f = func() {
		f()
	}
	f()

}

// 解空指针
func nilPanic() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var p *int
	*p = 1
}
