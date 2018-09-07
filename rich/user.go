package rich

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// 用户
type User struct {
	Id    goutils.ID `json:"id"`    // 主键
	Nick  string     `json:"nick"`  // 昵称
	Phone string     `json:"phone"` // 电话
	Pass  string     `json:"-"`     // 密码
	Ca    time.Time  `json:"ca"`    // 创建时间
}

// 用户路由
func (w *Web) userRoute(c *echo.Group) {
	c.GET("", w.usersGet) // 全部用户
}

// 获取全部用户
func (w *Web) usersGet(c echo.Context) error {
	return c.JSON(http.StatusOK, w.users())
}

func (w *Web) users() []User {
	us := []User{}
	w.Iterator([]byte{UserIdPrefix, '-'}, func(key, data []byte) {
		u := User{}
		goutils.Decode(data, &u)
		us = append(us, u)
	})
	return us
}

func (w *Web) UserInit() {
	us := w.users()
	if len(us) == 0 {
		u := User{
			Id:    goutils.NewId(UserIdPrefix),
			Nick:  "admin",
			Phone: "admin",
			Pass:  "6181",
			Ca:    time.Now(),
		}
		w.Put(u.Id[:], u)
	}
}
