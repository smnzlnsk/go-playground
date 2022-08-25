package main

import (
	"fmt"

	"github.com/smnzlnsk/go-utils/utils"

	//	"golang.org/x/net/ipv6"
	"net"
)

func main() {
	var i, j = 1, 2
	k := utils.Add(i, j)
	fmt.Println("Result:", k)

	var strc utils.Nic
	fmt.Println("Running playground executable:")
	var n, _ = net.Interfaces()
	for _, num := range n {
		fmt.Println(num)
	}
	strc.Mac = n[1].HardwareAddr
	strc.Mtu = n[1].MTU
	strc.Name = n[1].Name
	fmt.Println(strc)
}
