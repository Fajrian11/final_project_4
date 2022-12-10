package router

import (
	"final_project_4/config"
	"final_project_4/controllers"
	"final_project_4/database"
	"final_project_4/middlewares"
	"final_project_4/repositories"
	"final_project_4/service"

	"github.com/gin-gonic/gin"
)

func StartAPP() *gin.Engine {
	cfg := config.LoadConfig()
	db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)

	// USER
	userRepo := repositories.NewUserRepo(db)
	userService := service.NewUserService(&userRepo)
	userController := controllers.NewUserController(userService)
	// CATEGORY
	categoryRepo := repositories.NewCategoryRepo(db)
	categoryService := service.NewCategoryService(&categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)
	// PRODUCT
	productRepo := repositories.NewProductRepo(db)
	productService := service.NewProductService(&productRepo)
	productController := controllers.NewProductController(productService)
	// TRANSACTION
	transactionRepo := repositories.NewTransactionRepo(db)
	transactionService := service.NewTransactionService(&transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.UserRegisterControllers)
		userRouter.POST("/login", userController.UserLoginControllers)

		userRouter.Use(middlewares.Authentication())
		userRouter.PATCH("/topup", userController.TopUpBalanceControllers)
		userRouter.PUT("/update-account", userController.UpdateUserControllers)
		userRouter.DELETE("/delete-account", userController.DeleteUserControllers)
	}
	categoryRouter := router.Group("/categories")
	{
		categoryRouter.Use(middlewares.Authentication())
		categoryRouter.POST("/", middlewares.RoleAuthorization(), categoryController.CreateCategoryControllers)
		categoryRouter.GET("/", middlewares.RoleAuthorization(), categoryController.GetAllCategoryControllers)
		categoryRouter.PATCH("/:categoryId", middlewares.CategoryAuthorization(), categoryController.UpdateCategoryControllers)
		categoryRouter.DELETE("/:categoryId", middlewares.CategoryAuthorization(), categoryController.DeleteCategoryControllers)
	}
	productRouter := router.Group("/products")
	{
		productRouter.Use(middlewares.Authentication())
		productRouter.POST("/", middlewares.RoleAuthorization(), productController.CreateProductControllers)
		productRouter.GET("/", productController.GetAllProductControllers)
		productRouter.PUT("/:productId", middlewares.ProductAuthorization(), productController.UpdateProductControllers)
		productRouter.DELETE("/:productId", middlewares.ProductAuthorization(), productController.DeleteProductControllers)
	}
	transactionRouter := router.Group("/transactions")
	{
		transactionRouter.Use(middlewares.Authentication())
		transactionRouter.POST("/", transactionController.CreateTransactionControllers)
		transactionRouter.GET("/my-transaction", transactionController.GetMyTransactionController)
		transactionRouter.GET("/user-transaction", middlewares.RoleAuthorization(), transactionController.GetAllUserTransactionController)
	}

	return router
}
