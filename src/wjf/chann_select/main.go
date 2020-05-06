package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	c = make(chan int, 100000)
	for i := 0; i < 100000; i++ {
		c <- i
	}

	go func() {
		a := 0
		for {
			a++
			c <- a
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		default:
			fmt.Println("just wait")
			time.Sleep(time.Second)
		}
	}
}
