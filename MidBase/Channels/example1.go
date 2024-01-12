package main

import (
	"fmt"
	"sync"
)

// First Example
func ExampleGoroutine1() {

	var chanDefault = make(chan int) // Default chan, that can make both // read-write

	// WaitGroup for sync main goroutine and other
	var wg sync.WaitGroup
	wg.Add(1)

	// For Receiver // Writer
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			chanDefault <- i
			fmt.Printf("Gave %v \n", i)
		}
		close(chanDefault)
	}()

	wg.Add(1)

	// For Reader // Sender
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-chanDefault:
				if !ok {
					fmt.Println("Channel is closed")
					return
				}
				fmt.Printf("I got %v\n", val)
			}
		}
	}()
	wg.Wait()
}
