package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/panyingyun/vmspace/gateway"
	"github.com/panyingyun/vmspace/node"
	"github.com/smallnest/rpcx"
)

func main() {
	server := rpcx.NewServer()
	server.RegisterName("Node", node.NewNodeMgnager())
	server.RegisterName("GW", gateway.NewGWMgnager())
	fmt.Println("Start rpcx server...")
	err := server.Start("tcp", "127.0.0.1:8972")
	if err != nil {
		fmt.Println("Start rpcx fail!")
	}
	fmt.Println("Start rpcx server OK!")
	//quit when receive end signal
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	fmt.Printf("signal received signal %v\n", <-sigChan)
	fmt.Println("shutting down server")
}
