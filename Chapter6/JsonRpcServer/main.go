package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Op struct {
}

func (this *Op) Add(i int, pI *int) error {
	*pI = i * i
	return nil
}

func main() {
	server := rpc.NewServer()
	server.Register(&Op{})
	l, _ := net.Listen("tcp", ":1234")
	for {
		conn, _ := l.Accept()
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
