package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//draws zigzag

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nReceived SIGTERM. Shutting Down...")
		os.Exit(1)
	}()

	var start = 0
	var end = 4
	var current = start
	var isIncrease = true

	if !isIncrease {
		current = end
	}

	for {
		if isIncrease {
			for ; start <= current; start++ {
				fmt.Printf(" ")
			}
			current++
		} else {
			for ; start <= current; start++ {
				fmt.Printf(" ")
			}
			current--
		}
		fmt.Printf("*******\n")
		start = 0
		if current == start {
			isIncrease = true
		}
		if current == end {
			isIncrease = false
		}
		time.Sleep(50 * time.Millisecond)
	}

}
