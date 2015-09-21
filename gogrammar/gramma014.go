package main

import (
	"fmt"
	"net"
)

func localmachine() {
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, inter := range interfaces {
		ads, addr_err := inter.Addrs()
		if addr_err != nil {
			fmt.Println(addr_err)
		}
		for _, v := range ads {
			ip := net.ParseIP(string(v.String()))
			ipmaks := ip.DefaultMask().String()
			fmt.Println("ip addr:", ip)
			fmt.Println("default mask:", ipmaks)
		}
		mac := inter.HardwareAddr
		fmt.Println("mac:", mac)
		fmt.Println("--------------------")
	}
}
func main() {
	localmachine()
}
