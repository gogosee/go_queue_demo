package main

import (
	"fmt"
	"go_queue/core"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Print("connect failed, err:", err)
	}
	defer conn.Close()

	msg := core.Msg{Id: 1102, Topic: "topic-test", MsgType: 2, Payload: []byte("我")}
	n, err := conn.Write(core.MsgToBytes(msg))
	if err != nil {
		fmt.Print("write failed, err:", err)
	}

	fmt.Print(n)
}
