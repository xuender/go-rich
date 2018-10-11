package rich

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// Customer 客户
type Customer struct {
	Obj
	Phone  string            `json:"phone,omitempty"`  // 电话
	Trades []utils.ID        `json:"trades,omitempty"` // 消费记录
	Extend map[string]string `json:"extend"`           // 扩展属性
	Tags   Tags              `json:"tags"`             // 标签
}

// BeforePost 创建前设置拼音标签
func (c *Customer) BeforePost(key byte) utils.ID {
	c.Obj.BeforePost(key)
	for _, i := range Initial(c.Name) {
		c.Tags.Add(i)
	}
	return c.ID
}

// BeforePut 修改前设置拼音标签
func (c *Customer) BeforePut(id utils.ID) {
	c.Obj.BeforePut(id)
	c.Tags.DelPy() // 删除原拼音标签
	for _, i := range Initial(c.Name) {
		c.Tags.Add(i)
	}
}

// Includes 包含
func (c Customer) Includes(tags []string) bool {
	return c.Tags.Includes(tags)
}

// Match 匹配
func (c Customer) Match(txt string) bool {
	return c.Obj.Match(txt) || strings.Contains(c.Phone, txt)
}

// 客户路由
func (w *Web) customerRoute(c *echo.Group) {
	c.GET("", w.customersGet)                                                     // 客户列表
	c.POST("", w.customerPost)                                                    // 客户创建
	c.PUT("/:id", w.customerPut)                                                  // 客户修改
	c.DELETE("/:id", w.customerDelete)                                            // 删除客户
	c.DELETE("", w.customersDelete)                                               // 清除客户
	c.POST("/file", w.customersFile)                                              // 上传客户文件
	c.GET("/:id", func(c echo.Context) error { return w.ObjGet(c, &Customer{}) }) // 客户查看
}

// 给客户增加订单
func (w *Web) customerAddTrade(cid, tid utils.ID) error {
	delete(w.cache, CustomerIDPrefix)
	c := Customer{}
	w.Get(cid[:], &c)
	for _, id := range c.Trades {
		if id.Equal(tid) {
			return errors.New(fmt.Sprint("重复增加订单:", tid))
		}
	}
	c.Trades = append(c.Trades, tid)
	return w.Put(cid[:], c)
}

// 删除客户的订单
func (w *Web) customerDelTrade(cid, tid utils.ID) error {
	delete(w.cache, CustomerIDPrefix)
	c := Customer{}
	w.Get(cid[:], &c)
	for i, id := range c.Trades {
		if id.Equal(tid) {
			c.Trades = append(c.Trades[:i], c.Trades[i+1:]...)
			return w.Put(cid[:], c)
		}
	}
	return errors.New(fmt.Sprint("不存在:", tid))
}

// 客户列表
func (w *Web) customersGet(c echo.Context) error {
	customers := w.customers()
	w.ObjSearch(c, &customers)
	w.ObjSelect(c, &customers)
	excel := c.QueryParam("excel")
	if excel != "" {
		id := new(utils.ID)
		err := id.Parse(excel)
		if err != nil {
			return err
		}
		return w.excel(c, customers, id)
	}
	return w.ObjPaging(c, customers)
}

// 导出excel
func (w *Web) excel(c echo.Context, cs []Customer, id *utils.ID) error {
	xlsx := Xlsx{}
	w.Get(id[:], &xlsx)
	// 生成xlsx
	x := excelize.NewFile()
	m := map[string]string{"Name": "姓名", "Phone": "联系方式", "Note": "备注"}
	head, _ := x.NewStyle(`{"font":{"bold":true}}`)
	for k, v := range xlsx.Map {
		x.SetCellStyle("Sheet1", Axis(k, 0), Axis(k, 0), head)
		if t, ok := m[v]; ok {
			x.SetCellValue("Sheet1", Axis(k, 0), t)
		} else {
			x.SetCellValue("Sheet1", Axis(k, 0), v)
		}
	}
	for i, c := range cs {
		for k, v := range xlsx.Map {
			if _, ok := m[v]; ok {
				switch v {
				case "Name":
					x.SetCellValue("Sheet1", Axis(k, i+1), c.Name)
					break
				case "Phone":
					x.SetCellValue("Sheet1", Axis(k, i+1), c.Phone)
					break
				case "Note":
					x.SetCellValue("Sheet1", Axis(k, i+1), c.Note)
					break
				}
			} else {
				x.SetCellValue("Sheet1", Axis(k, i+1), c.Extend[v])
			}
		}
	}
	return x.Write(c.Response().Writer)
}

