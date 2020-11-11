package main

import (
	"net"
	"net/http"
	"net/rpc"
)

type Panda int

/*
argType是接受的参数
reply是返回的参数
*/
func (it *Panda) Calculation(argType int, reply *int) error {
	*reply = argType + 10086
	println("接受到的是：", argType, "返回值是：", *reply)
	return nil
}

func main() {
	pd := new(Panda)
	err := rpc.Register(pd)
	if err != nil {
		println("注册失败")
	}
	rpc.HandleHTTP()
	ls, err := net.Listen("tcp", ":10086")
	if err != nil {
		println("监听端口失败")
	}
	err = http.Serve(ls, nil)
	if err != nil {
		println("启动服务失败")
	}
}
