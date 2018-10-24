package rich

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// Obj 基础对象
type Obj struct {
	ID     utils.ID  `json:"id"`                              // 主键
	Name   string    `json:"name" validate:"required,lte=50"` // 名称 必填 最长50
	Pinyin string    `json:"pinyin"`                          // 拼音
	Note   string    `json:"note" validate:"lte=100"`         // 备注 最长100
	Ca     time.Time `json:"ca"`                              // 创建时间
	Da     time.Time `json:"da,omitempty"`                    // 删除时间
}

// Delete 删除对象
func (o *Obj) Delete() {
	o.Da = time.Now()
}

// IsDelete 对象是否删除
func (o *Obj) IsDelete() bool {
	return o.Da.Year() > 1
}

// Deleter 删除器
type Deleter interface {
	IsDelete() bool
	Delete()
}

// Binder 对象绑定
type Binder interface {
	// Bind 绑定数据
	Bind(c echo.Context) error
}

// Puter 对象修改
type Puter interface {
	// BeforePut 修改前
	BeforePut(id utils.ID)
}

// Poster 新对象操作
type Poster interface {
	// BeforePost 保存前
	BeforePost(key byte) utils.ID
}

// Matcher 对象比较
type Matcher interface {
	Match(txt string) bool
}

// Includeser 包含
type Includeser interface {
	Includes(tags []string) bool
}

// BeforePut 修改对象
func (o *Obj) BeforePut(id utils.ID) {
	o.Name = strings.Trim(o.Name, " ")
	o.ID = id
	o.Pinyin = PY(o.Name)
}

// BeforePost 对象新建前
func (o *Obj) BeforePost(key byte) utils.ID {
	o.Name = strings.Trim(o.Name, " ")
	o.ID = utils.NewID(key)
	o.Ca = time.Now()
	o.Pinyin = PY(o.Name)
	return o.ID
}

// ObjPut 对象修改
func (w *Web) ObjPut(c echo.Context, p Puter, key byte, bind func() error, check func() error) error {
	idstr := c.Param("id")
	if idstr == "" {
		return errors.New("id为空")
	}
	id := new(utils.ID)
	err := id.Parse(idstr)
	if err != nil {
		return err
	}
	if id[0] != key {
		return errors.New("前缀错误")
	}
	w.Get(id[:], p)
	if err := bind(); err != nil {
		return err
	}
	p.BeforePut(*id)
	if check != nil {
		if err := check(); err != nil {
			return err
		}
	}
	w.Put(id[:], p)
	return c.JSON(http.StatusOK, p)
}

// ObjPost 对象新增
func (w *Web) ObjPost(c echo.Context, p Poster, key byte, bind func() error, check func() error) error {
	if err := bind(); err != nil {
		return err
	}
	id := p.BeforePost(key)
	if check != nil {
		if err := check(); err != nil {
			return err
		}
	}
	w.Put(id[:], p)
	return c.JSON(http.StatusCreated, p)
}

// ObjLoad 对象加载
func (w *Web) ObjLoad(c echo.Context, o interface{}) error {
	id := new(utils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	return w.Get(id[:], o)
}

// ObjSearch 对象搜索
func (w *Web) ObjSearch(c echo.Context, i interface{}) (interface{}, error) {
	txt := c.QueryParam("search")
	if txt != "" {
		return utils.Filter(i, func(m Matcher) bool {
			return m.Match(txt)
		})
	}
	return i, nil
}

// ObjSelect 对象选择
func (w *Web) ObjSelect(c echo.Context, i interface{}) (interface{}, error) {
	str := c.QueryParam("tags")
	if str != "" {
		tags := []string{}
		json.Unmarshal([]byte(str), &tags)
		return utils.Filter(i, func(m Includeser) bool {
			return m.Includes(tags)
		})
	}
	return i, nil
}

// ObjPaging 分页
func (w *Web) ObjPaging(c echo.Context, i interface{}) error {
	rv := reflect.ValueOf(i)
	err := errors.New("分页对象必须是slice或array")
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return err
	}
	s := c.QueryParam("size")
	p := c.QueryParam("page")
	size := 20
	page := 0
	if s != "" {
		size, err = strconv.Atoi(s)
		if err != nil {
			return err
		}
	}
	if p != "" {
		page, err = strconv.Atoi(p)
		if err != nil {
			return err
		}
	}
	start := page * size
	if start > rv.Len() {
    return c.JSON(http.StatusOK, Paging{Data: []int{}, Total: rv.Len()})
	}
	end := start + size
	if end > rv.Len() {
		end = rv.Len()
	}
	t := rv.Type().Elem()
	out := reflect.MakeSlice(reflect.SliceOf(t), 0, end-start)
	for i := start; i < end; i++ {
		out = reflect.Append(out, rv.Index(i))
	}
	return c.JSON(http.StatusOK, Paging{Data: out.Interface(), Total: rv.Len()})
}

// ObjGet 对象获取
func (w *Web) ObjGet(c echo.Context, o interface{}) error {
	if err := w.ObjLoad(c, o); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, o)
}

// ObjDelete 对象删除
func (w *Web) ObjDelete(c echo.Context, key byte, check func(id utils.ID) error) error {
	id := utils.ID{}
	if err := id.Parse(c.Param("id")); err != nil {
		return err
	}
	if id[0] != key {
		return errors.New("前缀错误")
	}
	if err := check(id); err != nil {
		return err
	}
	w.Delete(id[:])
	return c.JSON(http.StatusNoContent, nil)
}

// ObjDeleter 对象删除器
func (w *Web) ObjDeleter(c echo.Context, key byte, obj Deleter, check func() error) error {
	id := utils.ID{}
	if err := id.Parse(c.Param("id")); err != nil {
		return err
	}
	if id[0] != key {
		return errors.New("前缀错误")
	}
	w.Get(id[:], obj)
	if obj.IsDelete() {
		return c.JSON(http.StatusOK, nil)
	}
	if err := check(); err != nil {
		return err
	}
	obj.Delete()
	w.Put(id[:], obj)
	return c.JSON(http.StatusNoContent, nil)
}

// Match 查找匹配
func (o Obj) Match(txt string) bool {
	if strings.Contains(o.Name, txt) ||
		strings.Contains(o.Pinyin, strings.ToLower(txt)) {
		return true
	}
	s := []byte{}
	for _, p := range strings.Split(o.Pinyin, " ") {
		if len(p) > 0 {
			s = append(s, p[0])
		}
	}
	return strings.Contains(string(s), strings.ToLower(txt))
}
