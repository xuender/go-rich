package rich

import (
	"crypto/rsa"
	"log"
	"net/http"
	"time"

	"../keys"

	jwt "github.com/dgrijalva/jwt-go"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/xuender/goutils"
)

type Web struct {
	Port string      // 端口号
	Temp string      // 临时文件目录
	Db   string      // 数据库目录
	db   *leveldb.DB // 数据库
	days Days        // 文件日期列表
}

// 文件日期列表主键
var DAYS_KEY = []byte("days")

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

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
	db, err := leveldb.OpenFile(w.Db, nil)
	if err != nil {
		return err
	}
	// 数据库链接
	w.db = db
	// 每日帐目初始化
	w.Get(DAYS_KEY, &w.days)
	return nil
}

// 数据库数据读取
func (w *Web) Get(key []byte, p interface{}) error {
	data, err := w.db.Get(key, nil)
	if err != nil {
		return err
	}
	return goutils.Decode(data, p)
}

// 数据库数据保存
func (w *Web) Put(key []byte, p interface{}) error {
	bs, err := goutils.Encode(p)
	if err != nil {
		return err
	}
	return w.db.Put(key, bs, nil)
}

// 迭代获取数据
func (w *Web) Iterator(prefix []byte, f func(bs []byte)) {
	iter := w.db.NewIterator(util.BytesPrefix(prefix), nil)
	for iter.Next() {
		f(iter.Value())
	}
	iter.Release()
}

func (w *Web) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// 跨域访问
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8100"},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	e.GET("/test", w.test)
	e.GET("/days", w.getDays)
	e.POST("/login", w.login)
	// 需要身份认证
	api := e.Group("/api")
	// api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey:    verifyKey,
	// 	SigningMethod: "RS256",
	// }))
	w.customerRoute(api.Group("/c"))
	// 静态资源
	e.Static("/", "www")
	// e.Use(static.ServeRoot("/", getAssets("www")))
	// 启动服务
	e.Start(w.Port)
}

func (w *Web) Close() {
	w.db.Close()
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
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// TODO 身份认证
	if u.Phone == "139" && u.Password == "123" {
		// 创建令牌
		token := jwt.New(jwt.SigningMethodRS256)
		// 设置用户信息
		claims := token.Claims.(jwt.MapClaims)
		claims["nick"] = "测试"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
		// 生成令牌
		t, err := token.SignedString(signKey)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
	return echo.ErrUnauthorized
}

// 测试
func (w *Web) test(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (w *Web) getDays(c echo.Context) error {
	return c.JSON(http.StatusOK, w.days)
}
