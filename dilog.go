package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
)

var (
	regexpPort = regexp.MustCompile(`:\d{2,5}$`)
	flagHTTP   = flag.String("http", ":8080", "dilog service port ex):8080")
	flagDBIP   = flag.String("dbip", "127.0.0.1", "mongodb ip")
	// add mode
	flagTool    = flag.String("tool", "", "tool name")
	flagProject = flag.String("project", "", "project name")
	flagSlug    = flag.String("slug", "", "shot or asset name")
	flagLog     = flag.String("log", "", "log strings")
	flagKeep    = flag.String("keep", "180", "Days to keep")
	flagUser    = flag.String("user", "", "custom Username.")
	// find mode
	flagFind = flag.String("find", "", "search word")
	// remove mode
	flagRm = flag.Bool("rm", false, "Delete data older than keep days")
)

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
	flag.Parse()

	if *flagTool != "" && *flagLog != "" && *flagHTTP == "" {
		if *flagUser == "" {
			*flagUser = username()
		}

		// MPAA를 대비하기 위해선, 로그기록시 IP, Port가 DB에 저장되는 것이 좋다.
		// Port는 restAPI 퀘스트 헤더에 존재하는데, 현재 많은 툴에서 사용중인 dilog의 cmd모드는 제거하고 restAPI만 남기면 좋다.
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		err = addDB(ip, *flagKeep, *flagLog, *flagProject, *flagSlug, *flagTool, *flagUser)
		if err != nil {
			log.Fatal("DB장애로 로그를 추가할 수 없습니다.")
		}
	} else if *flagFind != "" {
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
		items, err := findDB(*flagFind)
		if err != nil {
			log.Fatal("DB장애로 처리할 수 없습니다.")
		}
		for _, i := range items {
			fmt.Printf("%-25s %-4s %-15s %-20s %-10s %-10s %-14s %s\n", i.Time, i.Keep, i.Cip, i.User, i.Tool, i.Project, i.Slug, i.Log)
		}
	} else if *flagRm {
		//remove mode
		itemlist, err := allDB()
		if err != nil {
			log.Fatal("DB장애로 처리할 수 없습니다.")
		}
		for _, i := range itemlist {
			if timecheck(i.Time, i.Keep) {
				rmbool, err := rmDB(i.ID)
				if err != nil {
					log.Fatal(err)
				}
				if rmbool {
					return
				}
				fmt.Println("해당 slug를 삭제할 수 없습니다.")
				return
			}
		}
	} else if regexpPort.MatchString(*flagHTTP) {
		webserver()
	}
	fmt.Println("Digitalidea log server")
	flag.PrintDefaults()
}
