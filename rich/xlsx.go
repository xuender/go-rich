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
	Name string         `json:"name"` // 名称
	Map  map[int]string `json:"map"`  // 列定义
}

// Excel定义路由
func (w *Web) xlsxRoute(c *echo.Group) {
	c.GET("", w.xlsxesGet)         // 列表
	c.POST("", w.xlsxPost)         // 创建
	c.GET("/:id", w.xlsxGet)       // 获取
	c.PUT("/:id", w.xlsxPut)       // 修改
	c.DELETE("/:id", w.xlsxDelete) // 删除
}

// 全部Excel定义获取
func (w *Web) xlsxesGet(c echo.Context) error {
	log.Println("xlsxesGet")
	xs := []Xlsx{}
	w.Iterator([]byte{XlsxIDPrefix, '-'}, func(key, value []byte) {
		x := Xlsx{}
		if goutils.Decode(value, &x) == nil {
			xs = append(xs, x)
		} else {
			log.Printf("Excel定义解析失败 %x \n", key)
		}
	})
	return c.JSON(http.StatusOK, xs)
}

// Excle定义创建
func (w *Web) xlsxPost(c echo.Context) error {
	x := Xlsx{}
	return w.objPost(c, &x, func() error {
		if x.Name == "" {
			return errors.New("名程不能为空")
		}
		x.Init(XlsxIDPrefix)
		w.Put(x.ID[:], x)
		return nil
	})
}

// Excel定义获取
func (w *Web) xlsxGet(c echo.Context) error {
	x := Xlsx{}
	return w.objGet(c, &x)
}

// Excel定义修改
func (w *Web) xlsxPut(c echo.Context) error {
	x := Xlsx{}
	return w.objPut(c, &x)
}

// Excel定义删除
func (w *Web) xlsxDelete(c echo.Context) error {
	return w.objDelete(c, XlsxIDPrefix)
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
