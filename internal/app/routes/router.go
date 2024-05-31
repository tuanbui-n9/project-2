package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanbui-n9/project-2/internal/app/handlers"
	repository "github.com/tuanbui-n9/project-2/internal/app/repositories"
	"github.com/tuanbui-n9/project-2/internal/app/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/", userHandler.CreateUser)
		userRouter.GET("/:id", userHandler.GetUserByID)
		userRouter.PUT("/:id", userHandler.UpdateUser)
		userRouter.DELETE("/:id", userHandler.DeleteUser)
	}

	return router
}
