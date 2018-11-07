package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	//done chan bool
	done func()
}

func doWorker(id int, c chan int, wg worker) {
	for n := range c {
		fmt.Printf("worker %d received %c \n", id, n)
		wg.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, w.in, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	wg.Wait()

	// wait for all of them
	//for _, worker := range workers {
	//	<- worker.done
	//	<- worker.done
	//}
}

func main() {
	chanDemo()
	//bufferedChannel()
	//channelClose()
}
