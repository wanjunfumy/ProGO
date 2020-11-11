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
	// 这里直接暴露接口，不知道有没有其他的办法隐藏起来。如果是服务器能不，应该就不需要了。如果rpc是链接终端？那就不让他们链接到终端，
	err = cli.Call("Panda.Calculation", 10089, &reply)
	if err != nil {
		println("接口请求失败")
	}
	println("服务返回的值：", reply)
}
