package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		default:
			fmt.Println("default")
		}
	}()
	
	time.Sleep(time.Second)

	fmt.Println(runtime.NumGoroutine())
}