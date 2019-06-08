package main

import (
	"fmt"
	"strings"
	"text/template"
)

//템플릿 함수를 로딩합니다.
var funcMap = template.FuncMap{
	"addLink": addLink,
}

// addLink templatefunc 는 로그내용에 *flagProtocolPath에서 선언한 경로문자열이 포함된다면 dilink:// 링크를 생성시킨다.
func addLink(log string) string {
	var rstring string
	for _, word := range strings.Split(log, " ") {
		isPath := false
		for _, path := range strings.Split(*flagProtocolPath, ",") {
			if !strings.Contains(word, path) {
				continue
			}
			isPath = true
			break
		}
		if isPath {
			rstring += fmt.Sprintf(`<a href="dilink://%s">%s</a>`, word, word) + " "
		} else {
			rstring += word + " "
		}
	}
	return rstring
}
