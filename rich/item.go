package rich

import (
	"errors"
	"net/http"
	"os"
	"sort"
	"strings"
	"unsafe"

	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// Item 商品
type Item struct {
	Obj
	Price  int64             `json:"price"`  // 价格,单位分
	Cost   int64             `json:"cost"`   // 成本，单位分
	Batchs []Batch           `json:"batchs"` // 批次
	Extend map[string]string `json:"extend"` // 扩展属性
	Tags   Tags              `json:"tags"`   // 标签
}

// Batch 批次
type Batch struct {
	Cost      int64 `json:"cost"`      // 成本，单位分
	Total     int   `json:"total"`     // 采购总量
	Inventory int   `json:"inventory"` // 库存
}

// BeforePost 创建前设置拼音标签
func (i *Item) BeforePost(key byte) utils.ID {
	i.Obj.BeforePost(key)
	for _, f := range Initial(i.Name) {
		i.Tags.Add(f)
	}
	return i.ID
}

// BeforePut 修改前设置拼音标签
func (i *Item) BeforePut(id utils.ID) {
	i.Obj.BeforePut(id)
	i.Tags.DelPy() // 删除原拼音标签
	for _, f := range Initial(i.Name) {
		i.Tags.Add(f)
	}
}

// Includes 包含
func (i Item) Includes(tags []string) bool {
	return i.Tags.Includes(tags)
}

// 商品路由
func (w *Web) itemRoute(c *echo.Group) {
	c.GET("", w.itemsGet)                                                     // 商品列表
	c.POST("", w.itemPost)                                                    // 商品创建
	c.PUT("/:id", w.itemPut)                                                  // 商品修改
	c.DELETE("/:id", w.itemDelete)                                            // 商品删除
	c.GET("/:id", func(c echo.Context) error { return w.ObjGet(c, &Item{}) }) // 商品列表
	c.POST("/file", w.itemsFile)                                              // 上传商品文件
}

// 商品列表
func (w *Web) itemsGet(c echo.Context) error {
	items := w.items(true)
	w.ObjSearch(c, &items)
	w.ObjSelect(c, &items)
	return w.ObjPaging(c, items)
}

// 商品列表
func (w *Web) items(cache bool) []Item {
	if cache {
		if cs, ok := w.cache[ItemIDPrefix]; ok {
			return cs.([]Item)
		}
	}
	cs := []Item{}
	w.Iterator([]byte{ItemIDPrefix, '-'}, func(key, value []byte) {
		c := Item{}
		if utils.Decode(value, &c) == nil && !c.IsDelete() {
			cs = append(cs, c)
		}
	})
	sort.Slice(cs, func(i int, j int) bool {
		return cs[i].Name < cs[j].Name
	})
	w.cache[ItemIDPrefix] = cs
	return cs
}

// 商品创建
func (w *Web) itemPost(c echo.Context) error {
	delete(w.cache, ItemIDPrefix)
	i := Item{}
	return w.ObjPost(c, &i, ItemIDPrefix, func() error { return w.Bind(c, &i) }, func() error { return w.addTags(TagKeys[1], i.Tags) })
}

// 商品修改
func (w *Web) itemPut(c echo.Context) error {
	delete(w.cache, ItemIDPrefix)
	i := Item{}
	return w.ObjPut(c, &i, ItemIDPrefix, func() error { return w.Bind(c, &i) }, func() error { return w.addTags(TagKeys[1], i.Tags) })
}

// 商品删除
func (w *Web) itemDelete(c echo.Context) error {
	return w.ObjDeleter(c, ItemIDPrefix, &Item{}, func() error {
		w.tagsReset()
		delete(w.cache, ItemIDPrefix)
		return nil
	})
}

// 商品信息上传
func (w *Web) itemsFile(c echo.Context) error {
	delete(w.cache, ItemIDPrefix)
	// 读取Excel定义
	xid := new(utils.ID)
	if err := xid.Parse(c.Request().Header.Get("xlsx")); err != nil {
		return err
	}
	xlsx := Xlsx{}
	w.Get(xid[:], &xlsx)
	// 读取Excel文件
	file, err := w.saveTemp(c)
	if err != nil {
		return err
	}
	items := []Item{}
	err = ReadXlsx(file, func(row []string) {
		if is, e := newItems(row, xlsx.Map); e == nil {
			items = append(items, is...)
		}
	})
	// 生成商品信息
	if err == nil {
		for _, i := range items {
			i.BeforePost(ItemIDPrefix)
			w.addTags(TagKeys[1], i.Tags)
			w.Put(i.ID[:], i)
		}
		os.Remove(file)
		w.itemsMerge()
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusInternalServerError, err.Error())
}

// 新建商品
func newItems(row []string, m map[int]string) (items []Item, err error) {
	i := Item{}
	p, err := utils.Parse(row, m, &i)
	if err != nil {
		return
	}
	i.Cost = i.Cost * 100
	i.Price = i.Price * 100
	i.Extend = p
	for _, name := range []string{"", "商品", "货品", "商品名称"} {
		if i.Name == name {
			err = errors.New("商品名称错误")
			return
		}
	}
	if strings.Contains(i.Name, "+") {
		names := strings.Split(i.Name, "+")
		if len(names) == 1 {
			items = append(items, i)
			return
		}
		for _, name := range names {
			name = strings.Trim(name, " ")
			if name != "" {
				var b = *(*Item)(unsafe.Pointer(&i))
				b.Name = name
				items = append(items, b)
			}
		}
	}
	return
}

// itemsMerge 商品信息合并
func (w *Web) itemsMerge() {
	m := map[string]Item{}
	for _, c := range w.items(true) {
		k := c.Name
		if v, ok := m[k]; ok {
			// 重复
			for _, t := range c.Tags {
				v.Tags.Add(t)
			}
			w.Put(v.ID[:], v)
			w.Delete(c.ID[:])
		} else {
			m[k] = c
		}
	}
	delete(w.cache, ItemIDPrefix)
}

// 商品扣减库存
func (w *Web) deducting(order Order) error {
	delete(w.cache, ItemIDPrefix)
	i := Item{}
	if err := w.Get(order.ID[:], &i); err != nil {
		return err
	}
	i.deducting(order)
	return w.Put(i.ID[:], i)
}
func (i *Item) deducting(order Order) {
	num := order.Num
	for f, b := range i.Batchs {
		if b.Inventory >= num {
			i.Batchs[f].Inventory -= num
			break
		} else {
			num -= b.Inventory
			i.Batchs[f].Inventory = 0
		}
	}
}

// 商品撤销扣减
func (w *Web) unDeducting(order Order) error {
	delete(w.cache, ItemIDPrefix)
	i := Item{}
	if err := w.Get(order.ID[:], &i); err != nil {
		return err
	}
	i.unDeducting(order)
	return w.Put(i.ID[:], i)
}
func (i *Item) unDeducting(order Order) {
	num := order.Num
	reverse(i.Batchs)
	for f, b := range i.Batchs {
		if b.Inventory+num <= b.Total {
			i.Batchs[f].Inventory += num
			break
		} else {
			num -= b.Total - b.Inventory
			i.Batchs[f].Inventory = b.Total
		}
	}
	reverse(i.Batchs)
}
func reverse(s []Batch) []Batch {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
