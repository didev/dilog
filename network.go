package main

import (
	"os"
	"net"
	"strings"
	"runtime"
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
		if runtime.GOOS == "windows" && i.String() != "127.0.0.1"{
			ip = i.String()
		} else {
			if strings.Contains(i.String(), "/16") {
				ip = i.String()[0:len(i.String())-3]
			}
		}
	}
	return ip
}

/*
func LocalIP() string {
	var iplist []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString(err.Error()+"\n")
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				iplist = append(iplist, ipnet.IP.String())
			}
		}
	}
	return iplist[0]
}
*/
