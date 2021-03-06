package rich

// App 应用信息
type App struct {
	Name    string `json:"name"`    // 应用名称
	Usage   string `json:"usage"`   // 使用方法
	Version string `json:"version"` // 版本
	IsNew   bool   `json:"isNew"`   // 新应用
}

// NewAppVar 新建应用
func NewAppVar(name, usage, version string) *App {
	a := App{}
	a.Name = name
	a.Usage = usage
	a.Version = version
	a.IsNew = true
	return &a
}
