package main

import (
	"fmt"
)

func calc(intChan chan int, result chan int) {
	for v := range intChan {
		flag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			result <- v
		}
	}
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	go func() {
		for i := 0; i < 100000; i++ {
			intChan <- i
		}
		close(intChan)
	}()
	//利用管道的特性，取出来就没有了，所以，我们开8个线程同时处理一个chan
	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan)
	}
	for {
		v, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println("resultChan len:", len(resultChan), v)
	}

}
