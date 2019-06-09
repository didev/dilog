#!/bin/sh
APP="dilog"

# OS별로 빌드함.
# assets 폴더의 모든 에셋을 빌드전에 assets_vfsdata.go 파일로 생성한다.
go run assets/asset_generate.go

# OS별 필드
GOOS=linux GOARCH=amd64 go build -o ./bin/linux/${APP} dilog.go struct.go dbapi.go network.go timecheck.go http.go templatefunc.go assets_vfsdata.go
GOOS=windows GOARCH=amd64 go build -o ./bin/windows/${APP}.exe dilog.go struct.go dbapi.go network.go timecheck.go http.go templatefunc.go assets_vfsdata.go
GOOS=darwin GOARCH=amd64 go build -o ./bin/darwin/${APP} dilog.go struct.go dbapi.go network.go timecheck.go http.go templatefunc.go assets_vfsdata.go

# Github Release에 업로드 하기위해 압축
cd ./bin/linux/ && tar -zcvf ../${APP}_linux_x86-64.tgz . && cd -
cd ./bin/windows/ && tar -zcvf ../${APP}_windows_x86-64.tgz . && cd -
cd ./bin/darwin/ && tar -zcvf ../${APP}_darwin_x86-64.tgz . && cd -

# 삭제
rm -rf ./bin/linux
rm -rf ./bin/windows
rm -rf ./bin/darwin