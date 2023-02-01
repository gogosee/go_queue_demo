package main

import (
	"fmt"
	"go_queue/core"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Print("listen failed, err:", err)
		return
	}
	go core.Save()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("accept failed, err:", err)
			continue
		}
		go core.Process(conn)

	}
}
