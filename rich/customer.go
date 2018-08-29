package rich

import (
	"errors"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	pinyin "github.com/mozillazg/go-pinyin"
	"github.com/xuender/goutils"
)

type Customer struct {
	Id     goutils.ID   `json:"id"`               // 主键
	Name   string       `json:"name"`             // 姓名
	Pinyin string       `json:"pinyin"`           // 拼音
	Phone  string       `json:"phone,omitempty"`  // 电话
	Ca     time.Time    `json:"ca"`               // 创建时间
	Trades []goutils.ID `json:"trades,omitempty"` // 消费记录
	Note   string       `json:"note,omitempty"`   // 备注
	Extend []Ext        `json:"extend,omitempty"` // 扩展属性
}

func (w *Web) customerRoute(c *echo.Group) {
	// excel 格式定义
	c.PUT("/ext", w.postExt)
	c.GET("/ext", w.getExt)
	c.POST("/up", w.customerUp)
	c.GET("/all", w.customerAll)
	c.DELETE("/all", w.cleanCustomer)
	c.GET("/groups", w.customerGroups)
	c.GET("/g/:group", w.customerGroup)
	c.DELETE("/:id", w.deleteCustomer)
}

// 清除用户
func (w *Web) cleanCustomer(c echo.Context) error {
	for _, c := range w.customers() {
		w.Delete(c.Id[:])
	}
	return c.JSON(http.StatusOK, "清除完毕")
}

// 删除用户
func (w *Web) deleteCustomer(c echo.Context) error {
	id := new(goutils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	w.Delete(id[:])
	return c.JSON(http.StatusOK, "删除完毕")
}

// 获取客户分组列表
func (w *Web) customerGroups(c echo.Context) error {
	m := make(map[string]bool)
	for _, c := range w.customers() {
		rs := []rune(c.Pinyin)
		if len(rs) > 0 {
			m[strings.ToUpper(string(rs[0:1]))] = true
		}
	}
	ret := []string{}
	for item, _ := range m {
		ret = append(ret, item)
	}
	sort.Strings(ret)
	return c.JSON(http.StatusOK, ret)
}

// 获取客户分组
func (w *Web) customerGroup(c echo.Context) error {
	group := strings.ToUpper(c.Param("group"))
	ret := []Customer{}
	for _, c := range w.customers() {
		rs := []rune(c.Pinyin)
		if len(rs) > 0 {
			if group == strings.ToUpper(string(rs[0:1])) {
				ret = append(ret, c)
			}
		}
	}
	sort.Slice(ret, func(i int, j int) bool {
		return ret[i].Pinyin > ret[j].Pinyin
	})
	return c.JSON(http.StatusOK, ret)
}

// 客户列表
func (w *Web) customers() []Customer {
	cs := []Customer{}
	w.Iterator([]byte{CustomerIdPrefix, '-'}, func(data []byte) {
		c := Customer{}
		goutils.Decode(data, &c)
		cs = append(cs, c)
	})
	return cs
}

// 获取全部客户
func (w *Web) customerAll(c echo.Context) error {
	return c.JSON(http.StatusOK, w.customers())
}

var extCustomerKey = []byte("EXT-C")

// 设置 excel 定义
func (w *Web) postExt(c echo.Context) error {
	m := []Ext{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	w.Put(extCustomerKey, m)
	return c.JSON(http.StatusOK, m)
}

// 获取 excel 定义
func (w *Web) getExt(c echo.Context) error {
	m := []Ext{}
	w.Get(extCustomerKey, &m)
	return c.JSON(http.StatusOK, m)
}

// 客户信息上传
func (w *Web) customerUp(c echo.Context) error {
	promap := []Ext{}
	err := w.Get(extCustomerKey, &promap)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Excel尚未定义")
	}
	file, err := w.saveTemp(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	cs, err := readXlsx(file, promap)
	// log.Printf("size: %d %s\n", len(cs), promap)
	if err == nil {
		for _, c := range cs {
			// log.Println(c.Name)
			w.Put(c.Id[:], c)
		}
		os.Remove(file)
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusInternalServerError, err.Error())
}

// Excel 文件读取
func readXlsx(file string, promap []Ext) (cs []Customer, err error) {
	xlsx, err := excelize.OpenFile(file)
	if err != nil {
		return
	}
	rows := xlsx.GetRows(xlsx.GetSheetName(1))
	cs = []Customer{}
	for _, row := range rows {
		c, e := NewCustomer(row, promap)
		if e == nil {
			cs = append(cs, c)
		}
	}
	return
}

// var customerProMap map[int]string
var args = pinyin.NewArgs()

func init() {
	// customerProMap = make(map[int]string)
	// customerProMap[0] = "Name"
	// customerProMap[1] = "R球镜"
	// customerProMap[2] = "R柱镜"
	// customerProMap[3] = "R轴位"
	// customerProMap[4] = "L球镜"
	// customerProMap[5] = "L柱镜"
	// customerProMap[6] = "L轴位"
	// customerProMap[7] = "轴位"
	// customerProMap[8] = "镜架"
	// customerProMap[9] = "镜片"
	// customerProMap[10] = "金额"
	// customerProMap[11] = "Phone"
	// customerProMap[12] = "Note"
	args.Separator = " "
}
func NewCustomer(row []string, ext []Ext) (c Customer, err error) {
	customerProMap := make(map[int]string)
	for _, e := range ext {
		i, erro := strconv.Atoi(e.Value)
		if erro == nil {
			customerProMap[i] = e.Key
		}
	}
	p, err := goutils.Parse(row, customerProMap, &c)
	if err != nil {
		return
	}
	for _, e := range ext {
		v, ok := p[e.Key]
		if ok {
			c.Extend = append(c.Extend, Ext{Key: e.Key, Value: v})
		}
	}
	if c.Name == "" {
		err = errors.New("姓名为空")
		return
	}
	if c.Name == "姓名" {
		err = errors.New("姓名为姓名")
		return
	}
	c.Id = goutils.NewId(CustomerIdPrefix)
	c.Pinyin = pinyin.Slug(row[0], args)
	c.Ca = time.Now()
	return
}
