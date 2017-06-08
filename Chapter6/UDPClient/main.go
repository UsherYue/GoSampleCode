package main

import (
	"net"
	"time"
)

func main() {
	udpConn, _ := net.Dial("udp", "127.0.0.1:9999")
	udpConn.Write([]byte("hello,i am udp client"))
	time.Sleep(10 * time.Millisecond)
	udpConn.Write([]byte("hello,i am udp client"))
	time.Sleep(10 * time.Millisecond)
	udpConn.Write([]byte("hello,i am udp client"))
	time.Sleep(10 * time.Millisecond)
	udpConn.Write([]byte("hello,i am udp client"))
	time.Sleep(10 * time.Millisecond)
	udpConn.Write([]byte("hello,i am udp client"))
	time.Sleep(10 * time.Millisecond)
	udpConn.Write([]byte("hello,i am udp client"))
	defer udpConn.Close()
}
