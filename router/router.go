package router

import (
	"Hello-gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())


	//加载静态资源
	r.LoadHTMLGlob("dist/index.html")
	r.Static("/css", "./dist/css")
	r.Static("/fonts", "./dist/fonts")
	r.Static("/img", "./dist/img")
	r.Static("/js", "./dist/js")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")

	// 首页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})



	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		//v1.POST("user/register", api.UserRegister)

		// 用户登录
		//v1.POST("user/login", api.UserLogin)

		// 需要登录保护的
		//auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		//{
		//	// User Routing
		//	auth.GET("user/me", api.UserMe)
		//	auth.DELETE("user/logout", api.UserLogout)
		//}
	}
	return r
}
