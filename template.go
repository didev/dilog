package main

import (
	"fmt"
	"strings"
)

// template func가 되면 좋다.
func linklog(log string) string {
	var rstring string
	rstring = ""
	loglist := strings.Split(log, " ")
	for _, i := range loglist {
		if strings.Contains(i, "/show") || strings.Contains(i, "/lustre") {
			rstring = rstring + fmt.Sprintf(`<a href="dilink://%s">%s</a>`, i, i) + " "
		} else {
			rstring = rstring + i + " "
		}
	}
	return rstring
}
