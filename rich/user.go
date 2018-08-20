package rich

import (
	"time"

	"github.com/xuender/goutils"
)

// 用户
type User struct {
	Id       goutils.ID `json:"id"`    // 主键
	Nick     string     `json:"nick"`  // 昵称
	Phone    string     `json:"phone"` // 电话
	Password string     `json:"-"`     // 密码
	Ca       time.Time  `json:"ca"`    // 创建时间
}
