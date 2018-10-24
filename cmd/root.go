package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/labstack/gommon/color"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xuender/go-utils"

	"../rich"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "go-rich",
	Short:   "致力服务小商家",
	Version: "v0.2.1",
	Long: `
  Go Rich 致力于为小商家提供客户档案，商品库存，采购销售应用`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetString(cmd, _db)
		// 备份
		bak := GetString(cmd, _bak)
		if bak != "" {
			Backup(db, bak, GetString(cmd, _name))
		}
		address := GetString(cmd, _address)
		temp := GetString(cmd, _temp)
		protocol := GetString(cmd, _protocol)
		develop := GetBool(cmd, _develop)
		// 地址端口号
		if !strings.HasPrefix(address, ":") {
			address = ":" + address
		}
		c := cmd.Root()
		web := rich.Web{
			Temp:    temp,
			DB:      db,
			Dev:     develop,
			LogFile: GetString(cmd, _logfile),
			App:     rich.NewAppVar(c.Name(), c.Short, c.Version),
		}
		if url, err := rich.GetURL(address, !strings.EqualFold(protocol, "http")); err == nil {
			web.URL = url
		} else {
			return err
		}

		// 升级
		if !develop && GetBool(cmd, _upgrade) {
			go Upgrade(cmd.Version, temp)
		}

		// 扫码提示
		color.Println(" ", color.Blue(web.URL, color.U))
		color.Println("  使用", color.Green("微信"), "扫码，即刻使用")
		color.Println(" ", color.Red("《  "+web.App.Name+"  》", color.B, color.WhtBg), color.Cyan(web.App.Version, color.In))
		color.Println("      ", color.Red(web.App.Usage, color.U))
		rich.PrintQR(web.URL)
		color.Println("请保持本窗口处于打开状态，", color.Red("CTRL+C", color.B), "退出")
		// 初始化
		if err := web.Init(); err != nil {
			return err
		}
		defer web.Close()
		// 退出
		quitChan := make(chan os.Signal)
		signal.Notify(quitChan,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGHUP,
		)
		// 运行
		switch strings.ToLower(protocol) {
		case "tls":
			go web.StartTLS(address, GetString(cmd, _cert), GetString(cmd, _key))
		case "autotls":
			go web.StartAutoTLS(address)
		default:
			go web.Start(address)
		}
		// 打开浏览器
		if GetBool(cmd, _open) {
			utils.Open(web.URL + "/qr")
		}
		<-quitChan
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "y", "", "配置文件")
	flags := rootCmd.Flags()
	flags.StringP(_address, "a", "6181", "访问地址端口号")
	flags.StringP(_protocol, "p", "http", "访问协议: http, TLS, AutoTLS")
	flags.StringP(_cert, "c", "", "TLS证书文件")
	flags.StringP(_key, "k", "", "TLS密钥文件")
	flags.StringP(_db, "d", "db", "数据库目录")
	flags.StringP(_bak, "b", "", "备份目录")
	flags.StringP(_name, "n", time.Now().Format(rich.DayFormat), "备份文件名")
	flags.StringP(_temp, "t", "temp", "临时目录")
	flags.StringP(_logfile, "l", fmt.Sprintf("%s.log", time.Now().Format(rich.DayFormat)), "日志输出文件")
	flags.BoolP(_open, "o", false, "启动浏览器显示QR码")
	flags.BoolP(_develop, "m", false, "开发模式: 静态资源读自www目录,支持跨域访问,访问日志显示")
	flags.BoolP(_upgrade, "u", true, "更新下载最新版本")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName("rich")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		cfgFile = viper.ConfigFileUsed()
		log.Println("读取配置文件:", cfgFile)
	}
}
