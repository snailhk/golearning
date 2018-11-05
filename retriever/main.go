package main

import (
	"fmt"
	"learning/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	//var w io.ReadWriteCloser
	var r Retriever
	r = mock.Retriever{}
	fmt.Println(download(r))
}
