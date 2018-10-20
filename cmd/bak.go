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
	Long:    `备份数据库及用户配置，并压缩归档`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetString(cmd, _db)
		bak := GetString(cmd, "bak-path")
		name := GetString(cmd, "name")
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
		if err := rich.Zip(db, filepath.Join(bak, fmt.Sprintf("%s.zip", name))); err != nil {
			return err
		}
		fmt.Printf("备份成功: %s.zip\n", filepath.Join(bak, name))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bakCmd)
	bakCmd.Flags().StringP(_db, "d", "db", "数据库目录")
	bakCmd.Flags().StringP("bak-path", "b", "bak", "备份目录")
	bakCmd.Flags().StringP("name", "n", time.Now().Format(rich.DayFormat), "备份文件名")
}
