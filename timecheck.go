package main

import (
	"time"
	"strconv"
	"strings"
	)


func str2time(intime string) time.Time {
	var month time.Month
	info := strings.Split(intime, "-")
	datelist := strings.Split(info[0], "/")
	timelist := strings.Split(info[1], ":")
	year,_ := strconv.Atoi(datelist[0])
	monthint,_ := strconv.Atoi(datelist[1])
	switch monthint {
		case 1: month = 1
		case 2: month = 2
		case 3: month = 3
		case 4: month = 4
		case 5: month = 5
		case 6: month = 6
		case 7: month = 7
		case 8: month = 8
		case 9: month = 9
		case 10: month = 10
		case 11: month = 11
		case 12: month = 12
	}
	day,_ := strconv.Atoi(datelist[2])
	hour,_ := strconv.Atoi(timelist[0])
	min,_  := strconv.Atoi(timelist[1])
	sec,_  := strconv.Atoi(timelist[2])
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

func str2time_utc(intime string) time.Time {
	var month time.Month
	year,_ := strconv.Atoi(strings.Split(intime, "-")[0])
	monthint,_ := strconv.Atoi(strings.Split(intime, "-")[1])
	switch monthint {
		case 1: month = 1
		case 2: month = 2
		case 3: month = 3
		case 4: month = 4
		case 5: month = 5
		case 6: month = 6
		case 7: month = 7
		case 8: month = 8
		case 9: month = 9
		case 10: month = 10
		case 11: month = 11
		case 12: month = 12
	}
	day,_ := strconv.Atoi(strings.Split(strings.Split(intime, " ")[0], "-")[2])
	hour,_ := strconv.Atoi(strings.Split(strings.Split(intime, " ")[1], ":")[0])
	min,_  := strconv.Atoi(strings.Split(strings.Split(intime, " ")[1], ":")[1])
	return time.Date(year, month, day, hour, min, 0, 0, time.Local)
}

func timecheck(timestr, keepdate string) bool {
	if strings.Contains(timestr, "UTC") {
		recordtime := str2time_utc(timestr)
		alltime, _ := time.ParseDuration(keepdate + "h")
		addtime := recordtime.Add(alltime * 24)
		now := time.Now()
		return now.After(addtime) //if true then remove
	} else {
		recordtime := str2time(timestr)
		alltime, _ := time.ParseDuration(keepdate + "h")
		addtime := recordtime.Add(alltime * 24)
		now := time.Now()
		return now.After(addtime) //if true then remove
	}
}
