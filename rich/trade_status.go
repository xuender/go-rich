package rich

type TradeStatus int

const (
	预订 TradeStatus = iota
	待付款
	关闭 // 结束
	待发货
	待确认
	退货
	退款 // 结束
	完成 // 结束
)
