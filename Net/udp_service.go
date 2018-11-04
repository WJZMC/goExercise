package Net

//package main
//
//import (
//	"fmt"
//	"net"
//)
//
//func main() {
//	udpaddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8007")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	conn, err := net.ListenUDP("udp", udpaddr)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer conn.Close()
//
//	for {
//
//		buf := make([]byte, 1024)
//
//		n, addr, err := conn.ReadFromUDP(buf)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println(string(buf[:n]))
//
//		n, err = conn.WriteToUDP(buf[:n], addr)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
//
//}
