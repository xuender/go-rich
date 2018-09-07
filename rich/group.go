package rich

import (
	"net/http"
	"sort"
	"strings"

	"github.com/labstack/echo"
)

// 客户分组路由
func (w *Web) groupsRoute(c *echo.Group) {
	c.GET("", w.groupsGet)       // 查看所有分组
	c.GET("/:group", w.groupGet) // 根据分组获取用户
}

// 获取客户分组列表
func (w *Web) groupsGet(c echo.Context) error {
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
func (w *Web) groupGet(c echo.Context) error {
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
