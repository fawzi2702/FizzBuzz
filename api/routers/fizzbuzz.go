package routers

import (
	"github.com/fawzi2702/FizzBuzz/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupFizzbuzzRouter(r *gin.Engine) {
	g := r.Group("/fizzbuzz")
	{
		g.GET("/", handlers.Fizzbuzz)
	}
}
