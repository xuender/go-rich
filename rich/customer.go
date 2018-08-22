package rich

import (
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/labstack/echo"
	pinyin "github.com/mozillazg/go-pinyin"
	"github.com/xuender/goutils"
)

type Customer struct {
	Id         goutils.ID        `json:"id"`                   // 主键
	Name       string            `json:"name"`                 // 姓名
	Pinyin     string            `json:"pinyin"`               // 拼音
	Phone      string            `json:"phone,omitempty"`      // 电话
	Ca         time.Time         `json:"ca"`                   // 创建时间
	Trades     []goutils.ID      `json:"trades,omitempty"`     // 消费记录
	Note       string            `json:"note,omitempty"`       // 备注
	Properties map[string]string `json:"properties,omitempty"` // 扩展属性
}

func (w *Web) customerRoute(c *echo.Group) {
	// excel 格式定义
	c.POST("/promap", w.postPromap)
	c.GET("/promap", w.getPromap)
	c.POST("/up", w.cup)
	// TODO
	c.GET("/all", w.customerAll)
}
func (w *Web) customerAll(c echo.Context) error {
	cs := []Customer{}
	w.Iterator([]byte{CustomerIdPrefix, '-'}, func(data []byte) {
		c := Customer{}
		goutils.Decode(data, &c)
		cs = append(cs, c)
	})
	return c.JSON(http.StatusOK, cs)
}

var customerPromapKey = []byte("CP-MAP")

func (w *Web) postPromap(c echo.Context) error {
	m := make(map[int]string)
	if err := c.Bind(&m); err != nil {
		return err
	}
	w.Put(customerPromapKey, m)
	return c.JSON(http.StatusOK, m)
}
func (w *Web) getPromap(c echo.Context) error {
	m := make(map[int]string)
	w.Get(customerPromapKey, &m)
	return c.JSON(http.StatusOK, m)
}
func (w *Web) saveTemp(c echo.Context) (string, error) {
	// 来源
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// 目的
	mkdir(w.Temp)
	f := w.Temp + string(os.PathSeparator) + file.Filename
	dst, err := os.Create(f)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// 复制
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return f, nil
}
func (w *Web) cup(c echo.Context) error {
	promap := make(map[int]string)
	err := w.Get(customerPromapKey, &promap)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	file, err := w.saveTemp(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	cs, err := loadXlsx(file, promap)
	if err == nil {
		for _, c := range cs {
			w.Put(c.Id[:], c)
		}
		os.Remove(file)
		return c.String(http.StatusOK, "ok")
	}
	return c.String(http.StatusInternalServerError, err.Error())
}
func loadXlsx(file string, promap map[int]string) (cs []Customer, err error) {
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
func NewCustomer(row []string, customerProMap map[int]string) (c Customer, err error) {
	err = goutils.Parse(row, customerProMap, &c)
	if err != nil {
		return
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
	return
}
