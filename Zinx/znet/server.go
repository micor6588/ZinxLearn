package znet

import (
	"ZinxLearn/Zinx/ziface"
	"fmt"
	"net"
)

// Server 定义一个server的服务器模块,实现Iserver接口
type Server struct {
	Name      string         //服务器的名字
	IPVersion string         //服务器绑定的ip版本
	IP        string         //服务器监听的ip
	Port      int            //服务器监听的ip端口
	Router    ziface.IRouter //给当前的Server添加一个router,server注册的链接对应的处理业务

}

// // CallBackClient 定义当前客户端所绑定的handle(目前这个handle是写死的，以后优化应该由用户自定义handle)
// func CallBackClient(conn *net.TCPConn, data []byte, cnt int) error {
// 	//回显业务的实现
// 	fmt.Println("[Conn Handle] CallBackToClient ...")
// 	_, err := conn.Write(data[:cnt])
// 	if err != nil {
// 		fmt.Println("write back buf faild err=", err)
// 		return errors.New("CallBackToClient error")
// 	}
// 	return nil
// }

// Start 实现IServer接口的start方法，启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start]Server Listenner at IP :%s,Port %d,is starting \n", s.IP, s.Port)
	//开启一个go去做服务端Linster业务
	go func() {

		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		//2 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		//已经监听成功
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
		var countID uint32 //统计连接次数
		countID = 0
		//3 启动server网络连接业务,处理客户端连接业务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())
			countID++
			//将处理新连接的业务方法和conn进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, countID, s.Router)

			//启动当前业务逻辑处理
			dealConn.Start()
		}
	}()
}

// Stop 实现IServer接口的Stop方法，停止服务器
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//TODO:将其他需要清理的连接信息或者其他信息 也要一并停止或者清理

}

// Server 实现IServer接口的Server方法，运行服务器
func (s *Server) Server() {
	//启动服务器服务功能
	s.Start()

	// TODO 做一些启动服务器的额外业务

	//阻塞状态
	select {}
}

// AddRouter 路由功能：给当前的服务注册一个路由方法，供客户端链接使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succeed!")
}

// NewServer 初始化Server模块的方法
func NewServer(name string) ziface.IServer {
	server := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8080,
		Router:    nil,
	}
	return server
}
