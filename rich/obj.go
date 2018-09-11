package rich

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Obj 基础对象
type Obj struct {
	ID   goutils.ID `json:"id"`   // 主键
	Note string     `json:"note"` // 备注
	Ca   time.Time  `json:"ca"`   // 创建时间
}

// IDModifier ID修改接口
type IDModifier interface {
	ModifyID(id goutils.ID)
}

// ModifyID ID修改
func (o *Obj) ModifyID(id goutils.ID) {
	o.ID = id
}

// Init 对象初始化
func (o *Obj) Init(key byte) {
	o.ID = goutils.NewId(key)
	o.Ca = time.Now()
}

// 对象修改
func (w *Web) objPut(c echo.Context, o IDModifier) error {
	id := new(goutils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	w.Get(id[:], o)
	if err := c.Bind(o); err != nil {
		return err
	}
	o.ModifyID(*id)
	w.Put(id[:], o)
	return c.JSON(http.StatusOK, o)
}

// 对象创建
func (w *Web) objPost(c echo.Context, o interface{}, check func() error) error {
	if err := c.Bind(o); err != nil {
		return err
	}
	if err := check(); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, o)
}

// 对象获取
func (w *Web) objGet(c echo.Context, o interface{}) error {
	id := new(goutils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	w.Get(id[:], o)
	return c.JSON(http.StatusOK, o)
}

// 对象删除
func (w *Web) objDelete(c echo.Context, key byte) error {
	id := new(goutils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	if id[0] != key {
		return errors.New("前缀错误")
	}
	w.Delete(id[:])
	return c.JSON(http.StatusOK, nil)
}
