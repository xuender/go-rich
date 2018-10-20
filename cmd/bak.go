package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"

	"../rich"
)

var bakCmd = &cobra.Command{
	Use:     "bak",
	Aliases: []string{"b"},
	Short:   "数据库、配置备份",
	Long: `
  备份数据库及用户配置，并压缩归档`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Backup(GetString(cmd, _db), GetString(cmd, _bak), GetString(cmd, _name))
	},
}

// Backup 备份
func Backup(db, bak, name string) error {
	if fi, err := os.Stat(db); err == nil {
		if !fi.IsDir() {
			return errors.New(fmt.Sprintf("异常: 数据库目录 %s 不是目录。", db))
		}
	} else if os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("异常: 数据库目录 %s 不存在。", db))
	}
	if err := rich.Mkdir(bak); err != nil {
		return err
	}
	if err := rich.Zip(filepath.Join(bak, fmt.Sprintf("%s.zip", name)), db, cfgFile); err != nil {
		return err
	}
	fmt.Printf("备份成功: %s.zip\n", filepath.Join(bak, name))
	return nil
}

func init() {
	rootCmd.AddCommand(bakCmd)
	bakCmd.Flags().StringP(_db, "d", "db", "数据库目录")
	bakCmd.Flags().StringP(_bak, "b", "bak", "备份目录")
	bakCmd.Flags().StringP(_name, "n", time.Now().Format(rich.DayFormat), "备份文件名")
}
