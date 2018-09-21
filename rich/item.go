package rich

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Item 商品
type Item struct {
	Obj
	Price  int64             `json:"price"`  // 价格,单位分
	Extend map[string]string `json:"extend"` // 扩展属性
}

// 商品路由
func (w *Web) itemRoute(c *echo.Group) {
	c.GET("", w.itemsGet)          // 商品列表
	c.POST("", w.itemPost)         // 商品创建
	c.PUT("/:id", w.itemPut)       // 商品修改
	c.DELETE("/:id", w.itemDelete) // 商品删除
	// c.DELETE("", w.customersDelete)     // 清除客户
	// c.POST("/file", w.customersFile)    // 上传客户文件
}

// 商品列表
func (w *Web) itemsGet(c echo.Context) error {
	items := w.items()
	w.ObjSearch(c, &items)
	return c.JSON(http.StatusOK, items)
}

// 商品列表
func (w *Web) items() []Item {
	cs := []Item{}
	w.Iterator([]byte{ItemIDPrefix, '-'}, func(key, value []byte) {
		c := Item{}
		if goutils.Decode(value, &c) == nil {
			cs = append(cs, c)
		} else {
			log.Printf("商品解析失败 %x \n", key)
		}
	})
	return cs
}

// 商品创建
func (w *Web) itemPost(c echo.Context) error {
	return w.ObjPost(c, &Item{}, ItemIDPrefix, func() error { return nil })
}

// 商品修改
func (w *Web) itemPut(c echo.Context) error {
	return w.ObjPut(c, &Item{}, ItemIDPrefix, func() error { return nil })
}

// 商品删除
func (w *Web) itemDelete(c echo.Context) error {
	return w.ObjDelete(c, ItemIDPrefix)
}
