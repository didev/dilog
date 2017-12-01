package main

import (
	"net"
	"os"
	"runtime"
	"strings"
)

func LocalIP() string {
	hw, err := net.InterfaceAddrs()
	var ip string
	ip = "0.0.0.0"
	if err != nil {
		os.Exit(1)
	}
	for _, i := range hw {
		//linux and mac
		if runtime.GOOS == "windows" && i.String() != "127.0.0.1" {
			ip = i.String()
		} else {
			if strings.Contains(i.String(), "/16") {
				ip = i.String()[0 : len(i.String())-3]
			}
		}
	}
	return ip
}
