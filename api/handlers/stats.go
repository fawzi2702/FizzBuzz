package handlers

import (
	"fmt"
	"net/http"

	"github.com/fawzi2702/FizzBuzz/api/helpers/response"
	fizzbuzzParameters "github.com/fawzi2702/FizzBuzz/api/modules/fizzbuzz/parameters"
	"github.com/fawzi2702/FizzBuzz/api/modules/stats"
	"github.com/gin-gonic/gin"
)

type StatsResponse struct {
	Params *fizzbuzzParameters.Params `json:"params"`
	Score  int                        `json:"score"`
}

// GetTopRequest godoc
// @Summary Get the most requested parameters
// @Description Get the most requested parameters
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {object} response.OkResponse{data=StatsResponse}
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /stats [get]
func Stats(c *gin.Context) {
	p, s, err := stats.GetTopRequest()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err)
		return
	}

	if p == nil {
		response.Error(c, http.StatusNotFound, fmt.Errorf("no request has been made yet"))
		return
	}

	response.Ok(c, &StatsResponse{Params: p, Score: s})
}
