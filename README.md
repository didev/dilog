# dilog

![travisCI](https://secure.travis-ci.org/digital-idea/dilog.png)

디지털아이디어 웹용 로그서버 입니다.


### 로그추가(Commandline)
- 기본방법

```
$ dilog -tool 툴이름 -log="로그내용"
```

- 어드벤스드 방법
- 예시 : csi툴에서 busan프로젝트 SS_0010_org샷의 로그내용을 추가하고 보관일을 365일 1년으로 설정한다.

```
$ dilog -tool csi -project busan -slug SS_0010_org -log 로그내용 -keep 365
```

#### 로그추가 : RestAPI 사용법
- curl을 이용해서 로그를 POST하는 방법

```
$ curl -X POST -d "project=TEMP&slug=SS_0010_org&tool=csi3&user=d10191&keep=180&log=log_text" http://http://10.0.90.251:8080/api/setlog
```

#### 로그검색(Commandline)
- 터미널검색 : "test"라는 문자열로 로그검색

```
$ dilog -find test
```

- 터미널검색 : "test"라는 문자열의 검색 건수만 반환

```
$ dilog -findnum test
```

#### 로그삭제(Commandline)
- 시간이 지난 로그지우기

```
$ dilog -rm temp
```

#### 웹서버 실행방법
- web서버를 80번 포트로 실행
```
# dilog -http :80 -dbip 10.0.90.251
```

#### HISTORY
- '18.6.26 : RestAPI 추가
- '16.4.8 : 회사 확장에 따라 로그시간을 국제시로 변경.
- '15.5.26 ~ '15.6.11 : 로그서버 CSI 특허준비와 같이 문서작성.
- '15.3.24 ~ '15.5.26: 설계, 1차 완료


