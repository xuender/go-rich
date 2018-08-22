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
//go:generate go-bindata -nomemcopy -pkg keys -o ./keys/bindata.go keys/private.rsa keys/public.rsa.pub
//go:generate go-bindata -nomemcopy -pkg rich -o ./rich/bindata.go www/...

func main() {
	app := cli.NewApp()
	app.Name = "rich"
	app.Usage = "Rich 小商家服务"
	app.Version = "0.0.1"
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
	}
	// 初始化
	err := web.Init()
	if err != nil {
		return err
	}
	// 打开浏览器
	if !c.Bool("n") {
		url, err := rich.GetUrl(port)
		if err == nil {
			goutils.Open(url)
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
