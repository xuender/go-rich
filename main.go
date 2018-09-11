package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/urfave/cli"
	"github.com/xuender/goutils"

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
	app.Version = "v0.0.3"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "db,b",
			Value: "db",
			Usage: "数据库目录",
		},
		cli.StringFlag{
			Name:  "port,p",
			Value: "6181",
			Usage: "访问端口",
		},
		cli.StringFlag{
			Name:  "temp,t",
			Value: "temp",
			Usage: "临时文件目录",
		},
		cli.BoolFlag{
			Name:  "no-open,n",
			Usage: "启动不打开浏览器",
		},
		cli.BoolFlag{
			Name:  "develop-mode,d",
			Usage: "开发模式: 静态资源读自www目录,支持跨域访问,web日志输出",
		},
	}
	app.Action = runAction
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
func runAction(c *cli.Context) error {
	port := c.String("p")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	web := rich.Web{
		Port: port,
		Temp: c.String("t"),
		Db:   c.String("b"),
		Dev:  c.Bool("d"),
	}
	// 初始化
	err := web.Init()
	if err != nil {
		return err
	}
	// 打开浏览器
	if !c.Bool("n") {
		url, err := rich.GetURL(port)
		if err == nil {
			goutils.Open(url + "/qr")
		}
	}
	// 退出
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)
	// 运行
	go web.Run()
	fmt.Println(<-quitChan)
	web.Close()
	log.Println("退出")
	return nil
}
