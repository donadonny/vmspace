package node

import (
	"fmt"
)

type SendArgs struct {
	Deveui  string  `msg:"deveui"`
	Lng     float64 `msg:"lng"`
	Lat     float64 `msg:"lat"`
	Payload []byte  `msg:"payload"`
}

type SendReply struct {
	Code int `msg:"code"`
}

type ReceiveArgs struct {
	Deveui string `msg:"deveui"`
}

type ReceiveReply struct {
	Payload []byte `msg:"payload"`
}

type Node struct {
	deveui string
	lng    float64
	lat    float64
}

type NodeManager struct {
	nodes    map[string]Node
	uplink   map[string][]byte
	downlink map[string][]byte
}

// NewBackend creates a new Backend.
func NewNodeMgnager() *NodeManager {
	b := NodeManager{
		nodes:    make(map[string]Node),
		uplink:   make(map[string][]byte),
		downlink: make(map[string][]byte),
	}
	return &b
}

func (m *NodeManager) Send(args *SendArgs, reply *SendReply) error {
	fmt.Printf("Send Node[%v] Data here!!\n", args.Deveui)
	reply.Code = 200
	// node has or not
	if n, ok := m.nodes[args.Deveui]; ok {
		n.deveui = args.Deveui
		n.lat = args.Lat
		n.lng = args.Lng
		m.uplink[args.Deveui] = args.Payload
	} else {
		var newnode Node
		newnode.deveui = args.Deveui
		newnode.lat = args.Lat
		newnode.lng = args.Lng
		m.nodes[args.Deveui] = newnode
		m.uplink[args.Deveui] = args.Payload
	}
	return nil
}

func (m *NodeManager) Receive(args *ReceiveArgs, reply *ReceiveReply) error {
	fmt.Printf("Receive Node[%v] Data here!!\n", args.Deveui)
	reply.Payload = nil
	if down, ok := m.downlink[args.Deveui]; ok {
		reply.Payload = down
		delete(m.downlink, args.Deveui)
	}
	return nil
}
