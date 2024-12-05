package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gochat/internal/delivery/handlers"
	"gochat/internal/repository/chat"
	chatserv "gochat/internal/service/chat"
	"gochat/pkg/log"
)

func RegisterChatRouter(r *gin.Engine, db *sqlx.DB, logger *log.Logs) *gin.RouterGroup {
	chatRouter := r.Group("/chat")

	chatRepo := chat.InitChatRepository(db)
	chatService := chatserv.InitChatService(chatRepo, logger)
	chatHandler := handlers.InitChatHandler(chatService)

	chatRouter.POST("/create", chatHandler.Create)
	chatRouter.GET("/list", chatHandler.List)
	chatRouter.GET("/:id", chatHandler.Get)
	chatRouter.DELETE("/:id", chatHandler.Delete)
	return chatRouter
}
