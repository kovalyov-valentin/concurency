package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	alreadyStorage := make(map[int]struct{})
	mu := sync.Mutex{}
	capacity := 1000 

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10))
	}

	uniqueIDs := make(chan int, capacity)
	wg := sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			if _, ok := alreadyStorage[doubles[i]]; !ok {

				alreadyStorage[doubles[i]] = struct{}{}

				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	close(uniqueIDs)
	for val := range uniqueIDs {
		fmt.Println(val)
	}

	fmt.Printf("len of ids: #{len(uniqueIDs)}\n")
	fmt.Println(uniqueIDs)
}