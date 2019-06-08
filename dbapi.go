package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func addDB(ip, logstr, project, slug, tool, user string, keep int) error {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(*flagDBName).C(*flagCollectionName)
	now := time.Now()
	id := strconv.Itoa(int(now.UnixNano() / int64(time.Millisecond)))
	doc := Log{Cip: ip,
		ID:      id,
		Keep:    keep,
		Log:     logstr,
		Project: project,
		Slug:    slug,
		Time:    now.Format(time.RFC3339),
		Tool:    tool,
		User:    user,
	}
	err = c.Insert(doc)
	if err != nil {
		log.Println("DB Insert Err : ", err)
		return err
	}
	return nil
}

func allDB() ([]Log, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(*flagDBName).C(*flagCollectionName)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

func findtDB(toolname string) ([]Log, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(*flagDBName).C(*flagCollectionName)
	err = c.Find(bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}}).Sort("-time").All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

func findtpDB(toolname, project string) ([]Log, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(*flagDBName).C(*flagCollectionName)
	err = c.Find(bson.M{"$and": []bson.M{
		bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
		bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
	}}).Sort("-time").All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

// findtpsDB 함수는 툴이름, 프로젝트, Slug를 입력받아서 로그를 검색한다.
func findtpsDB(toolname, project, slug string) ([]Log, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(*flagDBName).C(*flagCollectionName)
	err = c.Find(bson.M{"$and": []bson.M{
		bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
		bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
		bson.M{"slug": &bson.RegEx{Pattern: slug, Options: "i"}},
	}}).Sort("-time").All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

func findDB(words string) ([]Log, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(*flagDBName).C(*flagCollectionName)
	wordQueries := []bson.M{}
	for _, word := range strings.Split(words, " ") {
		wordQueries = append(wordQueries, bson.M{"$or": []bson.M{
			bson.M{"cip": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"id": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"log": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"os": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"project": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"slug": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"time": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"tool": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"user": &bson.RegEx{Pattern: word, Options: "i"}},
		}})
	}
	err = c.Find(bson.M{"$and": wordQueries}).Sort("-time").All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

func rmDB(id string) (bool, error) {
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		log.Println("DB Connet Err : ", err)
		return false, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err = session.DB(*flagDBName).C(*flagCollectionName).Remove(bson.M{"id": id})
	if err != nil {
		log.Println("DB Remove Err : ", err)
		return false, err
	}
	return true, nil

}
