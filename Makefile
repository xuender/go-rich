TAG = $(shell git tag)
all: web generate build pack

web:
	@echo 编译 web
	ionic build --prod

generate:
	@echo 生成密钥和资源文件
	rm ./keys/bindata.go
	rm ./rich/bindata.go
	go generate

build:
	@echo 交叉编译
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o rich-386.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o rich-amd64.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rich-linux main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o rich-darwin main.go

pack:
	@echo 打包
	zip rich-386.zip rich-386.exe
	rm rich-386.exe
	zip rich-amd64.zip rich-amd64.exe
	rm rich-amd64.exe
	zip rich-linux.zip rich-linux
	rm rich-linux
	zip rich-darwin.zip rich-darwin
	rm rich-darwin

todo:
	@echo 寻找未完成代码
	grep -rnw "TODO" rich src

test:
	@echo 测试
	go test ./rich
