package rich

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// Tag 标签
type Tag struct {
	Obj
	Color string          `json:"color"` // 颜色
	Type  map[string]bool `json:"type"`  // 分类
	Nums  map[string]int  `json:"nums"`  // 数量
}

// TagKeys 标签键值
var TagKeys = []string{
	"tag-C", // 客户标签数据
	"tag-I", // 商品标签数据
}

// 标签路由
func (w *Web) tagRoute(c *echo.Group) {
	// 标签列表
	c.GET("", w.tagsGet)
	// 标签创建
	c.POST("", func(c echo.Context) error {
		t := Tag{}
		return w.ObjPost(c, &t, TagIDPrefix, func() error {
			for _, tag := range w.tags() {
				if t.Name == tag.Name {
					return errors.New("标签名称重复:" + t.Name)
				}
			}
			return nil
		})
	})
	// 标签修改
	c.PUT("/:id", func(c echo.Context) error {
		t := Tag{}
		return w.ObjPut(c, &t, TagIDPrefix, func() error {
			for _, tag := range w.tags() {
				if t.Name == tag.Name && t.ID != tag.ID {
					return errors.New("标签名称重复:" + t.Name)
				}
			}
			return nil
		})
	})
	// 标签删除
	c.DELETE("/:id", func(c echo.Context) error {
		return w.ObjDelete(c, TagIDPrefix)
	})
}

// 标签查询
func (w *Web) tagsGet(c echo.Context) error {
	// 数据库标签使用情况
	m := map[string]map[string]int{}
	for _, key := range TagKeys {
		m[key] = map[string]int{}
		w.Get([]byte(key), m[key])
	}
	ret := w.tags()
	for _, t := range ret {
		t.Nums = map[string]int{}
		// 标签总使用次数
		for _, key := range TagKeys {
			if v, ok := m[key]; ok {
				t.Nums[key], _ = v[t.Name]
			}
		}
	}
	// 搜索
	w.ObjSearch(c, &ret)
	return c.JSON(http.StatusOK, ret)
}

// 标签列表
func (w *Web) tags() []Tag {
	ts := []Tag{}
	w.Iterator([]byte{TagIDPrefix, '-'}, func(key, value []byte) {
		t := Tag{}
		if goutils.Decode(value, &t) == nil {
			ts = append(ts, t)
		} else {
			log.Printf("标签解析失败 %x \n", key)
		}
	})
	sort.Slice(ts, func(i int, j int) bool {
		return ts[i].Name > ts[j].Name
	})
	return ts
}
