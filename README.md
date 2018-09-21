# Go Rich

## 编译

### GO 安装
```
git clone https://github.com/golang/net.git $GOPATH/src/github.com/golang/net
git clone https://github.com/golang/sys.git $GOPATH/src/github.com/golang/sys
git clone https://github.com/golang/tools.git $GOPATH/src/github.com/golang/tools
git clone https://github.com/golang/crypto.git $GOPATH/src/github.com/golang/crypto
mkdir $GOPATH/src/golang.org
ln -s $GOPATH/src/github.com/golang $GOPATH/src/golang.org/x

# WEB框架
go get github.com/labstack/echo
go get github.com/Code-Hex/echo-static
# 校验
go get gopkg.in/go-playground/validator.v9
# 资源打包
go get github.com/jteeuwen/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs
# 数据库
go get github.com/syndtr/goleveldb/leveldb
# 命令行
go get github.com/urfave/cli
# 认证
go get github.com/dgrijalva/jwt-go
# Excel
go get github.com/360EntSecGroup-Skylar/excelize
# 汉字转拼音
go get github.com/mozillazg/go-slugify
# 工具
go get github.com/xuender/goutils
# 单元测试
go get github.com/smartystreets/goconvey
# 二维码生成
go get rsc.io/qr
```
### 编译
```shell
npm i
make
```
## 设计原则
* uri 中不能有动词,资源是名词
* 名词复数形式
