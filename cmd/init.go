package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// 命令参数
	_db       = "db-path"
	_temp     = "temp-path"
	_cert     = "tls-cert"
	_key      = "tls-key"
	_open     = "open-browser"
	_address  = "address"
	_protocol = "protocol"
	_logfile  = "log-file"
	_develop  = "develop-mode"
	_upgrade  = "upgrade"
	_bak      = "bak-path"
	_name     = "bak-name"
)

// GetString 读取配置
func GetString(cmd *cobra.Command, name string) string {
	f := cmd.Flag(name)
	// 命令行优先
	if f.Changed {
		return f.Value.String()
	}
	ret := viper.GetString(name)
	if ret == "" {
		return f.Value.String()
	}
	return ret
}

// GetString 读取配置
func GetBool(cmd *cobra.Command, name string) bool {
	f := cmd.Flag(name)
	b, _ := strconv.ParseBool(f.Value.String())
	// 命令行优先
	if f.Changed {
		return b
	}
	ret := viper.GetString(name)
	if ret == "" {
		return b
	}
	return viper.GetBool(name)
}
