package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {fmt.Println(<-ch1)}(
			)
	ch1 <- 5
	time.Sleep(3* time.Second)
}