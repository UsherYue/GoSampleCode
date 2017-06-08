package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:9999")
	udpConn, _ := net.ListenUDP("udp", addr)
	defer udpConn.Close()
	buffer := make([]byte, 1024)
	go func(udpConn *net.UDPConn) {
		for {
			rdLen, remoteAddr, err := udpConn.ReadFrom(buffer)
			if err != nil {
				return
			} else if rdLen > 0 {
				fmt.Println("RecvFrom ", remoteAddr.String(), ":", string(buffer[0:rdLen]))
			}
		}
	}(udpConn)
	time.Sleep(10 * time.Minute)
}
