package rich

import (
	"errors"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/go-utils"
)

// User 用户
type User struct {
	Obj
	Phone  string            `json:"phone"`  // 手机
	Pass   []byte            `json:"-"`      // 密码
	Admin  bool              `json:"admin"`  // 管理员
	Extend map[string]string `json:"extend"` // 扩展属性
}

// 用户路由
func (w *Web) userRoute(c *echo.Group) {
	c.GET("", w.usersGet)            // 全部用户
	c.POST("", w.userPost)           // 创建
	c.GET("/:id", w.userGet)         // 获取
	c.PUT("/:id", w.userPut)         // 修改
	c.DELETE("/:id", w.userDelete)   // 删除
	c.PATCH("/:id/pass", w.userPass) // 修改密码
}

// Match 匹配
func (u User) Match(txt string) bool {
	return u.Obj.Match(txt) || strings.Contains(u.Phone, txt)
}

// 用户修改密码
func (w *Web) userPass(c echo.Context) error {
	pass := c.QueryParam("pass")
	if pass == "" {
		return errors.New("密码不能为空")
	}
	u := &User{}
	if err := w.ObjLoad(c, u); err != nil {
		return err
	}
	u.Pass = Pass(pass)
	w.Put(u.ID[:], u)
	// 删除用户密钥缓存
	delete(w.cache, u.ID.String())
	return c.JSON(http.StatusOK, u)
}

// 获取全部用户
func (w *Web) usersGet(c echo.Context) error {
	ret := w.users()
	w.ObjSearch(c, &ret)
	return c.JSON(http.StatusOK, ret)
}

func (w *Web) users() []User {
	us := []User{}
	w.Iterator([]byte{UserIDPrefix, '-'}, func(key, data []byte) {
		u := User{}
		utils.Decode(data, &u)
		us = append(us, u)
	})
	sort.Slice(us, func(i int, j int) bool {
		return us[i].Name < us[j].Name
	})
	return us
}

// 用户创建
func (w *Web) userPost(c echo.Context) error {
	u := User{}
	return w.ObjPost(c, &u, UserIDPrefix, func() error { return w.Bind(c, &u) }, func() error {
		for _, o := range w.users() {
			if o.Name == u.Name {
				return errors.New("姓名重复")
			}
			if o.Phone == u.Phone {
				return errors.New("手机重复")
			}
		}
		return nil
	})
}

// 用户获取
func (w *Web) userGet(c echo.Context) error {
	return w.ObjGet(c, &User{})
}

// 用户修改
func (w *Web) userPut(c echo.Context) error {
	u := User{}
	return w.ObjPut(c, &u, UserIDPrefix, func() error { return w.Bind(c, &u) }, func() error {
		users := w.users()
		utils.Filter(&users, func(o User) bool { return !o.ID.Equal(u.ID) })
		for _, o := range users {
			if o.Name == u.Name {
				return errors.New("姓名重复")
			}
			if o.Phone == u.Phone {
				return errors.New("手机重复")
			}
		}
		return nil
	})
}

// 用户删除
func (w *Web) userDelete(c echo.Context) error {
	return w.ObjDelete(c, UserIDPrefix, func(id utils.ID) error { return nil })
}

// UserInit 用户初始化
func (w *Web) UserInit() {
	us := w.users()
	if len(us) == 0 {
		u := User{
			Obj: Obj{
				ID:   utils.NewID(UserIDPrefix),
				Name: "admin",
				Ca:   time.Now(),
			},
			Phone: "13900000000",
			Pass:  Pass(strings.Replace(w.App.Version, ".", "", -1)),
		}
		w.Put(u.ID[:], u)
	}
}
