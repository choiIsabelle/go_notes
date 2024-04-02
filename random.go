package main

import (
	"fmt"
	"sync"
)

func doSomething(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Two go routines")
}

func main() {
	// var x int
	var wg sync.WaitGroup
	wg.Add(2)
	// ch := make(chan int)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		go doSomething(wg)
		// ch <- 1
	}(&wg)
	// x = <-ch
	// fmt.Println(x)
	wg.Wait()
}
