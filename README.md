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

go get github.com/jteeuwen/go-bindata/...
go get github.com/360EntSecGroup-Skylar/excelize
go get github.com/Code-Hex/echo-static
go get github.com/dgrijalva/jwt-go
go get github.com/elazarl/go-bindata-assetfs
go get github.com/labstack/echo
go get github.com/mozillazg/go-slugify
go get github.com/syndtr/goleveldb/leveldb
go get github.com/urfave/cli
go get github.com/xuender/goutils
go get github.com/smartystreets/goconvey
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
