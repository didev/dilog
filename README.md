# dilog

![travisCI](https://secure.travis-ci.org/digital-idea/dilog.png)

디지털아이디어 웹용 로그서버 입니다.
프로젝트 매니징툴을 사용하면서 필요시 로그를 기록할 수 있도록 작성되었습니다.
로그 기능이 필요한 다른 툴도 활용할 수 있도록 설계되었습니다.

### mongoDB 설치, 실행
monogoDB를 설치하고 실행합니다.

macOS
```bash
$ brew install mongodb // macOS
$ mongod
```

CentOS
```bash
$ sudo yum install mongodb mongodb-server
```

### 로그추가: Commandline
기본적으로는 2개의 인수만 추가하더라도 로그를 사용할 수 있습니다.

```bash
$ dilog -tool 툴이름 -log="로그내용"
```

다른 인수들을 활용하여 로그를 복잡하게 넣을 수 있습니다.

> 예시 : toolA 에서 circle 프로젝트 SS_0010_org 샷에 "A버튼을 눌렀다." 라는 로그를 365일 보관

```bash
$ dilog -tool tooA -project circle -slug SS_0010_org -log "A버튼을 눌렀다." -keep 365
```

### 로그추가: RestAPI
- curl을 이용해서 로그를 POST하는 방법

```
$ curl -X POST -d "project=circle&slug=SS_0010_org&tool=csi3&user=woong&keep=180&log=log_text" http://127.0.0.1:8080/api/setlog
```

### 로그검색: Commandline
터미널검색 : "test"라는 문자열로 로그검색

```bash
$ dilog -find test
```

### 로그삭제: Commandline
기본적으로 로그 데이터는 가지고 있습니다.
아래 명령어를 통해서 시간이 지난 로그를 지울 수 있습니다.

```bash
$ dilog -rm
```

### 웹서버 실행
web서버를 8080번 포트로 실행

```bash
$ dilog -http :8080 -dbip 10.0.90.251
```

### Reference
- vfxgen: https://github.com/shurcooL/vfsgen

### History
- '18.6.26: RestAPI 추가
- '16.4.8: 회사 확장에 따라 로그시간을 국제시로 변경.
- '15.3.24~'15.5.26: 설계, 1차 완료


