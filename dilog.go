package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"runtime"
	"strings"
)

const DBIP = "10.0.90.251"

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
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

	if *toolPtr != "" && *logPtr != "" && *httpPtr == "" {
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
		// MPAA를 대비하기 위해선, 로그기록시 IP, Port가 DB에 저장되는 것이 좋다.
		// Port는 restAPI 퀘스트 헤더에 존재하는데, 현재 많은 툴에서 사용중인 dilog의 cmd모드는 제거하고 restAPI만 남기면 좋다.

		port := "" // 기존 cmd방식을 유지하기 위해서 빈문자열을 설정함.
		err := addDB(localip, port, *keepPtr, *logPtr, *projectPtr, *slugPtr, *toolPtr, user)
		if err != nil {
			log.Fatal("DB장애로 로그를 추가할 수 없습니다.")
		}
	} else if *findPtr != "" {
		//find mode
		fmt.Printf("%-25s %-04s %-15s %-20s %-10s %-10s %-14s %s\n", "Time", "Keep", "IP", "User", "Tool", "Project", "Slug", "Log")
		fmt.Printf("%-25s %-04s %-15s %-20s %-10s %-10s %-14s %s\n",
			strings.Repeat("-", 25),
			strings.Repeat("-", 4),
			strings.Repeat("-", 15),
			strings.Repeat("-", 20),
			strings.Repeat("-", 10),
			strings.Repeat("-", 10),
			strings.Repeat("-", 14),
			strings.Repeat("-", 20),
		)
		items, err := findDB(*findPtr)
		if err != nil {
			log.Fatal("DB장애로 처리할 수 없습니다.")
		}
		for _, i := range items {
			fmt.Printf("%-25s %-4s %-15s %-20s %-10s %-10s %-14s %s\n", i.Time, i.Keep, i.Cip, i.User, i.Tool, i.Project, i.Slug, i.Log)
		}
	} else if *findnumPtr != "" {
		num, err := findnumDB(*findnumPtr)
		if err != nil {
			log.Fatal("DB장애로 처리할 수 없습니다.")
		}
		fmt.Printf("%d\n", num)

	} else if *rmPtr == "temp" {
		//remove mode
		itemlist, err := allDB()
		if err != nil {
			log.Fatal("DB장애로 처리할 수 없습니다.")
		}
		for _, i := range itemlist {
			if timecheck(i.Time, i.Keep) {
				rmbool, err := rmDB(i.Id)
				if err != nil {
					log.Fatal(err)
				}
				if rmbool {
					return
				} else {
					fmt.Println("해당 slug를 삭제할 수 없습니다.")
				}
			}
		}
	} else if regexpPort.MatchString(*httpPtr) {
		//web server mode.
		webserver(*httpPtr)
	} else {
		fmt.Println("digitalidea log server")
		flag.PrintDefaults()
	}
}
