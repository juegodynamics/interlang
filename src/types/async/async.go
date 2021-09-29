package main

import (
	"fmt"
	"time"
)

func delayCallback(callback func(), c chan struct{}, wait float64) {
	time.Sleep(time.Millisecond * time.Duration(wait))
	callback()
	c <- struct{}{}
}

func main() {
	c := make(chan struct{})

	go delayCallback(func() { fmt.Println("One") }, c, 1000)
	fmt.Println("Two")

	<-c
}
