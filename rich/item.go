package rich

import (
	"log"
	"sort"
	"strings"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Item 商品
type Item struct {
	Obj
	Price  int64             `json:"price"`  // 价格,单位分
	Extend map[string]string `json:"extend"` // 扩展属性
	Tags   Tags              `json:"tags"`   // 标签
}

// BeforePost 创建前设置拼音标签
func (i *Item) BeforePost(key byte) goutils.ID {
	i.Obj.BeforePost(key)
	if len(i.Pinyin) > 0 {
		i.Tags.Add(strings.ToUpper(string(i.Pinyin[0])))
	}
	return i.ID
}

// BeforePut 修改前设置拼音标签
func (i *Item) BeforePut(id goutils.ID) {
	i.Obj.BeforePut(id)
	i.Tags.DelPy() // 删除原拼音标签
	if len(i.Pinyin) > 0 {
		i.Tags.Add(strings.ToUpper(string(i.Pinyin[0])))
	}
}

// Includes 包含
func (i Item) Includes(tags []string) bool {
	return i.Tags.Includes(tags)
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
	w.ObjSelect(c, &items)
	return w.ObjPaging(c, items)
}

// 商品列表
func (w *Web) items() []Item {
	if cs, ok := w.cache.Get(ItemIDPrefix); ok {
		return cs.([]Item)
	}
	cs := []Item{}
	w.Iterator([]byte{ItemIDPrefix, '-'}, func(key, value []byte) {
		c := Item{}
		if goutils.Decode(value, &c) == nil {
			cs = append(cs, c)
		} else {
			log.Printf("商品解析失败 %x \n", key)
		}
	})
	sort.Slice(cs, func(i int, j int) bool {
		return cs[i].Name < cs[j].Name
	})
	w.cache.Put(ItemIDPrefix, cs)
	return cs
}

// 商品创建
func (w *Web) itemPost(c echo.Context) error {
	w.cache.Remove(ItemIDPrefix)
	i := Item{}
	return w.ObjPost(c, &i, ItemIDPrefix, func() error { return w.Bind(c, &i) }, func() error { return w.addTags("tag-I", i.Tags) })
}

// 商品修改
func (w *Web) itemPut(c echo.Context) error {
	w.cache.Remove(ItemIDPrefix)
	i := Item{}
	return w.ObjPut(c, &i, ItemIDPrefix, func() error { return w.Bind(c, &i) }, func() error { return w.addTags("tag-I", i.Tags) })
}

// 商品删除
func (w *Web) itemDelete(c echo.Context) error {
	w.cache.Remove(ItemIDPrefix)
	return w.ObjDelete(c, ItemIDPrefix)
}
