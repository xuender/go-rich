package rich

type Paging struct {
	Data  interface{} `json:"data"`  // 数据
	Total int         `json:"total"` // 总数
}
