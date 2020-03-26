package ziface

// IServer 定义一个服务器
type IServer interface {
	//启动服务器
	Start()
	//关闭服务器
	Stop()
	//运行服务器
	Server()

	//路由的功能：给当前的服务注册一个路由方法，供给客户端链接处理使用
	AddRouter(router IRouter)
}
