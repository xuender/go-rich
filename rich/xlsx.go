package rich

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Xlsx Excel定义
type Xlsx struct {
	Obj
	Map map[int]string `json:"map"` // 列定义
}

// Excel定义路由
func (w *Web) xlsxRoute(c *echo.Group) {
	c.GET("", func(c echo.Context) error { return c.JSON(http.StatusOK, w.xlsxes()) })   // 列表
	c.GET("/:id", func(c echo.Context) error { return w.ObjGet(c, &Xlsx{}) })            // 获取
	c.POST("", w.xlsxPost)                                                               // 创建
	c.PUT("/:id", w.xlsxPut)                                                             // 修改
	c.DELETE("/:id", func(c echo.Context) error { return w.ObjDelete(c, XlsxIDPrefix) }) // 删除
}

func (w *Web) xlsxes() []Xlsx {
	xs := []Xlsx{}
	w.Iterator([]byte{XlsxIDPrefix, '-'}, func(key, value []byte) {
		x := Xlsx{}
		if goutils.Decode(value, &x) == nil {
			xs = append(xs, x)
		} else {
			log.Printf("Excel定义解析失败 %x \n", key)
		}
	})
	return xs
}

// Excle定义创建
func (w *Web) xlsxPost(c echo.Context) error {
	x := Xlsx{}
	return w.ObjPost(c, &x, XlsxIDPrefix, func() error { return w.Bind(c, &x) },
		func() error {
			for _, o := range w.xlsxes() {
				if o.Name == x.Name {
					return errors.New("名称重复")
				}
			}
			return nil
		})
}

// Excel定义修改
func (w *Web) xlsxPut(c echo.Context) error {
	x := Xlsx{}
	return w.ObjPut(c, &x, XlsxIDPrefix, func() error { return w.Bind(c, &x) },
		func() error {
			for _, o := range w.xlsxes() {
				if o.Name == x.Name && o.ID != x.ID {
					return errors.New("名称重复")
				}
			}
			return nil
		})
}

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
