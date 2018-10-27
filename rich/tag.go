package rich

import (
	"errors"
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// Tag 标签
type Tag struct {
	Obj
	Color string          `json:"color"` // 颜色
	Use   map[string]bool `json:"use"`   // 使用
	Nums  map[string]int  `json:"nums"`  // 数量
}

// TagItems 标签组
type TagItems []*Tag

func (t *TagItems) filter(c func(*Tag) bool) TagItems {
	ret := TagItems{}
	for _, tag := range *t {
		if c(tag) {
			ret = append(ret, tag)
		}
	}
	return ret
}

func (t TagItems) sort() {
	sort.Slice(t, func(i int, j int) bool {
		return t[i].Name < t[j].Name
	})
}

// TagKeys 标签键值
var TagKeys = []string{
	"tag-C", // 客户标签数据
	"tag-I", // 商品标签数据
}

// 标签重置
func (w *Web) tagsReset() {
	customers := w.customers(false)
	items := w.items(false)
	for _, t := range w.tags() {
		t.Use[TagKeys[0]] = false
		t.Use[TagKeys[1]] = false
		for _, c := range customers {
			if utils.Includes(c.Tags, t.Name) {
				t.Use[TagKeys[0]] = true
				break
			}
		}
		for _, i := range items {
			if utils.Includes(i.Tags, t.Name) {
				t.Use[TagKeys[1]] = true
				break
			}
		}
		w.Put(t.ID[:], t)
	}
}

func (w *Web) addTags(key string, tags Tags) error {
	for _, n := range tags {
		noHas := true
		for _, t := range w.tags() {
			if t.Name == n {
				t.Use[key] = true
				if err := w.Put(t.ID[:], t); err != nil {
					return err
				}
				noHas = false
				break
			}
		}
		if noHas {
			t := Tag{
				Obj: Obj{Name: n},
				Use: map[string]bool{key: true},
			}
			t.BeforePost(TagIDPrefix)
			if err := w.Put(t.ID[:], t); err != nil {
				return err
			}
		}
	}
	return nil
}

// 标签路由
func (w *Web) tagRoute(c *echo.Group) {
	// 标签列表
	c.GET("", w.tagsGet)
	// 标签类型
	c.GET("/:id", func(c echo.Context) error {
		key := c.Param("id")
		ret := w.tags()
		all := c.QueryParam("all") == "true"
		ret = ret.filter(func(t *Tag) bool { return t.Use[key] && (all || len(t.Name) > 1) })
		return c.JSON(http.StatusOK, ret)
	})
	// 标签创建
	c.POST("", func(c echo.Context) error {
		t := Tag{}
		return w.ObjPost(c, &t, TagIDPrefix, func() error { return w.Bind(c, &t) },
			func() error {
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
		return w.ObjPut(c, &t, TagIDPrefix, func() error { return w.Bind(c, &t) },
			func() error {
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
		return w.ObjDelete(c, TagIDPrefix, func(id utils.ID) error { return nil })
	})
}

// 标签查询
func (w *Web) tagsGet(c echo.Context) error {
	ret := w.tags()
	// 搜索
	w.ObjSearch(c, &ret)
	ret = ret.filter(func(t *Tag) bool { return len(t.Name) > 1 })
	// 数据库标签使用情况
	// m := map[string]map[string]int{}
	// for _, key := range TagKeys {
	// 	m[key] = map[string]int{}
	// 	w.Get([]byte(key), m[key])
	// }
	// for _, t := range ret {
	// 	t.Nums = map[string]int{}
	// 	// 标签总使用次数
	// 	for _, key := range TagKeys {
	// 		if v, ok := m[key]; ok {
	// 			t.Nums[key], _ = v[t.Name]
	// 		}
	// 	}
	// }
	return c.JSON(http.StatusOK, ret)
}

// 标签列表
func (w *Web) tags() TagItems {
	ts := TagItems{}
	w.Iterator([]byte{TagIDPrefix, '-'}, func(key, value []byte) {
		t := &Tag{}
		if utils.Decode(value, t) == nil {
			ts = append(ts, t)
		} else {
			log.Printf("标签解析失败 %x \n", key)
		}
	})
	ts.sort()
	return ts
}
