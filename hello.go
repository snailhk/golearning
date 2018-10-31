package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
}

func variableInitialValue() {
	var a, b int = 1, 2
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	e, f := false, "sss"
	fmt.Println(a, b, c, s)
	fmt.Println(e, f)
}

func enums() {
	//枚举类型
	const (
		cpp = iota
		_
		python
		java
		golang
	)
	fmt.Println(cpp, python, java, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	enums()
}
