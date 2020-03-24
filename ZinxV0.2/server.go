package main

import "ZinxLearn/Zinx/znet"

/*
基于Zinx框架来开发的服务端应用程序

*/

func main() {
	//1.创建一个server句柄，使用ZinX框架的API
	server := znet.NewServer("[zinx V0.2]")
	//2.启动server
	server.Server()
}
