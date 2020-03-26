package znet

import (
	"ZinxLearn/Zinx/ziface"
	"fmt"
	"net"
)

/*
连接模块
*/
type Connection struct {
	Conn     *net.TCPConn //当前连接的socket TCP套接字
	ConnID   uint32       //连接的ID
	isClosed bool         //当前连接状态
	//handleAPI ziface.HandleFunc //当前连接所绑定的处理业务方法API
	ExitChan chan bool      //告知当前连接已经退出的/停止channel
	Router   ziface.IRouter //该链接处理的方法Router

}

// NewConnection 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	con := &Connection{
		Conn:   conn,
		ConnID: connID,
		//handleAPI: callbackAPI,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return con
}

// StartReader 实现客户端连接后，读写业务逻辑
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutone is running...")
	defer fmt.Println("connID=", c.ConnID, "Reader is Exit,remote addr is ", c.RemoteAddr().String)
	defer c.Stop()
	//循环读取数据
	for {
		//读取客户端数据到buf当中，最大1024字节
		buf := make([]byte, 1024)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf faild ,err=", err)
			continue
		}

		// //调用当前连接所绑定的HandleAPI
		// err = c.handleAPI(c.Conn, buf, cnt)
		// if err != nil {
		// 	fmt.Println("ConnID", c.ConnID, "handle is error", err)
		// 	break
		// }

		//得到当前conn数据的Request请求
		req := Request{
			conn: c,
			data: buf,
		}
		//执行路由调用的方法
		//从路由当中，找到注册绑定的Conn对应的router调用
		go func(request ziface.IRquest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	fmt.Println("Conn start()... ConnID=", c.ConnID)
	//TODO:启动从当前连接读取数据的业务
	go c.StartReader()
	//TODO:启动从当前连接写入数据的业务
}

// Stop 停止连接，结束当前连接
func (c *Connection) Stop() {
	fmt.Println("Conn stop()... ConnID=", c.ConnID)
	//如果当前的链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//关闭socket连接
	c.Conn.Close()
	//回收资源
	close(c.ExitChan)
}

// GetTCPConnection 获取当前连接绑定的socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前链接模块的链接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端的TCP状态IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据,将数据发送给客户端
func (c *Connection) Send(data []byte) error {
	return nil
}
