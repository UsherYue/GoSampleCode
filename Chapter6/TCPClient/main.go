package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	if conn, err := net.Dial("tcp4", "127.0.0.1:9999"); err != nil {
		fmt.Println(err)
	} else {
		var bts []byte = make([]byte, 1024)
		length, _ := conn.Read(bts)
		fmt.Println(string(bts[0:length]))
		conn.Write([]byte("我是客户端"))
		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte("我是客户端"))
		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte("我是客户端"))
		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte("我是客户端"))
		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte("我是客户端"))
		time.Sleep(10 * time.Millisecond)
		conn.Close()
	}
}
