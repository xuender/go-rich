package rich

import "github.com/urfave/cli"

// App 应用信息
type App struct {
	Name    string `json:"name"`    // 应用名称
	Usage   string `json:"usage"`   // 使用方法
	Version string `json:"version"` // 版本
	IsNew   bool   `json:"isNew"`   // 新应用
}

// NewApp 新建应用
func NewApp(app *cli.App) *App {
	a := App{}
	a.Name = app.Name
	a.Usage = app.Usage
	a.Version = app.Version
	a.IsNew = true
	return &a
}
