package main

import (
	"fmt"

	"github.com/smnzlnsk/go-utils/utils"

	//	"golang.org/x/net/ipv6"
	"net"
)

func main() {
	k := utils.Add(1, 2)
	fmt.Println("Result:", k)

	fmt.Println("Running playground executable:")
	var n, _ = net.Interfaces()
	for _, num := range n {
		fmt.Println(num)
	}
	var addr, _ = n[0].Addrs()
	fmt.Println("IP:", addr)
}
