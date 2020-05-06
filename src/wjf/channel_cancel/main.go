package main

import (
	"context"
	"fmt"
)

// switch和select明显不同点，switch中case的是bool，select中case的是语句
func test() {
	// 该匿名函数接收一个上下文对象
	gen := func(ctx context.Context) <-chan int {
		// 构建一个管道，赋初始值为1后，并发一个匿名函数后，返回这个管道
		dst := make(chan int)
		n := 1

		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("i exited")
					return
				case dst <- n:
					n++
				}
			}
		}()

		return dst
	}

	ctx, Cancel := context.WithCancel(context.Background()) // 创建一个可退出的上下文，
	defer Cancel()                                          //
	for n := range gen(ctx) /*传入一个可退出的上下文*/ {               // 遍历这个管道，不要用java的思想去看，gen里面用了并发，所以不会有问题
		fmt.Println(n)
		if n == 10 {
			break
		}
	}
}

func main() {
	test()
	var quit int
again:
	fmt.Scanf("%d\n", &quit)
	if quit != 0 {
		fmt.Println("quit:", quit)
		goto again
	}
}
