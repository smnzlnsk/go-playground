package main

import (
	"fmt"
	"github.com/smnzlnsk/utils"
	"golang.org/x/net/ipv6"
	"net"
)

func main() {
	var strc itrfc
	fmt.Println("Running playground executable:")
	var n, _ = net.Interfaces()
	for _, num := range n {
		fmt.Println(num)
	}
	strc.mac = n[1].HardwareAddr
	strc.mtu = n[1].MTU
	strc.name = n[1].Name
	fmt.Println(strc)
}
