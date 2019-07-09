package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"time"

	"github.com/digital-idea/dilog"
)

var (
	// DBIP 값은 컴파일 단계에서 회사에 따라 값이 바뀐다.
	DBIP = "127.0.0.1"

	// server setting
	regexpPort       = regexp.MustCompile(`:\d{2,5}$`)
	regexpID         = regexp.MustCompile(`\d{13}$`)
	flagHTTP         = flag.String("http", "", "dilog service port ex):8080")
	flagDBIP         = flag.String("dbip", DBIP, "MongoDB IP")
	flagPagenum      = flag.Int("pagenum", 10, "Number of items on page")
	flagProtocolPath = flag.String("protocolpath", "/show,/lustre,/project,/storage", "A path-aware string to associate with the protocol(dilink). Separate each character with a comma.")
	// add mode
	now            = time.Now()
	regexpFullTime = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T\d{2}:\d{2}:\d{2}[-+]\d{2}:\d{2}$`) // 2016-10-19T16:41:24+09:00

	flagTool    = flag.String("tool", "", "tool name")
	flagProject = flag.String("project", "", "project name")
	flagSlug    = flag.String("slug", "", "shot or asset name")
	flagLog     = flag.String("log", "", "log strings")
	flagKeep    = flag.Int("keep", 180, "Days to keep")
	flagUser    = flag.String("user", "", "custom Username.")
	flagAddTime = flag.String("addtime", now.Format(time.RFC3339), "log add time.")
	// find mode
	flagFind = flag.String("find", "", "search word")
	// remove mode
	flagRm   = flag.Bool("rm", false, "Delete data older than keep days")
	flagRmID = flag.String("rmid", "", "ID number to dalete")
	// debug mode
	flagDebug = flag.Bool("debug", false, "dilog debug mode")
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
	log.SetPrefix("dilog: ")
	flag.Parse()

	// webserver
	if regexpPort.MatchString(*flagHTTP) {
		Webserver()
	}

	// remove mode
	if *flagRm {
		itemlist, err := dilog.All(*flagDBIP)
		if err != nil {
			log.Fatal(err)
		}
		if *flagDebug {
			fmt.Println(itemlist)
		}

		for _, i := range itemlist {
			isDelete, err := dilog.Timecheck(i.Time, i.Keep)
			if err != nil {
				log.Fatal(err)
			}
			if *flagDebug {
				fmt.Println(isDelete)
				fmt.Println(i.ID)
			}

			if isDelete {
				err := dilog.Remove(*flagDBIP, i.ID)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		return
	}
	// remove id mode
	if regexpID.MatchString(*flagRmID) {
		err := dilog.Remove(*flagDBIP, *flagRmID)
		if err != nil {
			log.Fatal(err)
		}
	}
	// add mode
	if *flagTool != "" && *flagLog != "" {
		if !regexpFullTime.MatchString(*flagAddTime) {
			log.Fatal("RFC3339 시간형식이 아닙니다.")
		}
		if *flagUser == "" {
			*flagUser = username()
		}
		ip, err := serviceIP()
		if err != nil {
			log.Fatal(err)
		}
		err = dilog.Add(*flagDBIP, ip, *flagLog, *flagProject, *flagSlug, *flagTool, *flagUser, *flagAddTime, *flagKeep)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	flag.PrintDefaults()
}
