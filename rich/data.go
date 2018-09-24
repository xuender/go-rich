package rich

import (
	"errors"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/xuender/goutils"
	validator "gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

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

// Get 数据库数据读取
func (w *Web) Get(key []byte, p interface{}) error {
	data, err := w.db.Get(key, nil)
	if err != nil {
		return err
	}
	return goutils.Decode(data, p)
}

// Put 数据库数据保存
func (w *Web) Put(key []byte, p interface{}) error {
	bs, err := goutils.Encode(p)
	if err != nil {
		return err
	}
	return w.db.Put(key, bs, nil)
}

// Iterator 迭代获取数据
func (w *Web) Iterator(prefix []byte, f func(key, value []byte)) error {
	iter := w.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		f(iter.Key(), iter.Value())
	}
	iter.Release()
	return iter.Error()
}

// Delete 删除
func (w *Web) Delete(key []byte) error {
	return w.db.Delete(key, nil)
}

// Bind 对象绑定
func (w *Web) Bind(c echo.Context, o interface{}) error {
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
