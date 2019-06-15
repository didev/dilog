package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/shurcooL/httpfs/html/vfstemplate"
)

// LoadTemplates 함수는 템플릿을 로딩합니다.
func LoadTemplates() (*template.Template, error) {
	t := template.New("").Funcs(funcMap)
	t, err := vfstemplate.ParseGlob(assets, t, "/template/*.html")
	return t, err
}

type recipe struct {
	Searchword   string
	Tool         string
	Project      string
	Slug         string
	Logs         []Log
	Page         int
	TotalPagenum []string
	Error        string
}

func num2pagelist(num int) []string {
	var page []string
	for i := 1; i < num+1; i++ {
		page = append(page, strconv.Itoa(i))
	}
	return page
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := LoadTemplates()
	if err != nil {
		log.Println("loadTemplates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	rcp := recipe{}
	err = t.ExecuteTemplate(w, "dilog", rcp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func search(w http.ResponseWriter, r *http.Request) {
	t, err := LoadTemplates()
	if err != nil {
		log.Println("loadTemplates:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	q := r.URL.Query()
	rcp := recipe{}
	rcp.Searchword = q.Get("searchword")
	rcp.Tool = q.Get("tool")
	rcp.Project = q.Get("project")
	rcp.Slug = q.Get("slug")
	page := q.Get("page")
	if page == "" {
		rcp.Page = 1
	} else {
		pagenum, err := strconv.Atoi(page)
		if err != nil {
			log.Println(err)
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "dilog", rcp)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		rcp.Page = pagenum
	}
	rcp.Searchword = r.FormValue("searchword")

	if rcp.Tool != "" && rcp.Project != "" && rcp.Slug != "" {
		logs, totalPagenum, err := findtpsDB(rcp.Tool, rcp.Project, rcp.Slug, rcp.Page)
		if err != nil {
			log.Println(err)
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "dilog", rcp)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		rcp.Logs = logs
		rcp.TotalPagenum = num2pagelist(totalPagenum)
		err = t.ExecuteTemplate(w, "dilog", rcp)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if rcp.Tool != "" && rcp.Project != "" {
		logs, totalPagenum, err := findtpDB(rcp.Tool, rcp.Project, rcp.Page)
		if err != nil {
			log.Println(err)
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "dilog", rcp)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		rcp.Logs = logs
		rcp.TotalPagenum = num2pagelist(totalPagenum)
		err = t.ExecuteTemplate(w, "dilog", rcp)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if rcp.Tool != "" {
		logs, totalPagenum, err := findtDB(rcp.Tool, rcp.Page)
		if err != nil {
			log.Println(err)
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "dilog", rcp)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		rcp.Logs = logs
		rcp.TotalPagenum = num2pagelist(totalPagenum)
		err = t.ExecuteTemplate(w, "dilog", rcp)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	if rcp.Searchword != "" {
		logs, totalPagenum, err := findDB(rcp.Searchword, rcp.Page)
		if err != nil {
			log.Println(err)
			rcp.Error = err.Error()
			err = t.ExecuteTemplate(w, "dilog", rcp)
			if err != nil {
				log.Println(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		rcp.Logs = logs
		rcp.TotalPagenum = num2pagelist(totalPagenum)
		err = t.ExecuteTemplate(w, "dilog", rcp)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	err = t.ExecuteTemplate(w, "dilog", rcp)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// PostFormValueInList 는 PostForm 쿼리시 Value값이 1개라면 값을 리턴한다.
func PostFormValueInList(key string, values []string) (string, error) {
	if len(values) != 1 {
		return "", errors.New(key + "값이 여러개 입니다.")
	}
	if values[0] == "" {
		return "", errors.New(key + "값이 빈 문자입니다.")
	}
	return values[0], nil
}

// handleApiSetLog 함수는 log를 등록하는 RestAPI이다.
func handleAPISetLog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Post Only", http.StatusMethodNotAllowed)
		return
	}
	var keep int
	var log string
	var project string
	var slug string
	var tool string
	var user string
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	r.ParseForm()
	defer r.Body.Close()
	args := r.PostForm
	for key, value := range args {
		switch key {
		case "keep":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			keep, err = strconv.Atoi(v)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
		case "log":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			log = v
		case "project":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			project = v
		case "slug":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			slug = v
		case "tool":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			tool = v
		case "user":
			v, err := PostFormValueInList(key, value)
			if err != nil {
				fmt.Fprintln(w, err)
				return
			}
			user = v
		}
	}
	err = addDB(ip, log, project, slug, tool, user, keep)
	if err != nil {
		fmt.Fprintln(w, err)
	}
}

// Webserver 함수는 웹서버를 실행합니다.
func Webserver() {
	ip, err := serviceIP()
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(assets)))
	http.HandleFunc("/search", search)
	http.HandleFunc("/", index)
	http.HandleFunc("/api/setlog", handleAPISetLog)
	fmt.Printf("Web Server Start : http://%s%s\n", ip, *flagHTTP)
	err = http.ListenAndServe(*flagHTTP, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// TotalPage 함수는 아이템의 갯수를 이용해서 총 페이지수를 반환한다.
func TotalPage(itemNum int) int {
	page := itemNum / *flagPagenum
	if itemNum%*flagPagenum != 0 {
		page++
	}
	return page
}
