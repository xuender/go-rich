package rich

import (
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// User 用户
type User struct {
	Obj
	Nick   string            `json:"nick"`   // 昵称
	Phone  string            `json:"phone"`  // 电话
	Pass   string            `json:"-"`      // 密码
	Extend map[string]string `json:"extend"` // 扩展属性
}

// 用户路由
func (w *Web) userRoute(c *echo.Group) {
	c.GET("", w.usersGet)          // 全部用户
	c.POST("", w.userPost)         // 创建
	c.GET("/:id", w.userGet)       // 获取
	c.PUT("/:id", w.userPut)       // 修改
	c.DELETE("/:id", w.userDelete) // 删除
}

// 获取全部用户
func (w *Web) usersGet(c echo.Context) error {
	return c.JSON(http.StatusOK, w.users())
}

func (w *Web) users() []User {
	us := []User{}
	w.Iterator([]byte{UserIDPrefix, '-'}, func(key, data []byte) {
		u := User{}
		goutils.Decode(data, &u)
		us = append(us, u)
	})
	return us
}

// 用户创建
func (w *Web) userPost(c echo.Context) error {
	u := User{}
	return w.objPost(c, &u, func() error {
		if u.Nick == "" {
			return errors.New("昵称不能为空")
		}
		u.Init(UserIDPrefix)
		w.Put(u.ID[:], u)
		return nil
	})
}

// 用户获取
func (w *Web) userGet(c echo.Context) error {
	u := User{}
	return w.objGet(c, &u)
}

// 用户修改
func (w *Web) userPut(c echo.Context) error {
	u := User{}
	return w.objPut(c, &u)
}

// 用户删除
func (w *Web) userDelete(c echo.Context) error {
	return w.objDelete(c, UserIDPrefix)
}

// UserInit 用户初始化
func (w *Web) UserInit() {
	us := w.users()
	if len(us) == 0 {
		u := User{
			Obj: Obj{
				ID: goutils.NewId(UserIDPrefix),
				Ca: time.Now(),
			},
			Nick:  "admin",
			Phone: "admin",
			Pass:  "6181",
		}
		w.Put(u.ID[:], u)
	}
}
