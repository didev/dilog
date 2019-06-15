package main

import (
	"fmt"
	"html/template"
	"strings"
)

//템플릿 함수를 로딩합니다.
var funcMap = template.FuncMap{
	"addLink": addLink,
	"hasPath": hasPath,
	"getPath": getPath,
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

// hasPath templatefunc 는 로그내용에 경로가 존재하는지 체크한다.
func hasPath(log string) bool {
	for _, path := range strings.Split(*flagProtocolPath, ",") {
		if !strings.Contains(log, path) {
			continue
		}
		return true
	}
	return false
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
