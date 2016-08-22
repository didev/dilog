package main

import (
	"os"
	"time"
	"strconv"
	"strings"
	"runtime"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Log struct {
	Cip string
	Id string
	Keep string
	Log string
	Os string
	Project string
	Sip string
	Slug string
	Time string
	Tool string
	User string
}

func genid() string {
	return strconv.Itoa(int(time.Now().UnixNano() / int64(time.Millisecond)))
}

func addDB(cip, keep, logstr, project, sip, slug, time, tool, user string) error {
	session, err := mgo.Dial(DBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dilog").C("log")
	doc := Log{ Cip: cip,
				Id: genid(),
				Keep: keep,
				Log: logstr,
				Os: runtime.GOOS,
				Project: project,
				Sip: sip,
				Slug: slug,
				Time: time,
				Tool: tool,
				User: user,
				}
	err = c.Insert(doc)
	if err != nil {
		log.Println("DB Insert Err : ", err)
		return err
	}
	return nil
}

func allDB() ([]Log, error) {
	var results []Log
	session, err := mgo.Dial(DBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return results, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("dilog").C("log")
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return results, err
	}
	return results, nil
}


func findtDB(toolname string) []Log {
	var results []Log
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c :=  session.DB("dilog").C("log")
	c.Find(bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}}).Sort("-time").All(&results)
	return results
}

func findtpDB(toolname,project string) []Log {
	var results []Log
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c :=  session.DB("dilog").C("log")
	c.Find(bson.M{"$and": []bson.M {
			bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
			bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
		}}).Sort("-time").All(&results)
	return results
}


func findtpsDB(toolname,project,slug string) []Log {
	var results []Log
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c :=  session.DB("dilog").C("log")
	c.Find(bson.M{"$and": []bson.M {
			bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
			bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
			bson.M{"slug": &bson.RegEx{Pattern: slug, Options: "i"}},
		}}).Sort("-time").All(&results)
	return results
}


func findDB(searchword string) []Log {
	var results []Log
	var wordlist []string
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c :=  session.DB("dilog").C("log")
	if len(strings.Split(searchword, " ")) == 1 {
		c.Find(bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: searchword, Options: "i"}},
				}}).Sort("-time").All(&results)
	} else if len(strings.Split(searchword, " ")) == 2 {
		wordlist = strings.Split(searchword, " ")
		c.Find(bson.M{"$and": []bson.M {
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				}},
			},
		}).Sort("-time").All(&results)
	} else {
		wordlist = strings.Split(searchword, " ")
		c.Find(bson.M{"$and": []bson.M {
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				}},
		}}).Sort("-time").All(&results)
	}
	return results
}


func findnumDB(searchword string) int {
	var results []Log
	var wordlist []string
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c :=  session.DB("dilog").C("log")
	if len(strings.Split(searchword, " ")) == 1 {
		c.Find(bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: searchword, Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: searchword, Options: "i"}},
				}}).All(&results)
	} else if len(strings.Split(searchword, " ")) == 2 {
		wordlist = strings.Split(searchword, " ")
		c.Find(bson.M{"$and": []bson.M {
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				}},
			},
		}).All(&results)
	} else {
		wordlist = strings.Split(searchword, " ")
		c.Find(bson.M{"$and": []bson.M {
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[0], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[1], Options: "i"}},
				}},
				bson.M{"$or": []bson.M {
				bson.M{"cip": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"id":  &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"log": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"os": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"project": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"slug": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"time": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"tool": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				bson.M{"user": &bson.RegEx{Pattern: wordlist[2], Options: "i"}},
				}},
		}}).All(&results)
	}
	return len(results)
}

func rmDB(id string) bool {
	session, err := mgo.Dial(DBIP)
	if err != nil {
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err = session.DB("dilog").C("log").Remove(bson.M{"id":id})
	if err == nil {
		return true
	} else {
		return false
	}
}
