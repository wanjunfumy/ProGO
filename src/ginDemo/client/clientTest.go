package main

import "net/rpc"

// rpc通信
func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		println("链接服务失败")
	}
	if cli == nil {
		println("没有发现可用服务")
		return
	}
	var reply int
	err = cli.Call("Panda.Calculation", 10089, &reply)
	if err != nil {
		println("接口请求失败")
	}
	println("服务返回的值：", reply)
}
