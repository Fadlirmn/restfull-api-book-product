package main

import(
	"go-roadmap/handlers"
	"go-roadmap/services"
	"go-roadmap/repository"
	"go-roadmap/config"

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

	//handlers
	productHandler := handlers.NewProductHandler(productService)
	bookHandler := handlers.NewBookHandler(bookService)
	userHandler := handlers.NewUserHandler(userService)

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
			books.POST("",bookHandler.CreateBook)
			books.PUT("/:id",bookHandler.UpdateBook)
			books.DELETE("/:id",bookHandler.DeleteBook)
		}
		users:= api.Group("/users")
		{
			users.GET("",userHandler.GetUsers)
			users.POST("",userHandler.CreateUser)
			users.PUT("/:id",userHandler.UpdateUser)
			users.DELETE("/:id",userHandler.DeleteUser)
		}
	}
	
	r.Run(":8080")
}