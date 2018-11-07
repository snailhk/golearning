package main

import (
	"fmt"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(j int) {
			for {
				fmt.Printf("hello from goroutine %d \n", i)
				a[j]++
			}
		}(i)
	}
	time.Sleep(time.Microsecond)
	fmt.Println(a)
}
