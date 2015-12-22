CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o /lustre/INHouse/Windows/bin/dilog.exe dilog.go dbapi.go network.go timecheck.go http.go template.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lustre/INHouse/CentOS/bin/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /lustre/INHouse/OSX/bin/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o /lustre/INHouse/Tool/dilog/bin/window/dilog.exe dilog.go dbapi.go network.go timecheck.go http.go template.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /lustre/INHouse/Tool/dilog/bin/linux/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /lustre/INHouse/Tool/dilog/bin/osx/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go

#CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /show/busan/tmp/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go
