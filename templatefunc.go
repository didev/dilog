package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/template"
)

//템플릿 함수를 로딩합니다.
var funcMap = template.FuncMap{
	"addLink":  addLink,
	"num2list": num2pagelist,
}

// addLink templatefunc 는 로그에 경로가 있다면 dilink:// 를 생성시킨다.
func addLink(log string) string {
	var rstring string
	words := strings.Split(log, " ")
	for _, word := range words {
		if strings.Contains(word, "/show") || strings.Contains(word, "/lustre") {
			rstring += fmt.Sprintf(`<a href="dilink://%s">%s</a>`, word, word) + " "
		} else {
			rstring += word + " "
		}
	}
	return rstring
}

func num2pagelist(num int) []string {
	var page []string
	for i := 1; i < num+1; i++ {
		page = append(page, strconv.Itoa(i))
	}
	return page
}
