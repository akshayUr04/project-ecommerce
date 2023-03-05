package http

import (
	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	handler "github.com/akshayur04/project-ecommerce/pkg/api/handler"
	"github.com/akshayur04/project-ecommerce/pkg/api/middleware"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	otpHandler *handler.OtpHandler,
	adminHandler *handler.AdminHandler) *ServerHTTP {
	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/")
	admin := engine.Group("/admin")

	user.POST("signup", userHandler.UserSignUp)
	user.POST("userlogin", userHandler.UserLogin)
	user.POST("sendotp", otpHandler.SendOtp)
	user.POST("verifyotp", otpHandler.ValidateOtp)
	user.POST("logout", userHandler.UserLogout)

	admin.POST("/adminlogin", adminHandler.AdminLoging)
	admin.POST("/creatadmin", adminHandler.CreateAdmin)

	admin.Use(middleware.AdminAuth)
	admin.POST("adminlogout", adminHandler.AdminLogout)
	admin.POST("blockuser", adminHandler.BlockUser)
	admin.PATCH("unblockuser/:id", adminHandler.UnblockUser)
	admin.GET("finduser/:id", adminHandler.FindUser)
	admin.GET("findall", adminHandler.FindAllUsers)

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
