package main

import (
	"fmt"
	"os"
	"os/signal"
)

//Example 2 - Os.Signal Exit, give smth signal to end main
func ExampleGoroutine2() {
	chDone := make(chan os.Signal, 1)
	signal.Notify(chDone, os.Interrupt, os.Kill)

	<-chDone
	fmt.Println("Got signal")
}
