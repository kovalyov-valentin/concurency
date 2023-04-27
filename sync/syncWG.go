package main

import (
	"fmt"
	"sync"
)

// Должен выводить числа возводя их в квадрат
func main() {
	wg := sync.WaitGroup{}
	counter := 20 

	for i := 0; i <= counter; i++ {
		i := i
		// wg надо запускать до того как мы запускаем горутину, 
		// так как горутина ляжет на планировщик, а он ее выполнит лишь когда-нибдь
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println(i * i)
		}()
	}

	wg.Wait()

	//time.Sleep(time.Second)
}