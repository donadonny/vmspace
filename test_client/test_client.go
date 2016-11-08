package main

import (
	"fmt"
	"time"

	"github.com/panyingyun/vmspace/gateway"
	"github.com/panyingyun/vmspace/node"
	"github.com/smallnest/rpcx"
)

func main() {
	selector := &rpcx.DirectClientSelector{
		Network:     "tcp",
		Address:     "127.0.0.1:8972",
		DialTimeout: 10 * time.Second,
	}
	client := rpcx.NewClient(selector)
	defer client.Close()

	args1 := &gateway.GWSendArgs{
		Gwid:    "1122334455667788",
		Lng:     120.3,
		Lat:     30.5,
		Payload: nil,
	}

	var reply1 gateway.GWSendReply
	client.Call("GW.Send", args1, &reply1)
	fmt.Println("reply1 = ", reply1)

	args2 := &node.SendArgs{
		Deveui:  "aabbccddeeff0000",
		Lng:     120.3,
		Lat:     30.5,
		Payload: nil,
	}

	var reply2 node.SendReply
	client.Call("Node.Send", args2, &reply2)
	fmt.Println("reply2 = ", reply2)

	args3 := &node.ReceiveArgs{
		Deveui: "aabbccddeeff0000",
	}
	var reply3 node.ReceiveReply
	client.Call("Node.Receive", args3, &reply3)
	fmt.Println("reply3 = ", len(reply3.Payload))

	args4 := &gateway.GWReceiveArgs{
		Gwid: "1122334455667788",
	}
	var reply4 gateway.GWReceiveReply
	client.Call("GW.Receive", args4, &reply4)
	fmt.Println("reply4 = ", len(reply4.Payload))
}
