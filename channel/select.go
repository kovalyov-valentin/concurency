package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	chanForResp := make(chan resp)
	go RPCCall(ctx, chanForResp)

	// select {
	// case <-ctx.Done():
	// 	fmt.Println()
	// case result := <-chanForResp:
	// 	fmt.Println(result)
	// }

	// <-ctx.Done()

	// cancel()

	res := <- chanForResp
	fmt.Println(res.id, res.err)
}

type resp struct {
	id  int
	err error
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			id:  0,
			err: errors.New("timeour expired"),
		}
	case <-time.After(time.Second * time.Duration(rand.Intn(5))):
		// sleep 0-4
		ch <- resp{
			id: rand.Int(),
		}
	}
}
