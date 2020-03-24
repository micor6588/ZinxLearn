package ziface

/*
Irequest接口；
实际上是把客户端请求的连接信息，和请求数据包装到一个Request当中

*/
type IRquest interface {
	//得到当前连接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}
