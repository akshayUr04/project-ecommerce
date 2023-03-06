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
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler) *ServerHTTP {

	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	// Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user := engine.Group("/")
	{
		user.POST("signup", userHandler.UserSignUp)
		user.POST("userlogin", userHandler.UserLogin)
		user.POST("sendotp", otpHandler.SendOtp)
		user.POST("verifyotp", otpHandler.ValidateOtp)
		user.POST("logout", userHandler.UserLogout)
	}

	admin := engine.Group("/admin")
	{
		admin.POST("/adminlogin", adminHandler.AdminLoging)

		admin.Use(middleware.AdminAuth)
		{
			admin.POST("creatadmin", adminHandler.CreateAdmin)
			admin.POST("adminlogout", adminHandler.AdminLogout)
			admin.POST("blockuser", adminHandler.BlockUser)
			admin.PATCH("unblockuser/:id", adminHandler.UnblockUser)
			admin.GET("finduser/:id", adminHandler.FindUser)
			admin.GET("findall", adminHandler.FindAllUsers)
			//categorys
			admin.POST("addcatergory", productHandler.CreateCategory)
			admin.PATCH("updatedcategory/:id", productHandler.UpdatCategory)
			admin.DELETE("deletecategory/:id", productHandler.DeleteCategory)
			admin.GET("listallcategories", productHandler.ListCategories)
			admin.GET("findcategories/:id", productHandler.DisplayCategory)
			admin.POST("addproduct", productHandler.AddProduct)
			admin.PATCH("updateproduct/:id", productHandler.UpdateProduct)
			admin.DELETE("deleteproduct/:id", productHandler.DeleteProduct)
			admin.POST("addproductitem", productHandler.AddProductItem)
			admin.PATCH("updatedproductitem/:id", productHandler.UpdateProductItem)
		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
