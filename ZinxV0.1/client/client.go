package main

import (
	"fmt"
	"net"
	"time"
)

/*
模拟客户端
*/

func main() {
	fmt.Println("client starting ....")
	//1.直接远程连接服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "192.168.73.1:8080")
	if err != nil {
		fmt.Println("client start server faild err=", err)
		return
	}
	for {
		//2.直接调用write写数据
		_, err := conn.Write([]byte("Hello Zinx V0.1"))
		if err != nil {
			fmt.Println("write conn faild err=", err)
			return
		}
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("write conn faild err=%v\n", err)
			return
		}
		fmt.Printf("server call back :%s,cnt=%d\n", buf, cnt)
		//cpu阻塞
		time.Sleep(1 * time.Second)
	}

}
