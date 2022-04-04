package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func goroutines() {
	// adding a total of goroutines
	wg.Add(2)
	go PrintName("Hi")
	go PrintName("Hello")

	// wait until all goroutines are finished
	wg.Wait()

	wg.Add(1)
	go PrintName("What's up!")
	wg.Wait()
}

func PrintName(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

	wg.Done()
}
