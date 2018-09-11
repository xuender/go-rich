package rich

// Order 订单子项
type Order struct {
	ItemID []byte // 商品
	Num    int64  // 数量
	Price  int64  // 价格,单位分
}
