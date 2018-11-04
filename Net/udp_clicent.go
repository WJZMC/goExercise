package Net

//package main
//
//import (
//	"fmt"
//	"net"
//)
//
//func main() {
//	cnn, err := net.Dial("udp", "127.0.0.1:8007")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer cnn.Close()
//
//	input := ""
//	buf := make([]byte, 1024)
//	for {
//		fmt.Scan(&input)
//		n, err := cnn.Write([]byte(input))
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		n, err = cnn.Read(buf)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Println("received ", n, "byte:", string(buf[:n]))
//	}
//}
