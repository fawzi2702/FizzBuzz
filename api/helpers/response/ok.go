package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ok is a helper function to send a 200 OK response
func Ok(c *gin.Context, data interface{}) {
	body := gin.H{"ok": true, "data": data}

	c.JSON(http.StatusOK, body)
}
