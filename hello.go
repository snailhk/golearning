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

func main() {
	fmt.Println("Hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
}
