package main

import(
	"go-roadmap/handlers"
	"go-roadmap/services"
	"go-roadmap/repository"
	"go-roadmap/config"
	"go-roadmap/middleware"

	"github.com/gin-gonic/gin"
)

func main()  {
	r:=gin.Default()

	db:=config.ConnectDB()

	//repo
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	bookRepo := repository.NewBookRepository(db)
	//service
	productService := services.NewProductService(productRepo) 
	bookService := services.NewBookService(bookRepo) 
	userService := services.NewUserService(userRepo) 
	authService := services.NewAuthService(userRepo) 

	//handlers
	productHandler := handlers.NewProductHandler(productService)
	bookHandler := handlers.NewBookHandler(bookService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)

	api := r.Group("/api")
	{
		products:= api.Group("/products")
		{
			products.GET("",productHandler.GetProducts)
			products.POST("",productHandler.CreateProduct)
			products.PUT("/:id",productHandler.UpdateProduct)
			products.DELETE("/:id",productHandler.DeleteProduct)
		}
		books:= api.Group("/books")
		{
			books.GET("",bookHandler.GetBooks)
			books.PUT("/:id",bookHandler.UpdateBook)
			books.DELETE("/:id",bookHandler.DeleteBook)
		}
		// users:= api.Group("/users")
		// {
		// 	users.GET("",userHandler.GetUsers)
		// 	users.POST("",userHandler.CreateUser)
		// 	users.PUT("/:id",userHandler.UpdateUser)
		// 	users.DELETE("/:id",userHandler.DeleteUser)
		// }
		public:= api.Group("/auth")
		{
			public.POST("/register", authHandler.Register)
			public.POST("/login", authHandler.Login)
		}
		protected:= api.Group("/admin")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/users", userHandler.GetUsers)
			protected.POST("/books", bookHandler.CreateBook)
		}
	}
	
	r.Run(":8080")
}