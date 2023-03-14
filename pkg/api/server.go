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
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler) *ServerHTTP {

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
		user.GET("disaplyaallproductItems", productHandler.DisaplyaAllProductItems)
		user.GET("disaplyproductItem/:id", productHandler.DisaplyProductItem)
		user.GET("listallproduct", productHandler.ListAllProduct)
		user.GET("listallcategories", productHandler.ListCategories)
		user.GET("findcategories/:id", productHandler.DisplayCategory)
		user.Use(middleware.UserAut)
		{
			user.POST("addtocaart/:id", cartHandler.AddToCart)
			user.PATCH("removefromcart/:id", cartHandler.RemoveFromCart)
			user.GET("listcart", cartHandler.ListCart)
			user.POST("addaddress", userHandler.AddAddress)
			user.PATCH("updateaddress/:id", userHandler.UpdateAddress)
			user.POST("orderall/:id", orderHandler.OrderAll)
			user.PATCH("usercancelordrder/:id", orderHandler.UserCancelOrder)
		}

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
			//product
			admin.POST("addproduct", productHandler.AddProduct)
			admin.PATCH("updateproduct/:id", productHandler.UpdateProduct)
			admin.DELETE("deleteproduct/:id", productHandler.DeleteProduct)
			admin.GET("listallproduct", productHandler.ListAllProduct)
			admin.GET("showproduct/:id", productHandler.ShowProduct)
			//product item
			admin.POST("addproductitem", productHandler.AddProductItem)
			admin.PATCH("updatedproductitem/:id", productHandler.UpdateProductItem)
			admin.DELETE("deleteproductitem/:id", productHandler.DeleteProductItem)
			admin.GET("disaplyaallproductItems", productHandler.DisaplyaAllProductItems)
			admin.GET("disaplyproductitem/:id", productHandler.DisaplyProductItem)

		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
