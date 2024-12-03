package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gochat/pkg/log"
)

func InitRouting(r *gin.Engine, db *sqlx.DB, logger *log.Logs) {
	_ = RegisterUserRouter(r, db, logger)
	_ = RegisterFriendRouter(r, db, logger)
}
