package Net

//package main
//
//import (
//	"fmt"
//	"net"
//	"strings"
//)
//
//func handelConnect(cnn net.Conn) {
//	defer cnn.Close()
//
//	buf := make([]byte, 1024)
//	for {
//		n, err := cnn.Read(buf)
//		if strings.Contains(string(buf[:n]), "exit") {
//			fmt.Println("服务端处理[", cnn.RemoteAddr(), "]客户端退出")
//			return
//		}
//		if n == 0 {
//			fmt.Println("[", cnn.RemoteAddr(), "]客户端暴力退出")
//			return
//		}
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("service received:[", cnn.RemoteAddr(), "]:", string(buf[:n]))
//
//		dealResult := "service replay [" + cnn.RemoteAddr().String() + "] " + string(buf[:n])
//
//		_, err = cnn.Write([]byte(dealResult))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//}
//func main() {
//	lister, err := net.Listen("tcp", "127.0.0.1:8004")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer lister.Close()
//
//	for {
//		cnn, err := lister.Accept()
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("客户端【", cnn.RemoteAddr(), "】连接成功")
//		go handelConnect(cnn)
//	}
//}
