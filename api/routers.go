package api

import (
	_ "api_gateway/api/docs"
	"api_gateway/api/handlers"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Restaurant Reservation System API
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host localhost:8080
type Server struct {
	Handlers handlers.MainHandler
}

func NewServer(hands handlers.MainHandler) *Server {
	return &Server{Handlers: hands}
}

func (s *Server) InitRoutes(r *gin.Engine) {
	r.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	auth := s.Handlers.Auth()
	product := s.Handlers.Product()

	api := r.Group("/api/v1")
	{

		authGroup := api.Group("/auth")
		{
			authGroup.POST("/login", auth.Login)
			authGroup.POST("/register", auth.Register)
			authGroup.GET("/profile/:id", auth.GetUserInfo)
		}

		productGroup := api.Group("/products")
		{
			productGroup.POST("", product.AddProduct)
			productGroup.PUT("/:id", product.EditProduct)
			productGroup.DELETE("/:id", product.DeleteProduct)
			productGroup.GET("", product.GetProducts)
			productGroup.GET("/:id", product.GetProduct)
			productGroup.POST("/search", product.SearchProducts)
			productGroup.POST("/:id/rate", product.AddRating)
			productGroup.GET("/:id/ratings", product.GetRatings)
		}

		orderGroup := api.Group("/orders")
		{
			orderGroup.POST("", product.PlaceOrder)
			orderGroup.PUT("/:id/cancel", product.CancelOrder)
			orderGroup.PUT("/:id/status", product.UpdateOrderStatus)
			orderGroup.GET("", product.GetOrders)
			orderGroup.GET("/:id", product.GetOrder)
			orderGroup.POST("/:id/pay", product.PayOrder)
			orderGroup.GET("/:id/payment-status", product.CheckPaymentStatus)
			orderGroup.PUT("/:id/shipping", product.UpdateShippingInfo)
		}

		categoryGroup := api.Group("/categories")
		{
			categoryGroup.POST("/artisan", product.AddArtisanCategory)
			categoryGroup.POST("/product", product.AddProductCategory)
		}

		api.GET("/statistics", product.GetStatistics)
		api.GET("/user-activity", product.GetUserActivity)
		api.GET("/recommendations", product.GetRecommendations)
		api.GET("/artisan-rankings", product.GetArtisanRankings)
	}

	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	// r.Use(middlewares.JWTMiddlewares)
}