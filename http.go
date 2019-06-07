package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var searchword string
	var tool string
	var project string
	var slug string
	var urllist []string
	searchword = r.FormValue("searchword")
	urllist = strings.Split(r.URL.Path, "/")

	if len(urllist) == 5 {
		slug = urllist[4]
		project = urllist[3]
		tool = urllist[2]
		logs, err := findtpsDB(tool, project, slug)
		if err != nil {
			io.WriteString(w, headHTML+"<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML+infoHTML(tool, project, slug)+searchboxHTML(searchword)+logHTML(logs))
		return
	} else if len(urllist) == 4 {
		project = urllist[3]
		tool = urllist[2]
		logs, err := findtpDB(tool, project)
		if err != nil {
			io.WriteString(w, headHTML+"<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML+infoHTML(tool, project, "")+searchboxHTML(searchword)+logHTML(logs))
		return
	} else if len(urllist) == 3 {
		tool = urllist[2]
		logs, err := findtDB(tool)
		if err != nil {
			io.WriteString(w, headHTML+"<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML+infoHTML(tool, "", "")+searchboxHTML(searchword)+logHTML(logs))
		return
	} else {
		if searchword != "" {
			logs, err := findDB(searchword)
			if err != nil {
				io.WriteString(w, headHTML+"<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
				return
			}
			io.WriteString(w, headHTML+infoHTML("", "", "")+searchboxHTML(searchword)+logHTML(logs))
			return
		}
	}
	var logs []Log
	io.WriteString(w, headHTML+infoHTML("", "", "")+searchboxHTML(searchword)+logHTML(logs))
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

func webserver() {
	ip, err := serviceIP()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", root)
	http.HandleFunc("/api/setlog", handleAPISetLog)
	fmt.Printf("Web Server Start : http://%s%s\n", ip, *flagHTTP)
	http.ListenAndServe(*flagHTTP, nil)
}
