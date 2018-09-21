package rich

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/xuender/goutils"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

// Obj 基础对象
type Obj struct {
	ID     goutils.ID `json:"id"`                              // 主键
	Name   string     `json:"name" validate:"required,lte=50"` // 名称 必填 最长50
	Pinyin string     `json:"pinyin"`                          // 拼音
	Note   string     `json:"note" validate:"lte=100"`         // 备注 最长100
	Ca     time.Time  `json:"ca"`                              // 创建时间
}

// Binder 对象绑定
type Binder interface {
	// Bind 绑定数据
	Bind(c echo.Context) error
}

// Puter 对象修改
type Puter interface {
	Binder
	// BeforePut 修改前
	BeforePut(id goutils.ID)
}

// Poster 新对象操作
type Poster interface {
	Binder
	// BeforePost 保存前
	BeforePost(key byte) goutils.ID
}

// Matcher 对象比较
type Matcher interface {
	Match(txt string) bool
}

// BeforePut 修改对象
func (o *Obj) BeforePut(id goutils.ID) {
	o.ID = id
}

// BeforePost 对象新建前
func (o *Obj) BeforePost(key byte) goutils.ID {
	o.ID = goutils.NewID(key)
	o.Ca = time.Now()
	o.Pinyin = py(o.Name)
	return o.ID
}

// Bind 对象绑定
func (o *Obj) Bind(c echo.Context) error {
	// 数据绑定
	if err := c.Bind(o); err != nil {
		return err
	}
	// 校验
	if err := validate.Struct(o); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			ret := ""
			for _, v := range errs.Translate(trans) {
				ret = v + "\n"
			}
			return errors.New(ret)
		}
		return errors.New("校验失败")
	}
	return nil
}

// ObjPut 对象修改
func (w *Web) ObjPut(c echo.Context, p Puter, key byte, check func() error) error {
	id := new(goutils.ID)
	err := id.Parse(c.Param("id"))
	if err != nil {
		return err
	}
	if id[0] != key {
		return errors.New("前缀错误")
	}
	w.Get(id[:], p)
	if err := p.Bind(c); err != nil {
		return err
	}
	if err := check(); err != nil {
		return err
	}
	p.BeforePut(*id)
	w.Put(id[:], p)
	return c.JSON(http.StatusOK, p)
}

// ObjPost 对象新增
func (w *Web) ObjPost(c echo.Context, p Poster, key byte, check func() error) error {
	if err := p.Bind(c); err != nil {
		return err
	}
	if err := check(); err != nil {
		return err
	}
	id := p.BeforePost(key)
	w.Put(id[:], p)
	return c.JSON(http.StatusOK, p)
}

// ObjLoad 对象加载
func (w *Web) ObjLoad(c echo.Context, o interface{}) error {
	id := new(goutils.ID)
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
		return goutils.Filter(i, func(m Matcher) bool {
			return m.Match(txt)
		})
	}
	return i, nil
}

// ObjGet 对象获取
func (w *Web) ObjGet(c echo.Context, o interface{}) error {
	if err := w.ObjLoad(c, o); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, o)
}

// ObjDelete 对象删除
func (w *Web) ObjDelete(c echo.Context, key byte) error {
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

// Match 查找匹配
func (o Obj) Match(txt string) bool {
	if strings.Contains(o.Name, txt) ||
		strings.Contains(o.Pinyin, strings.ToLower(txt)) ||
		strings.Contains(o.Pinyin, py(txt)) {
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

// Validate 校验
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	// 多国语言
	zhl := zh.New()
	uni = ut.New(zhl, zhl)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
}
