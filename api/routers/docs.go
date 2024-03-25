package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupDocsRouter(r *gin.Engine) {
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
