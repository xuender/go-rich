# 获取最后的tag
TAG = $(shell git tag | tail -n 1)

all: web generate build pack gren

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
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o rich-i386.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o rich-amd64.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o rich-linux main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o rich-darwin main.go

pack:
	@echo 打包 $(TAG)
	mv rich-i386.exe go-rich.exe
	zip go-rich-$(TAG)-windows_i386.zip go-rich.exe rich.bat
	rm go-rich.exe
	mv rich-amd64.exe go-rich.exe
	zip go-rich-$(TAG)-windows_amd64.zip go-rich.exe rich.bat
	rm go-rich.exe
	mv rich-linux go-rich
	zip go-rich-$(TAG)-linux_amd64.zip go-rich
	rm go-rich
	mv rich-darwin go-rich
	zip go-rich-$(TAG)-darwin_amd64.zip go-rich
	rm go-rich

gren:
	@echo 生成文档
	gren release --override
	gren changelog --override

todo:
	@echo 寻找未完成代码
	grep -rnw "TODO" rich src

test:
	@echo 测试
	go test ./rich
