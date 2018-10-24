package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"../rich"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Aliases: []string{"c"},
	Short:   "清理优化数据库",
	Long:    `整理帐目，删除的客户、商品记录，`,
	RunE: func(cmd *cobra.Command, args []string) error {
		db := GetString(cmd, _db)
		web := rich.Web{DB: db}
		// 初始化
		if err := web.DBInit(); err != nil {
			return err
		}
		defer web.Close()
		if err := web.Reset(); err != nil {
			return err
		}
		fmt.Println("数据库清理完毕。")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().StringP(_db, "d", "db", "数据库目录")
}
