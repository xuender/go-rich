package rich

import (
	"time"

	"github.com/xuender/goutils"
)

type Customer struct {
	Id     goutils.ID   `json:"id"`               // 主键
	Name   string       `json:"name"`             // 姓名
	Ca     time.Time    `json:"ca"`               // 创建时间
	Trades []goutils.ID `json:"trades,omitempty"` // 消费记录
	Note   string       `json:"note"`             // 备注
}
