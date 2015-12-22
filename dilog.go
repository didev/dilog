package main

import (
	"flag"
	"time"
	"fmt"
	"os/user"
	"runtime"
	"os"
	"strings"
	)

const DBIP = "10.0.90.251"

func timenow() string {
	t := time.Now()
	return fmt.Sprintf("%d/%02d/%02d-%02d:%02d:%02d",
			t.Year(), t.Month(), t.Day(),
			t.Hour(), t.Minute(), t.Second())
}

func username() string {
	user, err := user.Current()
	if err != nil {
		if runtime.GOOS == "darwin" {
			return os.ExpandEnv("$USER")
		} else if runtime.GOOS == "linux" {
			return os.ExpandEnv("$USER")
		} else {
			return user.Username
		}
	}
	return user.Username
}

func main() {
	var localip string
	var user string
	toolPtr := flag.String("tool", "", "add mode. | tool name (indispensable)")
	projectPtr := flag.String("project", "", "add mode. | project name")
	slugPtr := flag.String("slug", "", "add mode. | shot or asset name")
	logPtr := flag.String("log", "", "add mode. | log strings (indispensable)")
	keepPtr := flag.String("keep", "180", "add mode. | be kept")
	findPtr := flag.String("find", "", "find mode. | search word")
	findnumPtr := flag.String("findnum", "", "find mode. | return search count")
	rmPtr := flag.String("rm", "", "remove mode. | remove over kept")
	httpPtr := flag.String("http", "", "server mode. | service port ex):8080")
	ipPtr := flag.String("ip", "", "local ip.")
	userPtr := flag.String("user", "", "custom Username.")

	flag.Parse()

	if *toolPtr != "" && *logPtr != "" && *httpPtr == ""{
		if *ipPtr == "" {
			localip = LocalIP()
		} else {
			localip = *ipPtr
		}
		if *userPtr == "" {
			user = username()
		} else {
			user = *userPtr
		}
		addDB(localip, *keepPtr, *logPtr, *projectPtr, DBIP, *slugPtr, timenow(), *toolPtr, user)
	} else if *findPtr != "" {
		//find mode
		fmt.Printf("%-19s %-04s %-15s %-20s %-10s %-10s %-14s %s\n","Time", "Keep", "IP","User","Tool","Project","Slug","Log")
		fmt.Printf("%-19s %-04s %-15s %-20s %-10s %-10s %-14s %s\n",
					strings.Repeat("-",19),
					strings.Repeat("-", 4),
					strings.Repeat("-",15),
					strings.Repeat("-",20),
					strings.Repeat("-",10),
					strings.Repeat("-",10),
					strings.Repeat("-",14),
					strings.Repeat("-",20),
				)
		for _,i := range findDB(*findPtr) {
			fmt.Printf("%-19s %-4s %-15s %-20s %-10s %-10s %-14s %s\n", i.Time, i.Keep, i.Cip, i.User, i.Tool, i.Project, i.Slug, i.Log)
		}
	} else if *findnumPtr != "" {
		fmt.Printf("%d\n", findnumDB(*findnumPtr))

	} else if *rmPtr == "temp" {
		//remove mode
		for _, i := range allDB() {
			if timecheck(i.Time, i.Keep) {
				rmDB(i.Id)
			}
		}
	} else if *httpPtr != "" {
		//web server mode.
		webserver(*httpPtr)
	} else {
		fmt.Println("Digitalidea log utility")
		flag.PrintDefaults()
	}
}
