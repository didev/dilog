package main

import (
	"net"
	"os"
	"runtime"
	"strings"
)

// LocalIP 함수는 로컬아이피를 가지고 온다.
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
