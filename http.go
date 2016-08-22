package main

import (
	"io"
	"fmt"
	"net/http"
	"strings"
)

func www_root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html",)
	var searchword string = ""
	var tool string = ""
	var project string = ""
	var slug string = ""
	var urllist []string
	searchword = r.FormValue("searchword")
	urllist = strings.Split(r.URL.Path, "/")

	if len(urllist) == 5 {
		slug = urllist[4]
		project = urllist[3]
		tool = urllist[2]
		logs, err := findtpsDB(tool, project, slug)
		if err != nil {
			io.WriteString(w, headHTML + "<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML + infoHTML(tool,project,slug) + searchboxHTML(searchword)+logHTML(logs))
		return
	} else if len(urllist) == 4 {
		project = urllist[3]
		tool = urllist[2]
		logs, err := findtpDB(tool, project)
		if err != nil {
			io.WriteString(w, headHTML + "<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML + infoHTML(tool, project,"") + searchboxHTML(searchword) + logHTML(logs))
		return
	} else if len(urllist) == 3 {
		tool = urllist[2]
		logs, err := findtDB(tool)
		if err != nil {
			io.WriteString(w, headHTML + "<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
			return
		}
		io.WriteString(w, headHTML + infoHTML(tool,"","") + searchboxHTML(searchword) + logHTML(logs))
		return
	}  else {
		if searchword != "" {
			logs, err := findDB(searchword)
			if err != nil {
				io.WriteString(w, headHTML + "<br><center>DB 또는 네트워크 장애로 로그를 가지고 올 수 없습니다.</center>")
				return
			}
			io.WriteString(w, headHTML + infoHTML("","","") + searchboxHTML(searchword) + logHTML(logs))
			return
		}
	}
}

func webserver(port string) {
	http.HandleFunc("/", www_root)
	fmt.Printf("Web Server Start : http://%s%s\n", LocalIP(), port)
	http.ListenAndServe(port, nil)
}
