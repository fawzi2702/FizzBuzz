package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type OkResponse struct {
	Ok   bool        `json:"ok" example:"true"`
	Data interface{} `json:"data"`
}

// Ok is a helper function to send a 200 OK response
func Ok(c *gin.Context, data interface{}) {
	body := gin.H{"ok": true, "data": data}

	c.JSON(http.StatusOK, body)
}
