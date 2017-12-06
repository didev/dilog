GOOS=windows GOARCH=amd64 go build -o /lustre/INHouse/Windows/bin/dilog.exe dilog.go dbapi.go network.go timecheck.go http.go template.go
GOOS=linux GOARCH=amd64 go build -o /lustre/INHouse/CentOS/bin/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go
GOOS=darwin GOARCH=amd64 go build -o /lustre/INHouse/OSX/bin/dilog dilog.go dbapi.go network.go timecheck.go http.go template.go
