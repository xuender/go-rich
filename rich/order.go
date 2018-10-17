package rich

import (
	"github.com/xuender/go-utils"
)

// Order 订单子项
type Order struct {
	ID    utils.ID `json:"id"`    // 商品
	Num   int      `json:"num"`   // 数量
	Price int64    `json:"price"` // 价格,单位分
	Cost  int64    `json:"cost"`  // 成本，单位分
}
