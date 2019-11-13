package main

import "fmt"

type jacket struct {
	color string
	hood  bool
	size  int
}

func CreateJacket(_color string, _hood bool, _size int) *jacket {
	j := jacket{color: _color, hood: _hood, size: _size}
	return j
}

func main() {
	newJacket := CreateJacket("red", true, 4)
	fmt.Println(newJacket)
}