// 客户列表
func (w *Web) customers() []Customer {
	if cs, ok := w.cache[CustomerIDPrefix]; ok {
		return cs.([]Customer)
	}
	cs := []Customer{}
	w.Iterator([]byte{CustomerIDPrefix, '-'}, func(key, value []byte) {
		c := Customer{}
		if utils.Decode(value, &c) == nil {
			cs = append(cs, c)
		} else {
			log.Printf("客户信息解析失败 %x \n", key)
		}
	})
	sort.Slice(cs, func(i int, j int) bool {
		return cs[i].Name < cs[j].Name
	})
	w.cache[CustomerIDPrefix] = cs
	return cs
}

// 客户创建
func (w *Web) customerPost(c echo.Context) error {
	delete(w.cache, CustomerIDPrefix)
	cu := Customer{}
	return w.ObjPost(c, &cu, CustomerIDPrefix, func() error { return w.Bind(c, &cu) }, func() error { return w.addTags("tag-C", cu.Tags) })
}

// 客户修改
func (w *Web) customerPut(c echo.Context) error {
	delete(w.cache, CustomerIDPrefix)
	cu := Customer{}
	return w.ObjPut(c, &cu, CustomerIDPrefix, func() error { return w.Bind(c, &cu) }, func() error { return w.addTags("tag-C", cu.Tags) })
}

// 删除用户
func (w *Web) customerDelete(c echo.Context) error {
	return w.ObjDeleter(c, CustomerIDPrefix, &Customer{}, func() error {
		delete(w.cache, CustomerIDPrefix)
		return nil
	})
}

// 清除用户
func (w *Web) customersDelete(c echo.Context) error {
	delete(w.cache, CustomerIDPrefix)
	w.Iterator([]byte{CustomerIDPrefix, '-'}, func(key, value []byte) {
		w.Delete(key)
	})
	return c.JSON(http.StatusOK, "清除完毕")
}

// 客户信息上传
func (w *Web) customersFile(c echo.Context) error {
	delete(w.cache, CustomerIDPrefix)
	xid := new(utils.ID)
	err := xid.Parse(c.Request().Header.Get("xlsx"))
	if err != nil {
		return err
	}

	xlsx := Xlsx{}
	w.Get(xid[:], &xlsx)
	file, err := w.saveTemp(c)
	if err != nil {
		return err
	}
	cs, err := readXlsx(file, xlsx.Map)
	// log.Printf("size: %d %s\n", len(cs), promap)
	if err == nil {
		for _, c := range cs {
			c.BeforePost(CustomerIDPrefix)
			w.addTags("tag-C", c.Tags)
			w.Put(c.ID[:], c)
		}
		os.Remove(file)
		w.customersMerge()
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusInternalServerError, err.Error())
}

// customersMerge 客户信息合并
func (w *Web) customersMerge() {
	m := map[string]Customer{}
	for _, c := range w.customers() {
		k := c.Pinyin + "-" + c.Phone
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
	delete(w.cache, CustomerIDPrefix)
}

// Excel 文件读取
func readXlsx(file string, m map[int]string) (cs []Customer, err error) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return
	}
	for _, name := range xlsx.GetSheetMap() {
		rows := xlsx.GetRows(name)
		cs = []Customer{}
		for _, row := range rows {
			c, e := newCustomer(row, m)
			if e == nil {
				cs = append(cs, c)
			}
		}
	}
	return
}

// 新建客户
func newCustomer(row []string, m map[int]string) (c Customer, err error) {
	p, err := utils.Parse(row, m, &c)
	if err != nil {
		return
	}
	c.Extend = p
	if c.Name == "" {
		err = errors.New("姓名为空")
		return
	}
	if c.Name == "姓名" {
		err = errors.New("姓名为姓名")
		return
	}
	return
}
