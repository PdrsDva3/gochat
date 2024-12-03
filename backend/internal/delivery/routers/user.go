package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gochat/internal/delivery/handlers"
	"gochat/internal/repository/user"
	userserv "gochat/internal/service/user"
	"gochat/pkg/log"
)

func RegisterUserRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	userRouter := r.Group("/user")

	userRepo := user.InitUserRepository(db)
	userChangeRepo := user.InitUserChangeRepository(db)
	userService := userserv.InitUserService(userRepo, logger)
	userChangeService := userserv.InitUserChangeService(userChangeRepo, logger)
	userHandler := handlers.InitUserHandler(userService, userChangeService)

	userRouter.POST("/create", userHandler.Create)
	userRouter.POST("/login", userHandler.Login)
	userRouter.GET("/:id", userHandler.Get)
	userRouter.PUT("/pwd", userHandler.ChangePWD)
	userRouter.PUT("/phone", userHandler.ChangePhone)
	userRouter.PUT("/me", userHandler.Change)
	userRouter.PUT("/email", userHandler.ChangeEmail)
	userRouter.DELETE("/:id", userHandler.Delete)
	return userRouter
}
