package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 4, 5, 6}
	fmt.Println(arr1, arr2, arr3)
	var grid [4][5]int
	fmt.Println(grid)

	for _, i := range arr3 {
		fmt.Println(i)
	}

	a := make([]int, 8, 16)
	fmt.Println("aaaa")
	fmt.Println(a)
	b := make(map[string]string)
	b["aaa"] = "cccc"
	b["s"] = strconv.Itoa(9)
	fmt.Println(b)
	fmt.Println("bbb")

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] //2,3,4,5
	s2 := s1[3:5]  //5,6
	fmt.Println(s1)
	fmt.Println(s2)
	// slice可以向后扩展，不可以向前扩展
	// s[i]不可以超越len(s)，向后扩展不可以超越底层数组cap(s)
}
