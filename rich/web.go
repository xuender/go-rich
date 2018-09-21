package rich

import (
	"crypto/rsa"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"../keys"
	"rsc.io/qr"

	static "github.com/Code-Hex/echo-static"
	jwt "github.com/dgrijalva/jwt-go"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/xuender/goutils"
)

// Web 网络服务
type Web struct {
	Port string      // 端口号
	Temp string      // 临时文件目录
	Db   string      // 数据库目录
	Dev  bool        // 开发模式
	db   *leveldb.DB // 数据库
	days Days        // 文件日期列表
}

// DaysKey 文件日期列表主键
var DaysKey = []byte("days")
var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

// Init 初始化.
func (w *Web) Init() error {
	// 日志初始化
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	// 密钥初始化
	signBytes, err := keys.Asset("keys/private.rsa")
	if err != nil {
		return err
	}
	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		return err
	}
	verifyBytes, err := keys.Asset("keys/public.rsa.pub")
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return err
	}
	// 数据库初始化
	w.db, err = leveldb.OpenFile(w.Db, nil)
	if err != nil {
		return err
	}
	// 用户初始化
	w.UserInit()
	// 每日帐目初始化
	w.Get(DaysKey, &w.days)
	return nil
}

// Run 启动服务
func (w *Web) Run() (err error) {
	e := echo.New()
	e.HTTPErrorHandler = w.httpErrorHandler
	// 开发模式
	if w.Dev {
		e.Use(middleware.Recover())
		e.Use(middleware.Logger())
		// 跨域访问
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
			AllowCredentials: true,
		}))
	}
	e.GET("/qr", w.qrcode)   // 二维码访问
	e.GET("/info", w.info)   // 协议等
	e.GET("/login", w.login) // 登录
	api := e.Group("/api")   // API
	// 需要身份认证
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    verifyKey,
		SigningMethod: "RS256",
	}))

	w.customerRoute(api.Group("/customers")) // 客户
	w.groupsRoute(api.Group("/groups"))      // 客户分组
	w.extsRoute(api.Group("/exts"))          // 扩展定义
	w.xlsxRoute(api.Group("/xlsxes"))        // Excel定义
	w.userRoute(api.Group("/users"))         // 用户
	w.profileRoute(api.Group("/profile"))    // 账户
	w.itemRoute(api.Group("/items"))         // 商品
	w.tagRoute(api.Group("/tags"))         // 标签

	// 静态资源
	if w.Dev {
		e.Static("/", "www")
	} else {
		e.Use(static.ServeRoot("/", getAssets("www")))
	}
	// return e.Start(w.Port)
	// HTTP/2.0 启动
	return w.start(e)
}
func (w *Web) httpErrorHandler(err error, c echo.Context) {
	var code = http.StatusInternalServerError

	if !c.Response().Committed {
		if c.Request().Method == echo.HEAD {
			if err := c.NoContent(code); err != nil {
				c.Logger().Error(err)
			}
		} else {
			if es, ok := err.(*echo.HTTPError); ok {
				if err := c.JSON(es.Code, newHTTPError(es)); err != nil {
					c.Logger().Error(err)
				}
			} else {
				if err := c.JSON(code, newHTTPError(err)); err != nil {
					c.Logger().Error(err)
				}
			}
		}
	}
}

// 启动服务
func (w *Web) start(e *echo.Echo) error {
	s := e.TLSServer
	s.TLSConfig = new(tls.Config)
	s.TLSConfig.Certificates = make([]tls.Certificate, 1)
	cert, err := keys.Asset("keys/cert.pem")
	if err != nil {
		return err
	}
	keys, err := keys.Asset("keys/key.pem")
	if err != nil {
		return err
	}
	s.TLSConfig.Certificates[0], err = tls.X509KeyPair(cert, keys)
	if err != nil {
		return err
	}
	s.Addr = w.Port
	if !e.DisableHTTP2 {
		s.TLSConfig.NextProtos = append(s.TLSConfig.NextProtos, "h2")
	}
	return e.StartServer(e.TLSServer)
}

// Close 关闭服务
func (w *Web) Close() {
	w.db.Close()
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

// 静态资源
func getAssets(root string) *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    root,
	}
}

// 登录
func (w *Web) login(c echo.Context) error {
	// 登录信息绑定
	nick := c.QueryParam("nick")
	pass := c.QueryParam("pass")
	if nick == "" {
		return errors.New("昵称或电话不能为空")
	}
	if pass == "" {
		return errors.New("密码不能为空")
	}
	pass = Pass(pass)
	for _, u := range w.users() {
		// 身份认证
		if u.Name == nick || u.Phone == nick || pass == u.Pass {
			// 创建令牌
			token := jwt.New(jwt.SigningMethodRS256)
			// 设置用户信息
			claims := token.Claims.(jwt.MapClaims)
			claims["id"] = u.ID
			// 有效期 1 年
			claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
			// 生成令牌
			t, err := token.SignedString(signKey)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, map[string]string{
				"token": t,
			})
		}
	}
	return echo.ErrUnauthorized
}

// QR码
func (w *Web) qrcode(c echo.Context) error {
	url, err := GetURL(w.Port)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	code, err := qr.Encode(url, qr.Q)
	if err != nil {
		return c.String(http.StatusInternalServerError, "QR码生成错误")
	}
	return c.Blob(http.StatusOK, "image/png", code.PNG())
}
func (w *Web) info(c echo.Context) error {
	req := c.Request()
	m := make(map[string]string)
	m["Proto"] = req.Proto
	m["Host"] = req.Host
	m["RemoteAddr"] = req.RemoteAddr
	m["Method"] = req.Method
	m["Path"] = req.URL.Path
	return c.JSON(http.StatusOK, m)
}

// 上传文件临时保存
func (w *Web) saveTemp(c echo.Context) (string, error) {
	// 来源
	file, err := c.FormFile("file")
	if err != nil {
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	// 目的
	mkdir(w.Temp)
	f := w.Temp + string(os.PathSeparator) + file.Filename
	dst, err := os.Create(f)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	// 复制
	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}
	return f, nil
}
