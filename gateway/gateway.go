package gateway

import (
	"fmt"
)

type GWSendArgs struct {
	Gwid    string  `msg:"gwid"`
	Lng     float64 `msg:"lng"`
	Lat     float64 `msg:"lat"`
	Payload []byte  `msg:"payload"`
}

type GWSendReply struct {
	Code int `msg:"code"`
}

type GWReceiveArgs struct {
	Gwid string `msg:"gwid"`
}

type GWReceiveReply struct {
	Payload []byte `msg:"payload"`
}

type Gateway struct {
	gwid string
	lng  float64
	lat  float64
}

type GatewayManager struct {
	gws      map[string]Gateway
	uplink   map[string][]byte
	downlink map[string][]byte
}

// NewBackend creates a new Backend.
func NewGWMgnager() *GatewayManager {
	b := GatewayManager{
		gws:      make(map[string]Gateway),
		uplink:   make(map[string][]byte),
		downlink: make(map[string][]byte),
	}
	return &b
}

func (m *GatewayManager) Send(args *GWSendArgs, reply *GWSendReply) error {
	fmt.Printf("Send GW [%v] Data here!!\n", args.Gwid)
	reply.Code = 200
	// node has or not
	if n, ok := m.gws[args.Gwid]; ok {
		n.gwid = args.Gwid
		n.lat = args.Lat
		n.lng = args.Lng
		if args.Payload != nil {
			m.downlink[args.Gwid] = args.Payload
		}
	} else {
		var newgw Gateway
		newgw.gwid = args.Gwid
		newgw.lat = args.Lat
		newgw.lng = args.Lng
		m.gws[args.Gwid] = newgw
		if args.Payload != nil {
			m.downlink[args.Gwid] = args.Payload
		}
	}
	fmt.Println("m = ", m)
	return nil
}

func (m *GatewayManager) Receive(args *GWReceiveArgs, reply *GWReceiveReply) error {
	fmt.Printf("Receive GW [%v] Data here!!\n", args.Gwid)
	reply.Payload = nil
	if up, ok := m.uplink[args.Gwid]; ok {
		reply.Payload = up
		delete(m.uplink, args.Gwid)
	}
	fmt.Println("m = ", m)
	return nil
}
