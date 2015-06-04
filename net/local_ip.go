package main

import (
	"fmt"
	"net"
)

func getLocalIp(host string) []net.IP {
	ip, err := net.LookupIP(host)
	if err != nil {
		panic(err)
	}

	return ip
}

func main() {
	addrs := getLocalIp("google.com")
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			fmt.Println("IPv4: ", ipv4)
		}
	}
}
