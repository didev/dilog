# dilog
- 디지털아이디어 로그서버이다.

#### 기본정보
- 개발일시 : '15.3.24
- 관리부서 : 파이프라인

#### 사용법
- log추가하기
```
$ dilog -tool 툴이름 -log="로그내용"
```

- log추가 : 자세히
- 예시 : csi툴에서 busan프로젝트 SS_0010_org샷의 로그내용을 추가하고 보관일을 365일 1년으로 설정한다.
```
$ dilog -tool csi -project busan -slug SS_0010_org -log 로그내용 -keep 365
```

- 터미널검색 : "test"라는 문자열로 로그검색
```
$ dilog -find test
```

- 터미널검색 : "test"라는 문자열의 검색 건수만 반환
```
$ dilog -findnum test
```

- 시간이 지난 로그지우기
```
$ dilog -rm temp
```

- web서버를 8080번 포트로 실행
```
$ dilog -http=:8080
```

#### RestAPI 사용법
- curl을 이용해서 로그를 POST하는 방법

```
$ curl -X POST -d "project=TEMP&slug=SS_0010_org&tool=csi3&user=d10191&keep=180&log=log_text" http://http://10.0.90.251:8080/api/setlog
```

#### 실행방법
```
# dilog -http :80
```

#### 크로스플렛폼(맥,윈도우,리눅스) 빌드하기.
```
$ sh build.sh
```

#### HISTORY
- '16.4.8 : 회사 확장에 따라 로그시간을 국제시로 변경.
- '15.5.26 ~ '15.6.11 : 로그서버 CSI 특허준비와 같이 문서작성.
- '15.3.24 ~ '15.5.26: 설계, 1차 완료


