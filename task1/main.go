package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var canceled = make(chan struct{})

func main() {
	t := time.Now()
	fmt.Println("Hello World!")

	go func() {
		signalChanel := make(chan os.Signal, 1)
		signal.Notify(signalChanel, syscall.SIGINT)
		for {
			s := <-signalChanel
			switch s {
			case syscall.SIGINT:
				canceled <- struct{}{}
			}
		}
	}()

	select {
	case <-time.After(10 * time.Second):
		fmt.Println("GoodBye World!")
	case <-canceled:
		fmt.Printf("Stopped by the user after %v seconds\n", time.Since(t))
	}
}
