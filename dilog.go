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
	"text/template"
)

var (
	// server setting
	regexpPort         = regexp.MustCompile(`:\d{2,5}$`)
	flagHTTP           = flag.String("http", "", "dilog service port ex):8080")
	flagDBIP           = flag.String("dbip", "127.0.0.1", "mongodb ip")
	flagDBName         = flag.String("dbname", "dilog", "mongodb database name")
	flagCollectionName = flag.String("collection", "logs", "mongodb database name")
	templates          = template.New("main")
	// add mode
	flagTool    = flag.String("tool", "", "tool name")
	flagProject = flag.String("project", "", "project name")
	flagSlug    = flag.String("slug", "", "shot or asset name")
	flagLog     = flag.String("log", "", "log strings")
	flagKeep    = flag.Int("keep", 180, "Days to keep")
	flagUser    = flag.String("user", "", "custom Username.")
	// find mode
	flagFind = flag.String("find", "", "search word")
	// remove mode
	flagRm = flag.Bool("rm", false, "Delete data older than keep days")
	// flag help
	flagHelp = flag.Bool("help", false, "print help")
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
	// webserver
	if regexpPort.MatchString(*flagHTTP) {
		templates = template.Must(template.New("main").ParseGlob("assets/template/*"))
		Webserver()
	}
	//find mode
	if *flagFind != "" {
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
			fmt.Printf("%-25s %-4d %-15s %-20s %-10s %-10s %-14s %s\n", i.Time, i.Keep, i.Cip, i.User, i.Tool, i.Project, i.Slug, i.Log)
		}
		return
	}
	//remove mode
	if *flagRm {
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
		return
	}
	// add mode
	if *flagTool != "" && *flagLog != "" {
		if *flagUser == "" {
			*flagUser = username()
		}
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		err = addDB(ip, *flagLog, *flagProject, *flagSlug, *flagTool, *flagUser, *flagKeep)
		if err != nil {
			log.Fatal("DB장애로 로그를 추가할 수 없습니다.")
		}
		return
	}
	fmt.Println("Digitalidea log server")
	flag.PrintDefaults()
}
