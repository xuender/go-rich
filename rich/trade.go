package rich

import (
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
	Cost   int64       `json:"cost"`                    // 成本，单位分
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
	c.GET("", w.tradesGet)          // 订单列表
	c.GET("/:id", w.tradeGet)       // 订单列表
	c.POST("", w.tradePost)         // 订单创建
	c.PUT("/:id", w.tradePut)       // 订单修改
	c.DELETE("/:id", w.tradeDelete) // 订单删除
}

// 订单列表
func (w *Web) tradesGet(c echo.Context) error {
	day := c.QueryParam("day")
	if day == "" {
		return c.JSON(http.StatusOK, w.days)
	}
	if !w.days.Includes(day) {
		return c.JSON(http.StatusOK, []int{})
	}
	// TODO 日期校验判断日期是否在
	ids := []utils.ID{}
	w.Get([]byte(day), &ids)
	trades := []Trade{}
	for _, i := range ids {
		t := Trade{}
		w.Get(i[:], &t)
		trades = append(trades, t)
	}
	return c.JSON(http.StatusOK, trades)
}

// 订单获取
func (w *Web) tradeGet(c echo.Context) error {
	return w.ObjGet(c, &Trade{})
}

// 订单创建
func (w *Web) tradePost(c echo.Context) error {
	t := Trade{}
	return w.ObjPost(c, &t, TradeIDPrefix, func() error { return w.Bind(c, &t) }, func() error {
		w.dayAdd(t)
		if !t.CID.IsNew() {
			w.customerAddTrade(t.CID, t.ID)
		}
		return nil
	})
}

// 订单修改
func (w *Web) tradePut(c echo.Context) error {
	t := Trade{}
	return w.ObjPut(c, &t, TradeIDPrefix, func() error { return w.Bind(c, &t) }, func() error {
		old := Trade{}
		w.Get(t.ID[:], &old)
		if !t.CID.IsNew() {
			if !old.CID.IsNew() && !old.CID.Equal(t.CID) {
				// 删除旧订单
				w.customerDelTrade(old.CID, t.ID)
				// 增加新订单
				w.customerAddTrade(t.CID, t.ID)
			}
			if old.CID.IsNew() {
				// 增加新订单
				w.customerAddTrade(t.CID, t.ID)
			}
		}

		return nil
	})
}

// 订单删除
func (w *Web) tradeDelete(c echo.Context) error {
	return w.ObjDelete(c, TradeIDPrefix, func(id utils.ID) error {
		t := Trade{}
		w.Get(id[:], &t)
		if !t.CID.IsNew() {
			w.customerDelTrade(t.CID, t.ID)
		}
		w.dayDel(t)
		return nil
	})
}
