package main

import (
	"io"
	"fmt"
	//"net"
	"net/http"
	"strings"
)

func www_root(w http.ResponseWriter, r *http.Request) {
	var searchword string = ""
	var tool string = ""
	var project string = ""
	var slug string = ""
	var urllist []string
	var logs []Log
	w.Header().Set("Content-Type", "text/html",)
	searchword = r.FormValue("searchword")
	urllist = strings.Split(r.URL.Path, "/")

	if len(urllist) == 5 {
		slug = urllist[4]
		project = urllist[3]
		tool = urllist[2]
		logs = findtpsDB(tool, project, slug)
		io.WriteString(w, headHTML + infoHTML(tool,project,slug) + searchboxHTML(searchword)+logHTML(logs))
	} else if len(urllist) == 4 {
		project = urllist[3]
		tool = urllist[2]
		logs = findtpDB(tool, project)
		io.WriteString(w, headHTML + infoHTML(tool, project,"") + searchboxHTML(searchword) + logHTML(logs))
	} else if len(urllist) == 3 {
		tool = urllist[2]
		logs = findtDB(tool)
		io.WriteString(w, headHTML + infoHTML(tool,"","") + searchboxHTML(searchword) + logHTML(logs))
	}  else {
		if searchword != "" {
			logs = findDB(searchword)
		}
		io.WriteString(w, headHTML + infoHTML("","","") + searchboxHTML(searchword) + logHTML(logs))
	}
}

func webserver(port string) {
	http.HandleFunc("/", www_root)
	fmt.Printf("Web Server Start : http://%s%s\n", LocalIP(), port)
	http.ListenAndServe(port, nil)
}
