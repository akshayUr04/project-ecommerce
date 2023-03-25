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
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	couponHandler *handler.CouponHandler,
	favourites *handler.FavouriteHandler) *ServerHTTP {

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

		user.GET("order/razorpay/:id", paymentHandler.CreateRazorpayPayment)
		user.GET("payment-handler", paymentHandler.PaymentSuccess)

		user.Use(middleware.UserAuth)
		{
			user.GET("viewprfile", userHandler.Viewprfile)
			user.PATCH("editprofile", userHandler.UserEditProfile)
			user.PATCH("updatepassword", userHandler.UpdatePassword)
			user.POST("addtocaart/:id", cartHandler.AddToCart)
			user.PATCH("removefromcart/:id", cartHandler.RemoveFromCart)
			user.GET("listcart", cartHandler.ListCart)
			user.POST("addaddress", userHandler.AddAddress)
			user.PATCH("updateaddress/:id", userHandler.UpdateAddress)
			user.POST("orderall/:id", orderHandler.OrderAll)
			user.PATCH("usercancelordrder/:id", orderHandler.UserCancelOrder)
			user.GET("vieworder/:id", orderHandler.ListOrder)
			user.GET("listallorder", orderHandler.ListAllOrders)

			user.GET("userlistallcategories", productHandler.ListCategories)
			user.GET("userfindcategories/:id", productHandler.DisplayCategory)

			user.GET("userlistallproduct", productHandler.ListAllProduct)
			user.GET("usershowproduct/:id", productHandler.ShowProduct)

			user.GET("userdisaplayallproductItems", productHandler.DisaplyaAllProductItems)
			user.GET("userdisaplayproductitem/:id", productHandler.DisaplyProductItem)

			user.PATCH("addcoupontocart/", couponHandler.ApplyCoupon)

			user.POST("addtofav/:productId", favourites.AddToFavourites)
			user.DELETE("removefromfav/:productId", favourites.RemoveFromFav)
			user.GET("viewfav", favourites.ViewFavourites)

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
			//Dashboard
			admin.GET("getdashboard", adminHandler.AdminDashBoard)
			//Coupons
			admin.POST("createcoupon", couponHandler.AddCoupon)
			admin.PATCH("updatecoupen/:couponId", couponHandler.UpdateCoupon)
			admin.DELETE("deletecoupon/:couponId", couponHandler.DeleteCoupon)

		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.LoadHTMLGlob("template/*.html")
	sh.engine.Run(":3000")
}
