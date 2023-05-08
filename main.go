package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/ydhnwb/golang_heroku/handler/v1"
)

var (
// db             *gorm.DB               = config.SetupDatabaseConnection()
// userRepo       repo.UserRepository    = repo.NewUserRepo(db)
// productRepo    repo.ProductRepository = repo.NewProductRepo(db)
// authService    service.AuthService    = service.NewAuthService(userRepo)
// jwtService     service.JWTService     = service.NewJWTService()
// userService    service.UserService    = service.NewUserService(userRepo)
// productService service.ProductService = service.NewProductService(productRepo)
// authHandler v1.AuthHandler = v1.NewAuthHandler(authService, jwtService, userService)
// userHandler v1.UserHandler = v1.NewUserHandler(userService, jwtService)

// productHandler v1.ProductHandler      = v1.NewProductHandler(productService, jwtService)
)

func main() {
	// defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// authRoutes := server.Group("api/auth")
	// {
	// 	authRoutes.POST("/login", authHandler.Login)
	// 	authRoutes.POST("/register", authHandler.Register)
	// }

	// userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	// {
	// 	userRoutes.GET("/profile", userHandler.Profile)
	// 	userRoutes.PUT("/profile", userHandler.Update)
	// }

	// productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	// {
	// 	productRoutes.GET("/", productHandler.All)
	// 	productRoutes.POST("/", productHandler.CreateProduct)
	// 	productRoutes.GET("/:id", productHandler.FindOneProductByID)
	// 	productRoutes.PUT("/:id", productHandler.UpdateProduct)
	// 	productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	// }

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
		checkRoutes.GET("address/:address", v1.Address)
		// checkRoutes.GET("airdrop", v1.TotalSupply)
	}

	server.Run()
}
