package znet

import "ZinxLearn/Zinx/ziface"

// Request 请求消息的封装定义
type Request struct {
	//已经和客户端建立好链接
	conn ziface.IConnection
	//客户端的请求数据
	data []byte
}

// GetConnection 得到当前连接
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// GetData 得到当前用户数据
func (r *Request) GetData() []byte {
	return r.data
}
