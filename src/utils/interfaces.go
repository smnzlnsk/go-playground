package utils

import (
	"fmt"
	"net"
)

type itrfc struct {
	name string
	mtu  int
	mac  net.HardwareAddr
}

func printInterface(i itrfc) {
	fmt.Println(name, mtu, mac)
}
