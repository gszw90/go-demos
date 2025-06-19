package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func printConstants() {
	fmt.Println(x, y, z, k, p)
}
