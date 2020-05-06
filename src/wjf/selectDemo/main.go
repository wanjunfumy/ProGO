package main

import (
	"fmt"
	"time"
)

func selectTest() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		select { // select在选择case时候，是随机且可重复的
		case c <- 1:
		case c <- 2:
		case c <- 3:
		case c <- 4:
		case c <- 5:
		}
	}
	close(c)
	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	//selectTest()
	timeForChan()
}

var _time = time.Tick(time.Duration(1))

func timeForChan() {
	a := 0
	for {
		<-_time // 这里确实会阻碍到main的gorouting，可是，tick似乎不会生效，Duration值大小不会影响到打印的快慢
		fmt.Println("我是风阁阁主！", a)
		a++
	}
}
