set GOOS=linux
set GOARCH=amd64
go build cmd\main.go
del main.zip
tar.exe -a -cf main.zip main