package handlers

import (
	"fmt"
	"net/http"

	"github.com/fawzi2702/FizzBuzz/api/helpers/response"
	"github.com/fawzi2702/FizzBuzz/api/modules/stats"
	"github.com/gin-gonic/gin"
)

func Stats(c *gin.Context) {
	p, err := stats.GetTopRequest()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	if p == nil {
		response.Error(c, http.StatusNotFound, fmt.Errorf("no request has been made yet"))
		return
	}

	response.Ok(c, p)
}
