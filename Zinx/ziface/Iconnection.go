// 定义连接模块的抽象层
package ziface

import "net"

// IConnection 定义服务器连接的相关接口
type IConnection interface {
	Start() //启动连接，让当前连接开始工作

	Stop() //停止连接，结束当前连接

	GetTCPConnection() *net.TCPConn //获取当前连接绑定的socket conn

	GetConnID() uint32    //获取当前链接模块的链接ID
	RemoteAddr() net.Addr //获取远程客户端的TCP状态IP port

	Send(data []byte) error //发送数据,将数据发送给客户端

}

// HandleFunc 定义一个处理连接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
