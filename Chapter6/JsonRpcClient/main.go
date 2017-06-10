package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := net.DialTimeout("tcp", "127.0.0.1:1234", 1000*1000*1000*30) // 30秒超时时间
	if err != nil {
		log.Fatal("client\t-", err.Error())
	}
	defer client.Close()
	clientRpc := jsonrpc.NewClient(client)
	var reply interface{}
	clientRpc.Call("Op.Add", 12, &reply)
	fmt.Println(reply)
}
