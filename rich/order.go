package rich

type Order struct {
	ItemId []byte // 商品
	Num    int64  // 数量
	Price  int64  // 价格,单位分
}
