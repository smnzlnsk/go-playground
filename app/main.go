package main

import (
	"fmt"

	ut "github.com/smnzlnsk/go-utils/utils"

	//"golang.org/x/net/ipv6"
	"net"
)

func main() {
	var a = ut.Add(1, 2)

	fmt.Println("Running playground executable:", a)
	var m, _ = net.Interfaces()
	for _, iface := range m {
		iiface, _ := iface.Addrs()
		fmt.Println(iface.Name, iiface)
	}
}
