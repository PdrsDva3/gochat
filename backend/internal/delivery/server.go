package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"gochat/deploy/migration/docs"
	"gochat/internal/delivery/middleware"
	"gochat/internal/delivery/routers"
	"gochat/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, log *log.Logs) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	middlewareStruct := middleware.InitMiddleware(log)
	r.Use(middlewareStruct.CORSMiddleware())

	routers.InitRouting(r, db, log, middlewareStruct) // jwtUtils, session, tracer

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
