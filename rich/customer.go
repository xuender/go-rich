package rich

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Customer 客户
type Customer struct {
	Obj
	Phone  string            `json:"phone,omitempty"`  // 电话
	Trades []goutils.ID      `json:"trades,omitempty"` // 消费记录
	Extend map[string]string `json:"extend"`           // 扩展属性
}

// 客户路由
func (w *Web) customerRoute(c *echo.Group) {
	c.GET("", w.customersGet)          // 客户列表
	c.POST("", w.customerPost)         // 客户创建
	c.PUT("/:id", w.customerPut)       // 客户修改
	c.DELETE("/:id", w.customerDelete) // 删除客户
	c.DELETE("", w.customersDelete)    // 清除客户
	c.POST("/file", w.customersFile)   // 上传客户文件
}

// 获取全部客户
func (w *Web) customersGet(c echo.Context) error {
	customers := w.customers()
	w.ObjSearch(c, &customers)
	return c.JSON(http.StatusOK, customers)
}

// 客户列表
func (w *Web) customers() []Customer {
	cs := []Customer{}
	w.Iterator([]byte{CustomerIDPrefix, '-'}, func(key, value []byte) {
		c := Customer{}
		if goutils.Decode(value, &c) == nil {
			cs = append(cs, c)
		} else {
			log.Printf("客户信息解析失败 %x \n", key)
		}
	})
	return cs
}

// 客户创建
func (w *Web) customerPost(c echo.Context) error {
	return w.ObjPost(c, &Customer{}, CustomerIDPrefix, func() error { return nil })
}

// 客户修改
func (w *Web) customerPut(c echo.Context) error {
	return w.ObjPut(c, &Customer{}, CustomerIDPrefix, func() error { return nil })
}

// 删除用户
func (w *Web) customerDelete(c echo.Context) error {
	return w.ObjDelete(c, CustomerIDPrefix)
}

// 清除用户
func (w *Web) customersDelete(c echo.Context) error {
	w.Iterator([]byte{CustomerIDPrefix, '-'}, func(key, value []byte) {
		w.Delete(key)
	})
	return c.JSON(http.StatusOK, "清除完毕")
}

// 客户信息上传
func (w *Web) customersFile(c echo.Context) error {
	xid := new(goutils.ID)
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
			// log.Println(c.Name)
			w.Put(c.ID[:], c)
		}
		os.Remove(file)
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusInternalServerError, err.Error())
}

// Excel 文件读取
func readXlsx(file string, m map[int]string) (cs []Customer, err error) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return
	}
	rows := xlsx.GetRows(xlsx.GetSheetName(1))
	cs = []Customer{}
	for _, row := range rows {
		c, e := newCustomer(row, m)
		if e == nil {
			cs = append(cs, c)
		}
	}
	return
}

// 新建客户
func newCustomer(row []string, m map[int]string) (c Customer, err error) {
	p, err := goutils.Parse(row, m, &c)
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
	c.ID = goutils.NewID(CustomerIDPrefix)
	c.Pinyin = py(row[0])
	c.Ca = time.Now()
	return
}
