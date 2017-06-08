package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"time"
)

func Recv(conn net.Conn) {
	defer conn.Close()
	conn.Write([]byte("hello,client,I am Server Echo...."))
	for {
		if bts, err := ioutil.ReadAll(conn); err != nil {
			fmt.Println(conn.RemoteAddr().String(), " close connect!")
			break
		} else {
			//当读取到数据的时候
			if len(bts) > 0 {
				fmt.Println("Recv:", string(bts))
			}
		}
	}
}

func Accept(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Accect Connect:", conn.RemoteAddr().String())
		go Recv(conn)
	}
}

func main() {
	//解析TCP地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	l, _ := net.ListenTCP("tcp", tcpAddr)
	defer l.Close()
	go Accept(l)
	time.Sleep(10 * time.Minute)
}
