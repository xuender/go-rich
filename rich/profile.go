package rich

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/xuender/goutils"
)

// 账户路由
func (w *Web) profileRoute(c *echo.Group) {
	c.GET("", w.profileGet)     // 账户信息
	c.PATCH("", w.profilePatch) // 修改密码
}

// profile 获取当前账户
func (w *Web) profile(c echo.Context) (u User, err error) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := new(goutils.ID)
	if err = id.Parse(claims["id"].(string)); err != nil {
		return
	}
	w.Get(id[:], &u)
	return
}

// profileGet 当前账户信息
func (w *Web) profileGet(c echo.Context) error {
	u, err := w.profile(c)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// profilePatch 修改密码
func (w *Web) profilePatch(c echo.Context) error {
	old := c.QueryParam("old")
	if old == "" {
		return errors.New("原密码不能为空")
	}
	pass := c.QueryParam("pass")
	if pass == "" {
		return errors.New("新密码不能为空")
	}
	u, err := w.profile(c)
	if err != nil {
		return err
	}
	if bytes.Equal(u.Pass, Pass(old)) {
		return errors.New("原密码错误")
	}
	u.Pass = Pass(pass)
	w.Put(u.ID[:], u)
	// 删除用户密钥缓存
	w.cache.Remove(u.ID.String())
	t, err := w.Signed(u.ID.String(), u.Pass)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, t)
}
