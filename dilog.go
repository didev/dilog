package dilog

import (
	"log"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	// DBNAME 은 데이터베이스 이름이다.
	DBNAME = "dilog"
	// COLLECTION 은 MongoDB Collection 이름이다.
	COLLECTION = "log"
)

// Log 자료구조 이다.
type Log struct {
	Cip     string // Client IP
	ID      string // log ID
	Keep    int    // 보관일수
	Log     string // 로그내용
	Project string // 프로젝트
	Slug    string // 로그에 입력되는 Slug 예) 프로젝트 매니징툴에서 사용되는 에셋명 또는 샷명
	Time    string // 로그가 기입된 시간
	Tool    string // 로그가 보내진 툴
	User    string // 유저이름
}

// Add 함수는 log 를 추가합니다.
func Add(dbip, ip, logstr, project, slug, tool, user string, keep int) error {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(DBNAME).C(COLLECTION)
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

// All 는 DB의 모든 로그를 가지고 옵니다.
func All(dbip string) ([]Log, error) {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(DBNAME).C(COLLECTION)
	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, err
	}
	return results, nil
}

// FindTool 함수는 툴이름을 이용해서 log를 검색합니다.
func FindTool(dbip, toolname string, page, pageMaxItemNum int) ([]Log, int, error) {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, 0, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(DBNAME).C(COLLECTION)
	query := bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}}
	err = c.Find(query).Sort("-time").Skip((page - 1) * pageMaxItemNum).Limit(pageMaxItemNum).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	itemNum, err := c.Find(query).Count()
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	return results, totalPage(itemNum, pageMaxItemNum), nil
}

// TotalPage 함수는 아이템의 갯수를 이용해서 총 페이지수를 반환한다.
func totalPage(itemNum, pageMaxItemNum int) int {
	page := itemNum / pageMaxItemNum
	if itemNum%pageMaxItemNum != 0 {
		page++
	}
	return page
}

// FindtpDB 함수는 툴이름, 프로젝트 이름을 이용해서 log를 검색합니다.
func FindToolProject(dbip, toolname, project string, page, pageMaxItemNum int) ([]Log, int, error) {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, 0, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(DBNAME).C(COLLECTION)
	query := bson.M{"$and": []bson.M{
		bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
		bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
	}}
	err = c.Find(query).Sort("-time").Skip(page - 1).Limit(pageMaxItemNum).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	itemNum, err := c.Find(query).Count()
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	return results, totalPage(itemNum, pageMaxItemNum), nil
}

// FindtpsDB 함수는 툴이름, 프로젝트, Slug를 입력받아서 로그를 검색한다.
func FindToolProjectSlug(dbip, toolname, project, slug string, page, pageMaxItemNum int) ([]Log, int, error) {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, 0, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(DBNAME).C(COLLECTION)
	query := bson.M{"$and": []bson.M{
		bson.M{"tool": &bson.RegEx{Pattern: toolname, Options: "i"}},
		bson.M{"project": &bson.RegEx{Pattern: project, Options: "i"}},
		bson.M{"slug": &bson.RegEx{Pattern: slug, Options: "i"}},
	}}
	err = c.Find(query).Sort("-time").Skip((page - 1) * pageMaxItemNum).Limit(pageMaxItemNum).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	itemNum, err := c.Find(query).Count()
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	return results, totalPage(itemNum, pageMaxItemNum), nil
}

// Search 함수는 검색어를 이용해서 log를 검색합니다.
func Search(dbip, words string, page, pageMaxItemNum int) ([]Log, int, error) {
	session, err := mgo.Dial(dbip)
	if err != nil {
		log.Println("DB Connect Err : ", err)
		return nil, 0, err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []Log
	c := session.DB(DBNAME).C(COLLECTION)
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
			bson.M{"time": word},
			bson.M{"tool": &bson.RegEx{Pattern: word, Options: "i"}},
			bson.M{"user": &bson.RegEx{Pattern: word, Options: "i"}},
		}})
	}
	err = c.Find(bson.M{"$and": wordQueries}).Sort("-time").Skip((page - 1) * pageMaxItemNum).Limit(pageMaxItemNum).All(&results)
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	itemNum, err := c.Find(bson.M{"$and": wordQueries}).Count()
	if err != nil {
		log.Println("DB Find Err : ", err)
		return nil, 0, err
	}
	return results, totalPage(itemNum, pageMaxItemNum), nil
}

// Remove 함수는 log id를 이용해서 로그를 삭제합니다.
func Remove(dbip, id string) error {
	session, err := mgo.Dial(dbip)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	err = session.DB(DBNAME).C(COLLECTION).Remove(bson.M{"id": id})
	if err != nil {
		return err
	}
	return nil

}

// Timecheck 함수는 RFC3339 시간문자열과 보관일을 받아서 보관일이 지난지 체크한다.
func Timecheck(timestr string, keepdate int) bool {
	t, err := time.Parse(time.RFC3339, timestr)
	if err != nil {
		log.Println(err)
	}
	addtime := t.AddDate(0, 0, keepdate)
	now := time.Now()
	return addtime.After(now) //추후 이 결과를 이용해서 참이면 리무브 대상이다.
}
