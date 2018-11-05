package main

import (
	"bufio"
	"fmt"
	"learning/errorhandling/fib"
	"os"
)

func tryDefer() {
	for i := 0; i < 30; i++ {
		defer fmt.Println(i)

		if i == 30 {
			panic("printed to many!")
		}
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	writeFile("fib.txt")
}
