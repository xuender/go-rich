package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"../rich"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Aliases: []string{"u"},
	Short:   "升级",
	Long:    `更新下载最新版本`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return Upgrade(cmd.Root().Version, GetString(cmd, _temp))
	},
}

func Upgrade(version, temp string) error {
	if err := rich.Mkdir(temp); err != nil {
		return err
	}
	if url, tag, err := rich.LastURL("xuender", "go-rich", version); err == nil {
		fmt.Println("更新版本:", tag, "网址:", url)
		rich.Download(url, temp)
		os.Rename(filepath.Join(temp, "go-rich"), "go-rich")
		fmt.Println("更新成功:", tag)
		return nil
	} else {
		return err
	}
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
	upgradeCmd.Flags().StringP(_temp, "t", "temp", "临时目录")
	upgradeCmd.Runnable()
}
