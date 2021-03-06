package rich

import (
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Code-Hex/echo-static"
	"github.com/dgrijalva/jwt-go"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
	"github.com/syndtr/goleveldb/leveldb"
	"golang.org/x/crypto/acme/autocert"
	"rsc.io/qr"

	"../keys"
)

// Web 网络服务
type Web struct {
	Temp    string                      // 临时文件目录
	DB      string                      // 数据库目录
	Dev     bool                        // 开发模式
	URL     string                      // 网址
	LogFile string                      // 日志文件
	App     *App                        // 应用信息
	db      *leveldb.DB                 // 数据库
	days    Days                        // 文件日期列表
	cache   map[interface{}]interface{} // 缓存
}

// DaysKey 文件日期列表主键
var DaysKey = []byte("days")

// DBInit 数据库初始化
func (w *Web) DBInit() error {
	var err error
	if w.db, err = leveldb.OpenFile(w.DB, nil); err != nil {
		log.Println(w.db, w.DB)
		return err
	}
	// 每日帐目初始化
	if err := w.Get(DaysKey, &w.days); err != nil {
		w.days = []string{}
	}
	return nil
}

// Init 初始化.
func (w *Web) Init() error {
	if err := w.DBInit(); err != nil {
		return err
	}
	// 用户初始化
	w.UserInit()
	if len(w.days) > 0 {
		w.App.IsNew = false
	}
	if w.App.IsNew {
		color.Println("新用户提示:", color.Yellow("尽快修改初始密码", color.B))
		color.Println(color.Black("      登录帐号      ", color.B, color.YelBg))
		color.Println(color.Cyan(" 登录昵称 ", color.U), " : ", color.Green("admin", color.B))
		color.Println(color.Cyan(" 初始密码 ", color.U), " : ", color.Green(strings.Replace(w.App.Version, ".", "", -1), color.B))
		color.Println("创建首笔订单后不再提示以上信息")
	}
	// 缓存初始化
	w.cache = map[interface{}]interface{}{}
	return nil
}

func (w *Web) initEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.HTTPErrorHandler = w.httpErrorHandler
	// 开发模式
	if w.Dev {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
		middleware.DefaultLoggerConfig.Format = `${time_rfc3339_nano} [${remote_ip}] ${host}(${method})${uri}(${status}) ${error} ${latency} ` +
			`[${latency_human}] IN:${bytes_in} OUT:${bytes_out}` + "\n"
		e.Use(middleware.Recover())
		// 跨域访问
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
			AllowCredentials: true,
		}))
	} else {
		e.HidePort = true
		if f, err := os.OpenFile(w.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); err == nil {
			middleware.DefaultLoggerConfig.Output = f
			log.SetOutput(f)
		}
	}
	e.Use(middleware.Logger())
	// 二维码访问
	e.GET("/qr", func(c echo.Context) error {
		code, err := qr.Encode(w.URL, qr.Q)
		if err != nil {
			return c.String(http.StatusInternalServerError, "QR码生成错误: "+err.Error())
		}
		return c.Blob(http.StatusOK, "image/png", code.PNG())
	})
	// 应用信息
	e.GET("/about", func(c echo.Context) error {
		return c.JSON(http.StatusOK, w.App)
	})
	e.GET("/login", w.login) // 登录
	api := e.Group("/api")   // API
	// 需要身份认证
	api.Use(middlewareJWT(w, "HS256"))

	w.customerRoute(api.Group("/customers")) // 客户
	w.itemRoute(api.Group("/items"))         // 商品
	w.tagRoute(api.Group("/tags"))           // 标签
	w.extsRoute(api.Group("/exts"))          // 扩展定义
	w.xlsxRoute(api.Group("/xlsxes"))        // Excel定义
	w.userRoute(api.Group("/users"))         // 用户
	w.profileRoute(api.Group("/profile"))    // 账户
	w.tradeRoute(api.Group("/trades"))       // 订单
	// 静态资源
	if w.Dev {
		e.Static("/", "www")
	} else {
		e.Use(static.ServeRoot("/", getAssets("www")))
	}
	log.Println("Go Rich 启动...")
	return e
}

// Start starts an HTTP server.
func (w *Web) Start(address string) error {
	return w.initEcho().Start(address)
}

// StartTLS starts an HTTPS server.
func (w *Web) StartTLS(address, certFile, keyFile string) error {
	if certFile == "" || keyFile == "" {
		return w.startTLS(address)
	}
	return w.initEcho().StartTLS(address, certFile, keyFile)
}

func (w *Web) startTLS(address string) error {
	e := w.initEcho()
	s := e.TLSServer
	s.Addr = address
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
	if !e.DisableHTTP2 {
		s.TLSConfig.NextProtos = append(s.TLSConfig.NextProtos, "h2")
	}
	return e.StartServer(e.TLSServer)
}

// StartAutoTLS starts an HTTPS server using certificates automatically installed from https://letsencrypt.org.
func (w *Web) StartAutoTLS(address string) error {
	e := w.initEcho()
	e.AutoTLSManager.Cache = autocert.DirCache(w.Temp)
	return e.StartAutoTLS(address)
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
				if es.Code == 404 {
					c.Redirect(http.StatusMovedPermanently, "/")
					return
				}
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

// Close 关闭服务
func (w *Web) Close() {
	w.db.Close()
	log.Println("Go Rich 关闭.")
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

// Signed 创建令牌
func (w *Web) Signed(id string, pass []byte) (string, error) {
	// 创建令牌
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置用户信息
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	// 有效期 1 年
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	// 生成令牌
	return token.SignedString(pass)
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
	passBs := Pass(pass)
	u := w.FindUser(nick)
	// 身份认证
	if u != nil && bytes.Equal(passBs, u.Pass) {
		t, err := w.Signed(u.ID.String(), passBs)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, t)
	}
	return echo.ErrUnauthorized
}

// 查找用户
func (w *Web) FindUser(nick string) *User {
	for _, u := range w.users() {
		if u.Name == nick || u.Phone == nick {
			return u
		}
	}
	return nil
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
	if err := Mkdir(w.Temp); err != nil {
		return "", err
	}
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
