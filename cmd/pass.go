package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	"../rich"
)

var passCmd = &cobra.Command{
	Use:     "pass 用户名或手机号 新密码",
	Aliases: []string{"p"},
	Short:   "修改密码",
	Long:    `忘记密码后修改指定用户的登录密码`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("缺少参数，需要用户名或手机号加新密码")
		}
		db := GetString(cmd, _db)
		w := rich.Web{DB: db}
		// 初始化
		if err := w.DBInit(); err != nil {
			return err
		}
		defer w.Close()
		u := w.FindUser(args[0])
		if u == nil {
			return errors.New(fmt.Sprintf("查找用户 [ %s ] 失败", args[0]))
		}
		w.UpdatePass(u, args[1])
		fmt.Println("密码修改完毕")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(passCmd)
	passCmd.Flags().StringP(_db, "d", "db", "数据库目录")
}
