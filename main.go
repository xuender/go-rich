package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"github.com/xuender/go-utils"

	"./rich"
)

//go:generate openssl genrsa -out keys/private.rsa 1024
//go:generate openssl rsa -in keys/private.rsa -pubout -out keys/public.rsa.pub
//go:generate go run $GOROOT/src/crypto/tls/generate_cert.go --host=localhost
//go:generate mv cert.pem key.pem keys
//go:generate go-bindata -nomemcopy -pkg keys -o ./keys/bindata.go keys/private.rsa keys/public.rsa.pub keys/cert.pem keys/key.pem
//go:generate go-bindata -nomemcopy -pkg rich -o ./rich/bindata.go www/...

func main() {
	app := cli.NewApp()
	app.Name = "Go Rich"
	app.Usage = "服务小商家"
	app.Version = "v0.0.4"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address,a",
			Value: "6181",
			Usage: "访问地址的端口号",
		},
		cli.StringFlag{
			Name:  "protocol,p",
			Value: "http",
			Usage: "访问协议 http, TLS, AutoTLS",
		},
		cli.StringFlag{
			Name:  "cert,c",
			Usage: "TLS证书文件, 可省略",
		},
		cli.StringFlag{
			Name:  "key,k",
			Usage: "TLS密钥文件, 可省略",
		},
		cli.StringFlag{
			Name:  "db,b",
			Value: "db",
			Usage: "数据库目录",
		},
		cli.StringFlag{
			Name:  "temp,t",
			Value: "temp",
			Usage: "临时文件目录",
		},
		cli.StringFlag{
			Name:  "logfile,l",
			Value: "rich.log",
			Usage: "日志输出文件",
		},
		cli.BoolFlag{
			Name:  "no-open,n",
			Usage: "启动不打开浏览器",
		},
		cli.BoolFlag{
			Name:  "develop-mode,d",
			Usage: "开发模式: 静态资源读自www目录,支持跨域访问,访问日志显示",
		},
	}
	app.Action = runAction
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func runAction(c *cli.Context) error {
	address := c.String("a")
	if !strings.HasPrefix(address, ":") {
		address = ":" + address
	}
	web := rich.Web{
		Temp:    c.String("t"),
		Db:      c.String("b"),
		Dev:     c.Bool("d"),
		LogFile: c.String("l"),
	}
	if url, err := rich.GetURL(address, !strings.EqualFold(c.String("p"), "http")); err == nil {
		web.URL = url
	} else {
		return err
	}
	// 初始化
	if err := web.Init(); err != nil {
		return err
	}
	// 退出
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	// 运行
	switch strings.ToLower(c.String("p")) {
	case "tls":
		go web.StartTLS(address, c.String("c"), c.String("k"))
	case "autotls":
		go web.StartAutoTLS(address)
	default:
		go web.Start(address)
	}
	// 打开浏览器
	if !c.Bool("n") {
		utils.Open(web.URL + "/qr")
	}
	fmt.Println(<-quitChan)
	web.Close()
	return nil
}
