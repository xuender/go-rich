package rich

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// Trade 订单
type Trade struct {
	ID     utils.ID    `json:"id"`                      // 主键
	CID    utils.ID    `json:"cid"`                     // 客户ID
	Status TradeStatus `json:"status"`                  // 状态
	Ca     time.Time   `json:"ca"`                      // 创建时间
	Pa     time.Time   `json:"pa"`                      // 付款时间
	Orders []Order     `json:"orders"`                  // 订单详情
	Total  int64       `json:"total"`                   // 总额,单位分
	Note   string      `json:"note" validate:"lte=100"` // 备注 最长100
}

// BeforePost 订单新建前
func (t *Trade) BeforePost(key byte) utils.ID {
	t.ID = utils.NewID(key)
	t.Ca = time.Now()
	return t.ID
}

// BeforePut 修改订单
func (t *Trade) BeforePut(id utils.ID) {
	t.ID = id
}

// 订单路由
func (w *Web) tradeRoute(c *echo.Group) {
	c.GET("", func(c echo.Context) error { return c.JSON(http.StatusOK, w.days) }) // 订单列表
	c.GET("/:id", w.tradeGet)                                                      // 订单列表
	c.POST("", w.tradePost)                                                        // 订单创建
	c.PUT("/:id", w.tradePut)                                                      // 订单修改
	c.DELETE("/:id", w.tradeDelete)                                                // 订单删除
}

// 订单列表
func (w *Web) tradeGet(c echo.Context) error {
	id := c.Param("id")
	log.Println(id)
	log.Println(w.days.Includes(id))
	if !w.days.Includes(id) {
		return c.JSON(http.StatusOK, []int{})
	}
	// TODO 日期校验判断日期是否在
	ids := []utils.ID{}
	w.Get([]byte(id), &ids)
	trades := []Trade{}
	for _, i := range ids {
		t := Trade{}
		w.Get(i[:], &t)
		trades = append(trades, t)
	}
	return c.JSON(http.StatusOK, trades)
}

// 订单创建
func (w *Web) tradePost(c echo.Context) error {
	t := Trade{}
	return w.ObjPost(c, &t, TradeIDPrefix, func() error { return w.Bind(c, &t) }, func() error {
		day := t.Ca.Format("2006-01-02")
		w.days.Add(day)
		w.Put(DaysKey, w.days)
		idsKey := []byte(day)
		ids := []utils.ID{}
		w.Get(idsKey, &ids)
		ids = append(ids, t.ID)
		w.Put(idsKey, ids)
		log.Println(w.days)
		return nil
	})
}

// 订单修改
func (w *Web) tradePut(c echo.Context) error {
	t := Trade{}
	return w.ObjPut(c, &t, TradeIDPrefix, func() error { return w.Bind(c, &t) }, func() error { return nil })
}

// 订单删除
func (w *Web) tradeDelete(c echo.Context) error {
	return w.ObjDelete(c, TradeIDPrefix)
}
