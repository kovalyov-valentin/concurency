package main

import (
	"fmt"
	"runtime"
	"time"
)

func main_() {
	// В данном примере с небуферизированным каналом будет deadlock, так как нет второй горутины, которая бы читала из канала. 
	// У него нет буфера, куда данные сложить

	// ch := make(chan int) // Небуферизированный канал
	ch := make(chan int, 1) // Буферизированный канал
	fmt.Println(runtime.NumGoroutine())

	ch <- 1
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	go func ()  {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second)
			fmt.Println("iter ", i)
		}
		fmt.Println(runtime.NumGoroutine())
	}()
	select{}
	fmt.Println("finish")
}