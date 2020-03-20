package ziface

// IServer 定义一个服务器
type IServer interface {
	//启动服务器
	Start()
	//关闭服务器
	Stop()
	//运行服务器
	Server()
}
