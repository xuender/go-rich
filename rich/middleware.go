package rich

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// middlewareJWT returns a JWT auth middleware.
func middlewareJWT(w *Web, signingMethod string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			err := errors.New("身份认证错误")
			if auth == "" {
				return err
			}
			auth = auth[7:]
			log.Printf("[%s]\n", auth)
			token := new(jwt.Token)
			token, err = jwt.Parse(auth, func(t *jwt.Token) (interface{}, error) {
				if signingMethod != t.Method.Alg() {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				// 用户ID
				id := t.Claims.(jwt.MapClaims)["id"]
				if sign, ok := w.cache.Get(id); ok {
					return sign, nil
				}
				// 设置缓存
				for _, u := range w.users() {
					if u.ID.String() == id {
						w.cache.Put(id, u.Pass)
						return u.Pass, nil
					}
				}
				return nil, errors.New("用户未找到")
			})
			if err == nil && token.Valid {
				c.Set("user", token)
				return next(c)
			}
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "invalid or expired jwt",
				Internal: err,
			}
		}
	}
}
