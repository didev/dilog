package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	type recipe struct {
		Searchword string
		Tool       string
		Project    string
		Slug       string
		Logs       []Log
	}
	rcp := recipe{}
	templates.ExecuteTemplate(w, "index", rcp)
}

func search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	q := r.URL.Query()
	type recipe struct {
		Searchword string
		Tool       string
		Project    string
		Slug       string
		Logs       []Log
	}
	rcp := recipe{}
	rcp.Searchword = q.Get("searchword")
	rcp.Tool = q.Get("tool")
	rcp.Project = q.Get("project")
	rcp.Slug = q.Get("slug")

	rcp.Searchword = r.FormValue("searchword")

	if rcp.Tool != "" && rcp.Project != "" && rcp.Slug != "" {
		logs, err := findtpsDB(rcp.Tool, rcp.Project, rcp.Slug)
		if err != nil {
			templates.ExecuteTemplate(w, "dberr", nil)
			return
		}
		rcp.Logs = logs
		templates.ExecuteTemplate(w, "result", rcp)
		return
	}
	if rcp.Tool != "" && rcp.Project != "" {
		logs, err := findtpDB(rcp.Tool, rcp.Project)
		if err != nil {
			templates.ExecuteTemplate(w, "dberr", nil)
			return
		}
		rcp.Logs = logs
		templates.ExecuteTemplate(w, "result", rcp)
		return
	}
	if rcp.Tool != "" {
		logs, err := findtDB(rcp.Tool)
		if err != nil {
			templates.ExecuteTemplate(w, "dberr", nil)
			return
		}
		rcp.Logs = logs
		templates.ExecuteTemplate(w, "result", rcp)
		return
	}
	if rcp.Searchword != "" {
		logs, err := findDB(rcp.Searchword)
		if err != nil {
			templates.ExecuteTemplate(w, "dberr", nil)
			return
		}
		rcp.Logs = logs
		templates.ExecuteTemplate(w, "result", rcp)
		return
	}
	templates.ExecuteTemplate(w, "result", rcp)
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
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("assets/img"))))
	http.HandleFunc("/search", search)
	http.HandleFunc("/", index)
	http.HandleFunc("/api/setlog", handleAPISetLog)
	fmt.Printf("Web Server Start : http://%s%s\n", ip, *flagHTTP)
	err = http.ListenAndServe(*flagHTTP, nil)
	if err != nil {
		log.Fatal(err)
	}
}
