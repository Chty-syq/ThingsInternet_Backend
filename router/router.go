package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*")  // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.Use(Cors()) //允许跨域
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexPageHandler)
	r.GET("/index", controller.IndexPageHandler)
	r.GET("/login", controller.LoginPageHandler)
	r.GET("/register",controller.RegisterPageHandler)
	v1Group := r.Group("v1")
	{
		v1Group.POST("/login", controller.LoginHandler)
		v1Group.POST("/register", controller.RegisterHandler)
		v1Group.POST("/getOnlineDeviceNum", controller.GetOnlineDeviceNum)
		v1Group.POST("/getTotalInfo", controller.GetTotalInfo)
		v1Group.POST("/getAlertInfo", controller.GetAlertInfo)
		v1Group.POST("/getDeviceInfo", controller.GetDeviceInfo)
		v1Group.POST("/getDeviceNameInfo", controller.GetDeviceNameInfo)
		v1Group.POST("/clearDeviceInfo", controller.ClearDeviceInfoHandler)
		v1Group.POST("/modifyDeviceName", controller.ModifyDeviceNameHandler)
		v1Group.POST("/searchDeviceInfo", controller.SearchDeviceInfo)
	}
	return r
}
