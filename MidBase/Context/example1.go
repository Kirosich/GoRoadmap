package main

import (
	"context"
	"fmt"
	"time"
)

// Context with deadline

func Example1() {
	go func() {
		d := time.Now().Add(time.Second * 1)
		ctx, cancel := context.WithDeadline(context.Background(), d)
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Ctx done")
				return
			}
		}
	}()

	fmt.Println("Waiting context")
	time.Sleep(time.Second * 3)
}
