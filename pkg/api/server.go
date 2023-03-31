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

	user := engine.Group("/user")
	{
		user.POST("signup", userHandler.UserSignUp)
		user.POST("login", userHandler.UserLogin)
		user.POST("logout", userHandler.UserLogout)

		//Otp
		otp := user.Group("/otp")
		{
			otp.POST("send", otpHandler.SendOtp)
			otp.POST("verify", otpHandler.ValidateOtp)
		}

		products := user.Group("/products")
		{
			products.GET("listallproductItems", productHandler.DisaplyaAllProductItems)
			products.GET("disaplyproductItem/:id", productHandler.DisaplyProductItem)

			products.GET("listallproduct", productHandler.ListAllProduct)
			products.GET("showproduct/:id", productHandler.ShowProduct)

			products.GET("listallcategories", productHandler.ListCategories)
			products.GET("findcategories/:id", productHandler.DisplayCategory)
		}

		user.Use(middleware.UserAuth)
		{
			//Profile
			profile := user.Group("/profile")
			{
				profile.GET("view", userHandler.Viewprfile)
				profile.PATCH("edite", userHandler.UserEditProfile)
				profile.PATCH("updatepassword", userHandler.UpdatePassword)
			}

			//Address
			address := user.Group("/address")
			{
				address.POST("add", userHandler.AddAddress)
				address.PATCH("update/:addressId", userHandler.UpdateAddress)
			}

			//Cart
			cart := user.Group("/cart")
			{
				cart.POST("add/:product_item_id", cartHandler.AddToCart)
				cart.PATCH("remove/:product_item_id", cartHandler.RemoveFromCart)
				cart.GET("list", cartHandler.ListCart)
			}

			//Orders
			order := user.Group("/order")
			{
				order.POST("orderall/:paymentId", orderHandler.OrderAll)
				order.PATCH("cancel/:orderId", orderHandler.UserCancelOrder)
				order.GET("view/:orderId", orderHandler.ListOrder)
				order.GET("listall", orderHandler.ListAllOrders)
				order.PATCH("return/:orderId", orderHandler.ReturnOrder)
			}

			//Coupon
			coupon := user.Group("coupon")
			{
				coupon.PATCH("applay/:code", couponHandler.ApplayCoupon)
				coupon.PATCH("remove", couponHandler.RemoveCoupon)
			}

			//Favourites
			favourite := user.Group("/favourites")
			{

				favourite.POST("add/:productId", favourites.AddToFavourites)
				favourite.DELETE("remove/:productId", favourites.RemoveFromFav)
				favourite.GET("view", favourites.ViewFavourites)

			}

			//Payment
			user.GET("order/razorpay/:orderId", paymentHandler.CreateRazorpayPayment)
			user.GET("payment-handler", paymentHandler.PaymentSuccess)
		}

	}

	admin := engine.Group("/admin")
	{
		admin.POST("/adminlogin", adminHandler.AdminLoging)

		admin.Use(middleware.AdminAuth)
		{
			admin.POST("createadmin", adminHandler.CreateAdmin)
			admin.POST("logout", adminHandler.AdminLogout)

			adminUsers := admin.Group("/user")
			{
				adminUsers.PATCH("/block", adminHandler.BlockUser)
				adminUsers.PATCH("/unblock/:user_id", adminHandler.UnblockUser)
				adminUsers.GET("find/:user_id", adminHandler.FindUser)
				adminUsers.GET("findall", adminHandler.FindAllUsers)
			}

			//categorys
			category := admin.Group("/category")
			{
				category.POST("add", productHandler.CreateCategory)
				category.PATCH("update/:id", productHandler.UpdatCategory)
				category.DELETE("delete/:category_id", productHandler.DeleteCategory)
				category.GET("listall", productHandler.ListCategories)
				category.GET("find/:id", productHandler.DisplayCategory)
			}

			//product
			product := admin.Group("/product")
			{
				product.POST("add", productHandler.AddProduct)
				product.PATCH("update/:id", productHandler.UpdateProduct)
				product.DELETE("delete/:id", productHandler.DeleteProduct)
				product.GET("listall", productHandler.ListAllProduct)
				product.GET("show/:id", productHandler.ShowProduct)
			}

			//product item
			productItem := admin.Group("/product-item")
			{
				productItem.POST("add", productHandler.AddProductItem)
				productItem.PATCH("update/:id", productHandler.UpdateProductItem)
				productItem.DELETE("delete/:id", productHandler.DeleteProductItem)
				productItem.GET("listall", productHandler.DisaplyaAllProductItems)
				productItem.GET("show/:id", productHandler.DisaplyProductItem)
				productItem.POST("uploadimage/:id", productHandler.UploadImage)
			}

			//Dashboard
			dashboard := admin.Group("/dashboard")
			{
				dashboard.GET("get", adminHandler.AdminDashBoard)
			}

			//Coupons
			coupon := admin.Group("/coupon")
			{
				coupon.POST("create", couponHandler.CreateCoupon)
				coupon.PATCH("update/:couponId", couponHandler.UpdateCoupon)
				coupon.DELETE("delete/:couponId", couponHandler.DeleteCoupon)
				coupon.GET("view/:couponId", couponHandler.ViewCoupon)
				coupon.GET("viewall", couponHandler.ViewCoupons)
			}

			//Sales report
			sales := admin.Group("/sales")
			{
				sales.GET("get", adminHandler.ViewSalesReport)
				sales.GET("download", adminHandler.DownloadSalesReport)
			}

		}

	}

	return &ServerHTTP{engine: engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.LoadHTMLGlob("template/*.html")
	sh.engine.Run(":3000")
}
