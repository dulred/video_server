go mod init example.com/myapp

go mod tidy

go env -w GOPROXY=https://goproxy.cn,direct

go env | grep GOPROXY

go build

api.exe

cd 到指定test 目录用 go test -v