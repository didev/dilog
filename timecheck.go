package main

import (
	"time"
	"log"
	)


func str2time(str string) time.Time {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		log.Println("시간을 바꿀 수 없습니다.")
	}
	return t
}

func timecheck(timestr, keepdate string) bool {
	recordtime := str2time(timestr)
	alltime, _ := time.ParseDuration(keepdate + "h")
	addtime := recordtime.Add(alltime * 24)
	now := time.Now()
	return now.After(addtime) //추후 이 결과를 이용해서 참이면 리무브 대상이다.
}
