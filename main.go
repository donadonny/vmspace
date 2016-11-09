package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/panyingyun/vmspace/gateway"
	"github.com/panyingyun/vmspace/node"
	"github.com/smallnest/rpcx"
)

func main() {
	server := rpcx.NewServer()
	nodemgr := node.NewNodeMgnager()
	server.RegisterName("Node", nodemgr)
	gwmgr := gateway.NewGWMgnager()
	server.RegisterName("GW", gwmgr)

	//Start RPC Server
	fmt.Println("Start rpcx server...")
	err := server.Start("tcp", "127.0.0.1:8972")
	if err != nil {
		fmt.Println("Start rpcx fail!")
	}
	fmt.Println("Start rpcx server OK!")

	//Start Exchange Server
	go func() {
		for {
			downlink := gwmgr.GetDownlinkPayload()
			if downlink != nil {
				nodemgr.SetDownlinkPayload(downlink)
			}
			time.Sleep(time.Second)

			uplink := nodemgr.GetUplinkPayload()
			if uplink != nil {
				gwmgr.SetUplinkPayload(uplink)
			}
			time.Sleep(time.Second)
		}

	}()
	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("signal received signal %v\n", <-sigChan)
	fmt.Println("shutting down server")
}
