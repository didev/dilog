package main

import (
	"html/template"
	"strings"
)

//템플릿 함수를 로딩합니다.
var funcMap = template.FuncMap{
	"getPath": getPath,
}

// getPath templatefunc 는 로그내용에 경로가 존재하면 경로문자만 반환한다. 경로가 없다면 빈문자열을 반환한다.
func getPath(log string) string {
	for _, word := range strings.Split(log, " ") {
		for _, path := range strings.Split(*flagProtocolPath, ",") {
			if !strings.Contains(word, path) {
				continue
			}
			return word
		}
	}
	return ""
}
