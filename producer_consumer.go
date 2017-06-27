package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	tasks := make(chan string, 10)
	wg := new(sync.WaitGroup)
	wg.Add(runtime.NumCPU())
	go producer(tasks)
	go consumer(tasks, wg)
	wg.Wait()
	fmt.Println("mission end")
}

func producer(tasks chan<- string) {
	for i := 0; i < 100; i++ {
		tasks <- fmt.Sprintf("mission:%d", i)
	}
	close(tasks)
}

func consumer(tasks <-chan string, wg *sync.WaitGroup) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(i int) {
			for {
				if msg, ok := <-tasks; ok {
					fmt.Printf("process %d deal %s \n", i, msg)
				} else {
					wg.Done()
					break
				}
			}
		}(i)
	}
}
