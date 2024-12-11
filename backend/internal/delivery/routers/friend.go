package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gochat/internal/delivery/handlers"
	"gochat/internal/repository/friend"
	friendserv "gochat/internal/service/friend"
	"gochat/pkg/log"
)

func RegisterFriendRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	friendRouter := r.Group("/friend")

	friendRepo := friend.InitFriendRepository(db)
	friendService := friendserv.InitFriendService(friendRepo, logger)
	friendHandler := handlers.InitFriendHandler(friendService)

	friendRouter.POST("/", friendHandler.AddFriend)
	friendRouter.GET("/:id", friendHandler.Get)
	friendRouter.GET("/list/:id", friendHandler.GetFriendsInfo)
	friendRouter.DELETE("/:id_1/:id_2", friendHandler.Delete)
	return friendRouter
}
