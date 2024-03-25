package routers

import (
	"github.com/fawzi2702/FizzBuzz/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupStatsRouter(r *gin.Engine) {
	g := r.Group("/stats")
	{
		g.GET("/", handlers.Stats)
	}
}
