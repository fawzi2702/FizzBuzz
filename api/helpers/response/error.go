package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Ok    bool   `json:"ok" example:"false"`
	Error string `json:"error"`
}

// Error is an helper function to send an error response
func Error(c *gin.Context, status int, err error) {
	body := gin.H{"ok": false, "error": err.Error()}

	c.JSON(status, body)
}
