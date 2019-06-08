package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
)

var (
	// server setting
	regexpPort         = regexp.MustCompile(`:\d{2,5}$`)
	flagHTTP           = flag.String("http", "", "dilog service port ex):8080")
	flagDBIP           = flag.String("dbip", "127.0.0.1", "Mongodb ip")
	flagDBName         = flag.String("dbname", "dilog", "Mongodb database name")
	flagCollectionName = flag.String("collection", "logs", "Mongodb database name")
	flagPagenum        = flag.Int("pagenum", 20, "Number of items on page")
	flagProtocolPath   = flag.String("protocolpath", "/show,/lustre,/project", "A path-aware string to associate with the protocol(dilink). Separate each character with a comma.")
	tmpl               = template.New("main")
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
	// test rice, 먼저 파일을 하나 만들고 테스트 진행할 것.
	templateBox, err := rice.FindBox("assets/template")
	if err != nil {
		log.Fatal(err)
	}
	// get file contents as string
	templateString, err := templateBox.String("result.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(templateString)
	// parse and execute the template
	tmplMessage, err := template.New("message").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}
	tmplMessage.Execute(os.Stdout, map[string]string{"Searchword": "woong"})

	// webserver
	if regexpPort.MatchString(*flagHTTP) {
		tmpl = template.Must(template.New("main").Funcs(funcMap).ParseGlob("assets/template/*.html"))
		//tmpl, _ = vfstemplate.ParseGlob(assets, tmpl, "/template/*.html")

		Webserver()
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
