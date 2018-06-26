package main

import (
	"regexp"
)

var regexpPort = regexp.MustCompile(`:\d{2,5}$`)
