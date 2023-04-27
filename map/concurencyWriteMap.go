package main

import (
	"fmt"
	"sync"
)

func main() {
	writes := 1000
	// var storage map[int]int
	storage := make(map[int]int, writes)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			mu.Lock()
			defer mu.Unlock()
			storage[i] = i
		}()
	}

	wg.Wait()
	fmt.Println(storage)
}