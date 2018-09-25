package rich

// Paging 分页对象
type Paging struct {
	Data  interface{} `json:"data"`  // 数据
	Total int         `json:"total"` // 总数
}
