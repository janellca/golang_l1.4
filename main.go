package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var n int
	fmt.Fscan(os.Stdin, &n)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()

	ch := make(chan int)
	for i := 0; i < n; i++ {
		go func(ch chan int) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("worker is stoped")
					return
				case value, ok := <-ch:
					if !ok {
						fmt.Println("channel is closed")
					}
					fmt.Println(value)
				}
			}
		}(ch)
	}
	x := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker is stoped")
			close(ch)
			return
		default:
			ch <- x
			x++
			time.Sleep(1 * time.Second)
		}

	}

}
